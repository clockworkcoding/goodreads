package goodreads

import (
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	config := readConfig()
	c := NewClientWithToken(config.GoodReadsKey, config.GoodReadsSecret, config.SlackToken, config.GoodReadsHost)
	if xml, err := c.QueryUser(); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	} else {
		fmt.Println("*************************************************************\n", xml.User_link.Text, xml.Attr_id, xml.User_name.Text)
	}
}

func TestAddToShelf(t *testing.T) {
	config := readConfig()
	c := NewClientWithToken(config.GoodReadsKey, config.GoodReadsSecret, config.SlackToken, config.GoodReadsHost)
	if err := c.AddBookToShelf("645180", "read"); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}
