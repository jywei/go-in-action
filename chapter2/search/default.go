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

// Method declared with a value receiver of type defaultMatcher
// func (m defaultMatcher) Search(feed *Feed, searchTerm string)
// Declare a pointer of type defaultMatch
// dm := new(defaultMatch)
// The compiler will dereference the dm pointer to make the call
// dm.Search(feed, "test")
// Method declared with a pointer receiver of type defaultMatcher
// func (m *defaultMatcher) Search(feed *Feed, searchTerm string)
// Declare a value of type defaultMatch
// var dm defaultMatch
// The compiler will reference the dm value to make the call
// dm.Search(feed, "test")
