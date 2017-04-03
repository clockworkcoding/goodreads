package goodreads

import (
	"encoding/xml"
	"log"
)

func (c *Client) QueryUser() (user User_user, err error) {

	client, err := c.GetHttpClient()
	if err != nil {
		return
	}

	resp, err := client.Get(
		"https://www.goodreads.com/api/auth_user")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var response User_GoodreadsResponse
	if err := xml.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println(err)
	}
	return response.User_user, nil
}

type User_GoodreadsResponse struct {
	User_Request User_Request `xml:" Request,omitempty" json:"Request,omitempty"`
	User_user    User_user    `xml:" user,omitempty" json:"user,omitempty"`
}

type User_Request struct {
	User_authentication User_authentication `xml:" authentication,omitempty" json:"authentication,omitempty"`
	User_key            User_key            `xml:" key,omitempty" json:"key,omitempty"`
	User_method         User_method         `xml:" method,omitempty" json:"method,omitempty"`
}

type User_authentication struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type User_key struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type User_link struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type User_method struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type User_name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type User_root struct {
	User_GoodreadsResponse User_GoodreadsResponse `xml:" GoodreadsResponse,omitempty" json:"GoodreadsResponse,omitempty"`
}

type User_user struct {
	Attr_id   string    `xml:" id,attr"  json:",omitempty"`
	User_link User_link `xml:" link,omitempty" json:"link,omitempty"`
	User_name User_name `xml:" name,omitempty" json:"name,omitempty"`
}
