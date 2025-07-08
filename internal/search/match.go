package search

import (
	"fmt"
	"log"
)

// Result is a feed search Result.
type Result struct {
	Field   string
	Content string
}

// Matcher defines the behavior requred by types that want
// to implement a new search type.
// Interfaces tipically represents a sdingle action, when an interface
// has a single method,
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for search each individual feed to run
// searches concurrently.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Print(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

// Display display the search results coming from a channel.
func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%sn\n\n", result.Field, result.Content)
	}
}
