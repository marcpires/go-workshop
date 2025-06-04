package main

import (
	"log"

	_ "github.com/marcpires/rss/internal/matchers"
	"github.com/marcpires/rss/internal/search"
)

func main() {
	log.Print("Starting search")

	search.Run("president")
}
