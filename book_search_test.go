package goodreads

import "testing"

func TestGetSearch(t *testing.T) {
	config := readConfig()
	gr := NewClient(config.GoodReadsKey, config.GoodReadsSecret)
	if response, _ := gr.GetSearch("The Name of the Wind"); response.Search_work[0].Search_best_book.Search_title.Text != "The Name of the Wind (The Kingkiller Chronicle, #1)" {
		t.Fail()
	}
}

func TestBadKey(t *testing.T) {
	config := readConfig()
	gr := NewClient("fake key", config.GoodReadsSecret)
	if _, err := gr.GetSearch("Collapsing Empire"); err == nil {
		t.Fail()
	}
}
