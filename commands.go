package main

import (
	"errors"
)

// command struct holds command 'name' and <args> from the CLI passed in by the user.
type command struct {
	name string
	args []string
}

// Struct and methods to register and run handler functions
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.registeredCommands[cmd.name]
	if !ok {
		return errors.New("Command not found")
	}
	return handler(s, cmd)
}

func getProgramCommands() commands {
	programCommands := commands{
		make(map[string]func(*state,command)error),
	}
	programCommands.register("login", handlerLogin)
	programCommands.register("register", handlerRegister)
	programCommands.register("reset", handlerReset)
	programCommands.register("users", handlerUsers)
	programCommands.register("agg", handlerAgg)
	programCommands.register("addfeed", handlerAddFeed)
	programCommands.register("feeds", handlerFeeds)
	return programCommands
}
