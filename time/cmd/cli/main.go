package main

import (
	"fmt"
	"log"
	"os"

	poker "go-learning/time"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.StdOutAlerter), store)

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
