package goodreads

import (
	"fmt"
	"testing"
)

func TestGetShelf(t *testing.T) {
	config := readConfig()
	c := NewClientWithToken(config.GoodReadsKey, config.GoodReadsSecret, config.UserToken, config.UserSecret)

	if response, err := c.GetShelf("22966785", ReviewListParameters{Shelf: "read-next"}); err != nil && response.Reviews_review[0].Reviews_book.Reviews_title.Text == "The Collapsing Empire (The Interdependency #1)" {
		fmt.Println("Get Shelf: " + err.Error())
		t.Fail()
	} else {

		fmt.Println(response)
	}
}

func TestGetShelfWithPages(t *testing.T) {
	config := readConfig()
	c := NewClientWithToken(config.GoodReadsKey, config.GoodReadsSecret, config.UserToken, config.UserSecret)

	if _, err := c.GetShelf("22966785", ReviewListParameters{Shelf: "read", Page: 2, PerPage: 50}); err != nil {
		fmt.Println("Get Shelf: " + err.Error())
		t.Fail()
	} else {
		fmt.Println("Get Shelf")
	}
}
