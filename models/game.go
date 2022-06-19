package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./y.sqlite")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type TogyzGame struct {
	Id        int    `json:"id"`
	WhiteName string `json:"white_name"`
	BlackName string `json:"black_name"`
	Result    string `json:"result"`
	Event     string `json:"event"`
	Date      string `json:"date"`
	Site      string `json:"site"`
	Notation  string `json:"notaton"`
}

func GetGames(count int) ([]TogyzGame, error) {

	rows, err := DB.Query("SELECT id, _WhiteName, _BlackName, _Result, _Event, _Date, _Site, _Notation FROM games LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	games := make([]TogyzGame, 0)

	for rows.Next() {
		game := TogyzGame{}
		err = rows.Scan(&game.Id, &game.WhiteName, &game.BlackName, &game.Result, &game.Event, &game.Date, &game.Site, &game.Notation)

		if err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return games, err
}

func GetGameById(id string) (TogyzGame, error) {

	stmt, err := DB.Prepare("SELECT id, _WhiteName, _BlackName, _Result, _Event, _Date, _Site, _Notation FROM games WHERE id = ?")

	if err != nil {
		return TogyzGame{}, err
	}

	game := TogyzGame{}

	sqlErr := stmt.QueryRow(id).Scan(&game.Id, &game.WhiteName, &game.BlackName, &game.Result, &game.Event, &game.Date, &game.Site, &game.Notation)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return TogyzGame{}, nil
		}
		return TogyzGame{}, sqlErr
	}
	return game, nil
}
