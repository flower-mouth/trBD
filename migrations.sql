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
CREATE TABLE ChessPlayersIntermediateResults (
    ID SERIAL PRIMARY KEY,
    ParticipantID INTEGER REFERENCES Participants(ID),
    Points DECIMAL(3, 1) NOT NULL
);

-- Создание таблицы "Промежуточные результаты - Не умеют играть в шахматы"
CREATE TABLE NonChessPlayersIntermediateResults (
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



