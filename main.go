package main

import (
	"fmt"
	"github.com/borisfritz/gator/internal/config"
)

func main() {
	firstCFG := config.Read()
	firstCFG.SetUser("Matthew")
	secondCFG := config.Read()
	fmt.Printf("%+v\n", secondCFG)
}
