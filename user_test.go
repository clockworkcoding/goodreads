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
		fmt.Println(xml)
	}
}
