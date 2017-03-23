package goodreads

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var (
	apiRoot = "http://www.goodreads.com/"
)

func (c *Client) GetSearch(query string) (Search_results, error) {
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	safeQuery := url.QueryEscape(query)
	safeKey := url.QueryEscape(c.consumer.ConsumerKey)

	var model Search_GoodreadsResponse

	url := apiRoot + fmt.Sprintf("search/index.xml?key=%s&q=%s", safeKey, safeQuery)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return model.Search_search.Search_results, err
	}

	client := c.GetHttpClient()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return model.Search_search.Search_results, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return model.Search_search.Search_results, errors.New(resp.Status)
	}

	if err := xml.NewDecoder(resp.Body).Decode(&model); err != nil {
		log.Println(err)
	}

	return model.Search_search.Search_results, nil
}
