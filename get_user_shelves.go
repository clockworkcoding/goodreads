package goodreads

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (c Client) GetUserShelves(id string) (shelves Shelf_shelves, err error) {
	var response Shelf_GoodreadsResponse

	url := apiRoot + fmt.Sprintf("shelf/list.xml?key=%s&user_id=%s", c.consumerKey, id)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("NewRequest: ", err)
		return
	}

	client, err := c.GetHttpClient()
	if err != nil {
		log.Println("Do: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}

	if err = xml.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println(err)
		return
	}

	return response.Shelf_shelves, nil
}

type Shelf_GoodreadsResponse struct {
	Shelf_Request Shelf_Request `xml:" Request,omitempty" json:"Request,omitempty"`
	Shelf_shelves Shelf_shelves `xml:" shelves,omitempty" json:"shelves,omitempty"`
}

type Shelf_Request struct {
	Shelf_authentication Shelf_authentication `xml:" authentication,omitempty" json:"authentication,omitempty"`
	Shelf_key            Shelf_key            `xml:" key,omitempty" json:"key,omitempty"`
	Shelf_method         Shelf_method         `xml:" method,omitempty" json:"method,omitempty"`
}

type Shelf_authentication struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Shelf_book_count struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Shelf_description struct {
	Attr_nil string `xml:" nil,attr"  json:",omitempty"`
}

type Shelf_display_fields struct {
}

type Shelf_exclusive_flag struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Shelf_featured struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Shelf_id struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Shelf_key struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Shelf_method struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Shelf_name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Shelf_order struct {
	Attr_nil string `xml:" nil,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

type Shelf_per_page struct {
	Attr_nil  string `xml:" nil,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
}

type Shelf_recommend_for struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Shelf_root struct {
	Shelf_GoodreadsResponse Shelf_GoodreadsResponse `xml:" GoodreadsResponse,omitempty" json:"GoodreadsResponse,omitempty"`
}

type Shelf_shelves struct {
	Attr_end         string             `xml:" end,attr"  json:",omitempty"`
	Attr_start       string             `xml:" start,attr"  json:",omitempty"`
	Attr_total       string             `xml:" total,attr"  json:",omitempty"`
	Shelf_user_shelf []Shelf_user_shelf `xml:" user_shelf,omitempty" json:"user_shelf,omitempty"`
}

type Shelf_sort struct {
	Attr_nil string `xml:" nil,attr"  json:",omitempty"`
}

type Shelf_sticky struct {
	Attr_nil  string `xml:" nil,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
}

type Shelf_user_shelf struct {
	Shelf_book_count     Shelf_book_count     `xml:" book_count,omitempty" json:"book_count,omitempty"`
	Shelf_description    Shelf_description    `xml:" description,omitempty" json:"description,omitempty"`
	Shelf_display_fields Shelf_display_fields `xml:" display_fields,omitempty" json:"display_fields,omitempty"`
	Shelf_exclusive_flag Shelf_exclusive_flag `xml:" exclusive_flag,omitempty" json:"exclusive_flag,omitempty"`
	Shelf_featured       Shelf_featured       `xml:" featured,omitempty" json:"featured,omitempty"`
	Shelf_id             Shelf_id             `xml:" id,omitempty" json:"id,omitempty"`
	Shelf_name           Shelf_name           `xml:" name,omitempty" json:"name,omitempty"`
	Shelf_order          Shelf_order          `xml:" order,omitempty" json:"order,omitempty"`
	Shelf_per_page       Shelf_per_page       `xml:" per_page,omitempty" json:"per_page,omitempty"`
	Shelf_recommend_for  Shelf_recommend_for  `xml:" recommend_for,omitempty" json:"recommend_for,omitempty"`
	Shelf_sort           Shelf_sort           `xml:" sort,omitempty" json:"sort,omitempty"`
	Shelf_sticky         Shelf_sticky         `xml:" sticky,omitempty" json:"sticky,omitempty"`
}
