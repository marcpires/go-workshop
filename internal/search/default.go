package search

type defaultMatcher struct{}

func init() {
  var matcher defaultMatcher
  Register("default", matcher)
}

// Search implements the behavior for the default matcher
// uses a value receiver as it results in zero allocation and we don´t want to
// manipulate any state.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
  return nil, nil
}
