package search

import (
	"log"
	"sync"
)

// matcher is map (reference type)
// var are alawys initialized to their zero value:
// false for bool, 0 for int, nil for pointers, slices, maps, channels, interfaces, and function types.
var matchers = make(map[string]Matcher)

// Run uses a matcher to search for feeds and diplay it results.
func Run(searchTerm string) {

	// declare and initialize variables with short variable declaration operator
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all work are done
	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)

}

// Register is called to register a matcher.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
