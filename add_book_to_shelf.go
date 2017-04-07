package goodreads

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) AddBookToShelf(bookid, shelfName string) (err error) {

	postURL := fmt.Sprintf("%sshelf/add_to_shelf.xml", apiRoot)
	form := url.Values{}
	form.Add("name", shelfName)
	form.Add("book_id", bookid)

	// Build the request
	req, err := http.NewRequest("POST", postURL, strings.NewReader(form.Encode()))
	if err != nil {
		log.Println("NewRequest: ", err)
		return
	}

	client, err := c.GetHttpClient()
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	return nil
}
