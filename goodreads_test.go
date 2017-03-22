package goodreads

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGetSearch(t *testing.T) {
	config := readConfig()
	if response, _ := GetSearch("The Name of the Wind", config.GoodReadsKey); response.Search_work[0].Search_best_book.Search_title.Text != "The Name of the Wind (The Kingkiller Chronicle, #1)" {
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
