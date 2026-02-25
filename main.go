package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/borisfritz/gator/internal/config"
	"github.com/borisfritz/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	// Initilize Program by reading config file, setting programState,
	// and registering program commands.
	programConfig, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	programState := state{
		cfg: &programConfig,
	}
	programCommands := getProgramCommands()

	// LOAD Database
	db, err := sql.Open("postgres", programConfig.DBURL)
	if err != nil {
		log.Fatalf("Unable to connect to DataBase: %v", err)
	}
	dbQueries := database.New(db)
	programState.db = dbQueries
	
	// Capture command from user
	if len(os.Args) < 2 {
		log.Fatal("Incorrect usage. Gator requires <command> [arguments...]")
	}
	userCommand := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	// Run Captured command
	err = programCommands.run(&programState, userCommand)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
