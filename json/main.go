package main

import (
	"log"
	"net/http"
)

const sqliteFile = "../players.db"

func main() {
	store := NewSqlitePlayerStore(sqliteFile)
	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5001", server))
}
