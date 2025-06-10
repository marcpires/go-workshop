package main

import (
	"log"
	"os"

	_ "github.com/marcpires/rss/internal/matchers" // Cool stuff happens here related to rss.go
	"github.com/marcpires/rss/internal/search"
)

func init() {
	// Change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

func main() {
	log.Print("Starting search")

	search.Run("president")
}
