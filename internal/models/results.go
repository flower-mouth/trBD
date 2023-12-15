package models

type ChessPlayersResults struct {
	Participant1FIO    string
	Participant2FIO    string
	PointsParticipant1 float64
	PointsParticipant2 float64
}

type NonChessPlayersResults struct {
	Participant1FIO    string
	Participant2FIO    string
	PointsParticipant1 float64
	PointsParticipant2 float64
}

type TournamentResults struct {
	ParticipantID int
	FIO           string
	TotalPoints   float64
}
