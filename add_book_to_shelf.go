package goodreads

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//AddBookToShelf Adds a book to the given shelf. If not one of the defaults, it will also be added to "read"
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

	if resp.StatusCode != 201 {
		err = errors.New(resp.Status)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	return nil
}
