package search

import "log"

type Result struct {
	Field   string
	Content string
}

// Matcher defines the behavior required by types that want to implement a new search type
// If you have a search method, you are a Matcher interface
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// It takes Feed type's address and a string of searchTerm, then return a slice of Result type's address

// Match is launched as a goroutine for each individual feed to run searches concurrently
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchTerm, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write results to the channel
	for _, result := range searchTerm {
		results <- result
	}
}

// Display writes results to the console, as they are received by the individual goroutines
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
