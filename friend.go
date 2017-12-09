package goodreads

import (
	"errors"
	"fmt"
	"log"
	"net/url"
)

//AddFriend sends a friend request to the selected userId
func (c *Client) AddFriend(userid string) (err error) {
	postURL := fmt.Sprintf("%sfriend/add_as_friend.xml", apiRoot)
	form := url.Values{}
	form.Add("id", userid)

	client, err := c.GetHttpClient()
	if err != nil {
		return
	}

	resp, err := client.PostForm(postURL, form)
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
