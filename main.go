package main

import (
	"os"

	"github.com/lukebrobbs/pokemon-go/server"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	server.Start(PORT)
}
