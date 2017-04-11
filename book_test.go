package goodreads

import "testing"

func TestBookLookup(t *testing.T) {
	config := readConfig()
	gr := NewClient(config.GoodReadsKey, config.GoodReadsSecret)
	if book, _ := gr.GetBook("30078567"); book.Book_title[0].Text != "The Collapsing Empire (The Interdependency #1)" {

		t.Fail()
	}
}

func TestBadBookLookup(t *testing.T) {
	config := readConfig()
	gr := NewClient(config.GoodReadsKey, config.GoodReadsSecret)
	if _, err := gr.GetBook("300785673"); err == nil {

		t.Fail()
	}
}
