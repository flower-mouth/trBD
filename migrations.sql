-- Database: trbd

	-- Создание таблицы "Участники турнира"
CREATE TABLE Participants (
    ID SERIAL PRIMARY KEY,
    FIO VARCHAR(255) NOT NULL,
    BirthDate DATE NOT NULL,
    GroupNumber VARCHAR(10) NOT NULL,
    PhoneNumber VARCHAR(15) NOT NULL,
    Experience BOOLEAN NOT NULL
);

-- Создание таблицы "Промежуточные результаты - Умеют играть в шахматы"
CREATE TABLE ChessPlayersPoints (
    ID SERIAL PRIMARY KEY,
    ParticipantID INTEGER REFERENCES Participants(ID),
    Points DECIMAL(3, 1) NOT NULL
);

-- Создание таблицы "Промежуточные результаты - Не умеют играть в шахматы"
CREATE TABLE NonChessPlayersPoints (
    ID SERIAL PRIMARY KEY,
    ParticipantID INTEGER REFERENCES Participants(ID),
    Points DECIMAL(3, 1) NOT NULL
);

-- Создание таблицы "Турнирная таблица - Умеют играть в шахматы"
CREATE TABLE ChessPlayersResults (
    ID SERIAL PRIMARY KEY,
    Participant1ID INTEGER REFERENCES Participants(ID),
    Participant2ID INTEGER REFERENCES Participants(ID),
    PointsParticipant1 DECIMAL(3, 1) NOT NULL,
    PointsParticipant2 DECIMAL(3, 1) NOT NULL
);

-- Создание таблицы "Турнирная таблица - Не умеют играть в шахматы"
CREATE TABLE NonChessPlayersResults (
    ID SERIAL PRIMARY KEY,
    Participant1ID INTEGER REFERENCES Participants(ID),
    Participant2ID INTEGER REFERENCES Participants(ID),
    PointsParticipant1 DECIMAL(3, 1) NOT NULL,
    PointsParticipant2 DECIMAL(3, 1) NOT NULL
);

create function update_chessplayerspoints() returns trigger
    language plpgsql
as
$$
BEGIN
    -- Обновление значений в таблице chessplayerspoints для participant1id
    UPDATE chessplayerspoints
    SET points = points + NEW.pointsparticipant1
    WHERE participantid = NEW.participant1id;

    -- Если participant1id отсутствует, добавляем новую запись
    IF NOT FOUND THEN
        INSERT INTO chessplayerspoints (participantid, points)
        VALUES (NEW.participant1id, NEW.pointsparticipant1);
    END IF;

    -- Обновление значений в таблице chessplayerspoints для participant2id
    UPDATE chessplayerspoints
    SET points = points + NEW.pointsparticipant2
    WHERE participantid = NEW.participant2id;

    -- Если participant2id отсутствует, добавляем новую запись
    IF NOT FOUND THEN
        INSERT INTO chessplayerspoints (participantid, points)
        VALUES (NEW.participant2id, NEW.pointsparticipant2);
    END IF;

    RETURN NEW;
END;
$$;

alter function update_chessplayerspoints() owner to postgres;

create function update_nonchessplayerspoints() returns trigger
    language plpgsql
as
$$
BEGIN
    -- Обновление значений в таблице nonchessplayerspoints для participant1id
    UPDATE nonchessplayerspoints
    SET points = points + NEW.pointsparticipant1
    WHERE participantid = NEW.participant1id;

    -- Если participant1id отсутствует, добавляем новую запись
    IF NOT FOUND THEN
        INSERT INTO nonchessplayerspoints (participantid, points)
        VALUES (NEW.participant1id, NEW.pointsparticipant1);
    END IF;

    -- Обновление значений в таблице nonchessplayerspoints для participant2id
    UPDATE nonchessplayerspoints
    SET points = points + NEW.pointsparticipant2
    WHERE participantid = NEW.participant2id;

    -- Если participant2id отсутствует, добавляем новую запись
    IF NOT FOUND THEN
        INSERT INTO nonchessplayerspoints (participantid, points)
        VALUES (NEW.participant2id, NEW.pointsparticipant2);
    END IF;

    RETURN NEW;
END;
$$;

alter function update_nonchessplayerspoints() owner to postgres;


-- ТАБЛИЧНАЯ ФУНКЦИЯ --
-- Создание функции для получения результатов турнира
CREATE OR REPLACE FUNCTION get_tournament_results()
    RETURNS TABLE (
                      ParticipantID INTEGER,
                      FIO VARCHAR(255),
                      Points DECIMAL(3, 1)
                  ) AS $$
BEGIN
    -- Получение результатов шахматных игр
    RETURN QUERY
        SELECT
            cp.ParticipantID,
            p.FIO,
            cp.Points
        FROM
            ChessPlayersPoints cp
                INNER JOIN
            Participants p ON cp.ParticipantID = p.ID;

    -- Получение результатов игр для тех, кто не играет в шахматы
    RETURN QUERY
        SELECT
            ncp.ParticipantID,
            p.FIO,
            ncp.Points
        FROM
            NonChessPlayersPoints ncp
                INNER JOIN
            Participants p ON ncp.ParticipantID = p.ID;
END;
$$ LANGUAGE plpgsql;


-- СКАЛЯРНАЯ ФУНКЦИЯ --
CREATE OR REPLACE FUNCTION calculate_total_points(participant_id INTEGER)
    RETURNS DECIMAL(3, 1) AS $$
DECLARE
    total_points DECIMAL(3, 1);
BEGIN
    -- Рассчитать общее количество очков для заданного participant_id
    SELECT
        COALESCE(SUM(Points), 0)
    INTO
        total_points
    FROM
        ChessPlayersPoints
    WHERE
            ParticipantID = participant_id;

    -- Добавить очки из NonChessPlayersPoints
    SELECT
        COALESCE(SUM(Points), 0)
    INTO
        total_points
    FROM
        NonChessPlayersPoints
    WHERE
            ParticipantID = participant_id;

    RETURN total_points;
END;
$$ LANGUAGE plpgsql;

/*SELECT
    calculate_total_points(1) AS total_points_for_participant_1;*/

-- ПРЕДСТАВЛЕНИЕ --
CREATE OR REPLACE VIEW TournamentResults AS
SELECT
    p.ID AS ParticipantID,
    p.FIO,
    COALESCE(cp.Points, 0) + COALESCE(ncp.Points, 0) AS TotalPoints
FROM
    Participants p
        LEFT JOIN
    ChessPlayersPoints cp ON p.ID = cp.ParticipantID
        LEFT JOIN
    NonChessPlayersPoints ncp ON p.ID = ncp.ParticipantID;




