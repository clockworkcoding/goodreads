package goodreads

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var (
	apiRoot = "http://www.goodreads.com/"
)

func GetSearch(query, key string) Search_GoodreadsResponse {
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	safeQuery := url.QueryEscape(query)
	safeKey := url.QueryEscape(key)

	var model Search_GoodreadsResponse

	url := apiRoot + fmt.Sprintf("search/index.xml?key=%s&q=%s", safeKey, safeQuery)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return model
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return model
	}
	defer resp.Body.Close()

	if err := xml.NewDecoder(resp.Body).Decode(&model); err != nil {
		log.Println(err)
	}
	fmt.Println(model.Search_search.Search_results.Search_work[0].Search_best_book.Search_title)
	return model
}
