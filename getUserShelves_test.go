package goodreads

import (
	"fmt"
	"testing"
)

func TestGetUserShelves(t *testing.T) {
	config := readConfig()
	gr := NewClient(config.GoodReadsKey, config.GoodReadsSecret)
	if response, _ := gr.GetUserShelves("22966785"); response.Shelf_user_shelf[3].Shelf_name.Text != "given-up-reading" {
		fmt.Println(response.Shelf_user_shelf[3].Shelf_name.Text)
		t.Fail()
	} else {

		fmt.Println(response.Shelf_user_shelf[3].Shelf_name.Text)
	}
}
