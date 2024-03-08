package search

import (
  "encoding/json"
  "os"
)

const dataFile = "data/data.json"

type Feed struct {
  Name string `json:"site"`
  Link string `json:"link"`
  Type string `json:"site"`
}

func RetrieveFeeds()([]*Feed, error) {
  file, err := os.OpenFile(dataFile)
  if err != nil {
    return nil,err
  }

  defer file.close()

  var feeds []*Feed
  err = json.NewEncoder(file).Decode(&feeds)

  //don´t need to verify for errors. The caller is responsible to do so
  return feeds, err
}
