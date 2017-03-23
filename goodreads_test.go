package goodreads

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

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

type Configuration struct {
	GoodReadsKey    string `json:"goodReadsKey"`
	GoodReadsSecret string `json:"goodReadsSecret"`
	SlackToken      string `json:"slackToken"`
	GoodReadsHost   string `json:"goodReadsHost"`
	GoodReadsPort   string `json:"goodReadsPort"`
	SlackHost       string `json:"slackHost"`
	SlackHostHTTP   string `json:"slackHostHttp"`
	IsHTTPS         string `json:"isHTTPS"`
}

func readConfig() Configuration {

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
