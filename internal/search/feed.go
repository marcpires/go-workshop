package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed defines the feed file structure.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds decodes the feed config data into the Feed structure.
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.OpenFile(dataFile, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
