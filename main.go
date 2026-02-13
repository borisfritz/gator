package main

import (
	"log"
	"os"
	"github.com/borisfritz/gator/internal/config"
)

// State struct to save the state of the program.
//
// cfg holds a pointer to a Config struct (internal/config) which 
// is read from the file located at $HOME/.gatorconfig.json
type state struct {
	cfg *config.Config
}

func main() {
	// Initilize Program by reading config file, setting programState,
	// programCommands and registering commands.
	programConfig, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	programState := state{
		cfg: &programConfig,
	}
	programCommands := commands{
		make(map[string]func(*state,command)error),
	}
	programCommands.register("login", handlerLogin)
	
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
