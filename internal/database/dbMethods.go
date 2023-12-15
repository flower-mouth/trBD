package database

import (
	"context"
	"trBD/internal/models"
)

func InsertParticipant(fio, birthdate, groupnumber, phonenumber string, experience bool) error {
	// Начинаем транзакцию
	tx, err := DBClient.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// Подготовка SQL-запроса
	query := "CALL insert_participant($1, $2, $3, $4, $5)"
	_, err = tx.Exec(context.Background(), query, fio, birthdate, groupnumber, phonenumber, experience)
	if err != nil {
		return err
	}

	// Коммитим транзакцию
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func InsertChessPlayerResult(participant1id, participant2id int, pointsparticipant1, pointsparticipant2 float64) error {
	// Начинаем транзакцию
	tx, err := DBClient.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// Подготовка SQL-запроса
	query := "CALL insert_chess_player_result($1, $2, $3, $4)"
	_, err = tx.Exec(context.Background(), query, participant1id, participant2id, pointsparticipant1, pointsparticipant2)
	if err != nil {
		return err
	}

	// Коммитим транзакцию
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func InsertNonChessPlayerResult(participant1id, participant2id int, pointsparticipant1, pointsparticipant2 float64) error {
	// Начинаем транзакцию
	tx, err := DBClient.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// Подготовка SQL-запроса
	query := "CALL insert_non_chess_player_result($1, $2, $3, $4)"
	_, err = tx.Exec(context.Background(), query, participant1id, participant2id, pointsparticipant1, pointsparticipant2)
	if err != nil {
		return err
	}

	// Коммитим транзакцию
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// Получение участников из базы данных
func GetExpParticipantsFromDB(dbClient Client) ([]models.Participants, error) {
	rows, err := dbClient.Query(context.Background(), "SELECT fio FROM Participants WHERE experience = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []models.Participants

	for rows.Next() {
		var p models.Participants
		err := rows.Scan(&p.FIO)
		if err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	return participants, nil
}

// Получение участников из базы данных
func GetZeroExpParticipantsFromDB(dbClient Client) ([]models.Participants, error) {
	rows, err := dbClient.Query(context.Background(), "SELECT fio FROM Participants WHERE experience = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []models.Participants

	for rows.Next() {
		var p models.Participants
		err := rows.Scan(&p.FIO)
		if err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	return participants, nil
}

// Функция получения ID участника по его ФИО
func GetParticipantIDByFIO(dbClient Client, fio string) (int, error) {
	var participantID int

	err := dbClient.QueryRow(context.Background(), "SELECT ID FROM Participants WHERE FIO = $1", fio).Scan(&participantID)
	if err != nil {
		return 0, err
	}

	return participantID, nil
}

// Функция для получения данных из таблицы ChessPlayersPoints
func GetChessPlayersPointsFromDB(dbClient Client) ([]models.ChessPlayersPoints, error) {
	// Запрос к базе данных
	query := "SELECT cp.Points, p.FIO FROM ChessPlayersPoints cp INNER JOIN Participants p ON cp.ParticipantID = p.ID"
	rows, err := dbClient.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pointsData []models.ChessPlayersPoints

	for rows.Next() {
		var pointData models.ChessPlayersPoints
		err := rows.Scan(&pointData.Points, &pointData.ParticipantFIO)
		if err != nil {
			return nil, err
		}
		pointsData = append(pointsData, pointData)
	}

	return pointsData, nil
}

// Функция для получения данных из таблицы ChessPlayersPoints
func GetNonChessPlayersPointsFromDB(dbClient Client) ([]models.NonChessPlayersPoints, error) {
	// Запрос к базе данных
	query := "SELECT cp.Points, p.FIO FROM NonChessPlayersPoints cp INNER JOIN Participants p ON cp.ParticipantID = p.ID"
	rows, err := dbClient.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pointsData []models.NonChessPlayersPoints

	for rows.Next() {
		var pointData models.NonChessPlayersPoints
		err := rows.Scan(&pointData.Points, &pointData.ParticipantFIO)
		if err != nil {
			return nil, err
		}
		pointsData = append(pointsData, pointData)
	}

	return pointsData, nil
}

// Функция для получения данных из представления TournamentResults
func GetTournamentResultsFromDB(dbClient Client) ([]models.TournamentResults, error) {
	// Запрос к базе данных
	query := "SELECT * FROM TournamentResults"
	rows, err := dbClient.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.TournamentResults

	for rows.Next() {
		var result models.TournamentResults
		err := rows.Scan(&result.ParticipantID, &result.FIO, &result.TotalPoints)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
