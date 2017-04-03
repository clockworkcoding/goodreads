package goodreads

import (
	"fmt"
	"testing"
)

func TestGetShelf(t *testing.T) {
	config := readConfig()
	c := NewClientWithToken(config.GoodReadsKey, config.GoodReadsSecret, config.SlackToken, config.GoodReadsHost)

	if response, err := c.GetShelf("22966785", ReviewListParameters{Shelf: "read-next"}); err != nil && response.Reviews_review[0].Reviews_book.Reviews_title.Text == "The Collapsing Empire (The Interdependency #1)" {
		fmt.Println("Get Shelf: " + err.Error())
		t.Fail()
	} else {

		fmt.Println(response)
	}
}
