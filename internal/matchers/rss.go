package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	 "github.com/marcpires/rss/internal/search"
)

type (
  item struct {
    XMLName xml.Name `xmk:"item"`
    PubDate string `xml:"pubDate"`
    Title string `xml:"title"`
    Description string `xml:"description"`
    Link string `xml:"link"`
    GUID string `xml:"guid"`
    GeoRssPoint string `xml:"georss:point"`
  }

  image struct {
    XMLName xml.Name `xml:"image"`
    URL string `xml:"url"`
    Title string `xml:"title"`
    Link string `xml:"link"`
  }

  channel struct {
    XMLName xml.Name `xml:"channel"`
    Title string `xml:"title"`
    Description string `xml:"description"`
    Link string `xml:"link"`
    PubDate string `xml:"pubDate"`
    LastBuildDate string `xml:"lastBuildDate"`
    TTL string `xml:"ttl"`
    Language string `xml:"language"`
    ManagingEditor string `xml:"managingEditor"`
    WebMaster string `xml:"webMaster"`
    Image image `xml:"image"`
    Item []item `xml:"item"`
  }

  rssDocument struct {
    XMLName xml.Name `xml:"rss"`
    Channel channel `xml:"channel"`
  }

)

// rssMatcher implements the Matcher interface
type rssMatcher struct {}

// init registers the matcher with the program
func init() {
  var matcher rssMatcher
  search.Register("rss", matcher)
}

func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
  if feed.URI == "" {
    return nil, errors.New("No rss feed URI provided")
  }

  // Retrieve the rss feed document
  resp, err := http.Get(feed.URI)
  if err != nil {
    return nil, err
  }

  // close the response once we return from the function
  defer resp.Body.Close()

  // Check the status code to show proper response
  if resp.StatusCode != 200 {
    return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
  }

  // Decode the rss feed document
  var document rssDocument
  err = xml.NewDecoder(resp.Body).Decode(&document)
  return &document, err
}

/*
Search searches a feed for term and return a []*search.Result and an error
It uses a value receiver to implement the Matcher interface
*/
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
  var results []*search.Result

  log.Printf("Search Feed Type[%s] Site[%s] for Url[%s]\n", 
  feed.Type, feed.Name, feed.URI)

  // retrieve the data to search
  document, err := m.retrieve(feed)
  if err != nil {
    return nil, err
  }

  for _, channelItem := range document.Channel.Item {
    // Check the title for search term
    matched, err := regexp.MatchString(searchTerm, channelItem.Title)

    if err != nil {
       return nil, err 
    }

    // We found a match
    if matched {
      results = append(results, &search.Result{
         Field: "Title",
         Content: channelItem.Title,
      })
    }

    // Check the description for the search term
    matchedDesc, err := regexp.MatchString(searchTerm, channelItem.Description)

    if err != nil {
       return nil, err 
    }

    // We found a match
    // TODO: Move this to a pkg/utils/regex.go
    if matchedDesc {
      results = append(results, &search.Result{
         Field: "Description",
         Content: channelItem.Description,
      })
    }

  }
    return results, nil
}

