package main

import (
  "os"
  "log"
  _ "github.com/marcpires/rss/matchers"
  "github.com/marcpires/rss/search"
)

func init() {
  log.SetOutput(os.Stdout)
}

func main() {
  search.Run("president")
}
