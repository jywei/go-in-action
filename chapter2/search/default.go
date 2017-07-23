package search

// defaultMatcher implements the default matcher
type defaultMatcher struct{}

// init registers the default matcher with the program
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

// the implementation of the Matcher interface by the defaultMatcher type
// type Matcher interface {
// 	Search(feed *Feed, searchTerm string) ([]*Result, error)
// }
