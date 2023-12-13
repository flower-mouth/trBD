-- Database: trbd

	-- Создание таблицы "Участники турнира"
CREATE TABLE Participants (
    ID SERIAL PRIMARY KEY,
    FIO VARCHAR(255) NOT NULL,
    BirthDate DATE NOT NULL,
    GroupNumber VARCHAR(10) NOT NULL,
    PhoneNumber VARCHAR(15) NOT NULL,
    Experience BOOLEAN NOT NULL,
    ParticipantGroup VARCHAR(10) NOT NULL
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
