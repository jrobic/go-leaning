package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const createPlayerTable = `CREATE TABLE IF NOT EXISTS players (
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	createdAt timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	name TEXT, 
	wins INTEGER NOT NULL DEFAULT 0
	);`

type SqlitePlayerStore struct {
	db *sql.DB
}

func NewSqlitePlayerStore(dbFile string) *SqlitePlayerStore {
	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	if _, err := db.Exec(createPlayerTable); err != nil {
		log.Fatal(err)
		return nil
	}

	return &SqlitePlayerStore{db}
}

func (s *SqlitePlayerStore) GetPlayerScore(name string) int {
	row := s.db.QueryRow("SELECT SUM(wins) FROM players WHERE name = ?;", name)

	var score int
	row.Scan(&score)

	return score
}

func (s *SqlitePlayerStore) RecordWin(name string) {
	_, err := s.db.Exec("INSERT INTO players (name, wins) VALUES(?,?);", name, 1)

	if err != nil {
		log.Fatal(err)
	}
}

func (s *SqlitePlayerStore) GetLeague() []Player {
	rows, err := s.db.Query("SELECT name, SUM(wins) FROM players GROUP BY name")

	if err != nil {
		log.Fatal(err)
	}

	var league []Player

	for rows.Next() {
		player := Player{}

		rows.Scan(&player.Name, &player.Wins)

		league = append(league, player)
	}

	return league
}
