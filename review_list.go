package goodreads

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const ()

type ReviewListParameters struct {
	Shelf   string
	Sort    string
	Search  string
	Order   string
	Page    int
	PerPage int
}

func (c Client) GetShelf(userID string, params ReviewListParameters) (response Reviews_reviews, err error) {
	return c.ReviewList(userID, params)
}

func (c Client) ReviewList(userID string, params ReviewListParameters) (response Reviews_reviews, err error) {

	form := url.Values{}
	form.Add("v", "2")
	form.Add("id", userID)
	form.Add("key", c.consumerKey)
	if params.Shelf != "" {
		form.Add("shelf", params.Shelf)
	}
	if params.Sort != "" {
		form.Add("sort", params.Sort)
	}
	if params.Search != "" {
		form.Add("search", params.Search)
	}
	if params.Order != "" {
		form.Add("order", params.Order)
	}
	if params.Page > 0 {
		form.Add("Page", string(params.Page))
	}
	if params.PerPage > 0 {
		form.Add("per_page", string(params.PerPage))
	}

	apiURL := fmt.Sprintf("%sreview/list.xml?%s", apiRoot, form.Encode())
	fmt.Println(apiURL)
	// Build the request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Println("NewRequest: ", err)
		return
	}

	client, err := c.GetHttpClient()
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return response, errors.New(resp.Status)
	}
	var goodreadsResponse Reviews_GoodreadsResponse
	if err := xml.NewDecoder(resp.Body).Decode(&goodreadsResponse); err != nil {
		log.Println(err)
	}

	return goodreadsResponse.Reviews_reviews, nil
}

type Reviews_GoodreadsResponse struct {
	Reviews_Request Reviews_Request `xml:" Request,omitempty" json:"Request,omitempty"`
	Reviews_reviews Reviews_reviews `xml:" reviews,omitempty" json:"reviews,omitempty"`
}

type Reviews_Request struct {
	Reviews_authentication Reviews_authentication `xml:" authentication,omitempty" json:"authentication,omitempty"`
	Reviews_key            Reviews_key            `xml:" key,omitempty" json:"key,omitempty"`
	Reviews_method         Reviews_method         `xml:" method,omitempty" json:"method,omitempty"`
}

type Reviews_a struct {
	Attr_href string `xml:" href,attr"  json:",omitempty"`
	Attr_rel  string `xml:" rel,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Reviews_authentication struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_author struct {
	Reviews_average_rating     Reviews_average_rating     `xml:" average_rating,omitempty" json:"average_rating,omitempty"`
	Reviews_id                 Reviews_id                 `xml:" id,omitempty" json:"id,omitempty"`
	Reviews_image_url          Reviews_image_url          `xml:" image_url,omitempty" json:"image_url,omitempty"`
	Reviews_link               Reviews_link               `xml:" link,omitempty" json:"link,omitempty"`
	Reviews_name               Reviews_name               `xml:" name,omitempty" json:"name,omitempty"`
	Reviews_ratings_count      Reviews_ratings_count      `xml:" ratings_count,omitempty" json:"ratings_count,omitempty"`
	Reviews_role               Reviews_role               `xml:" role,omitempty" json:"role,omitempty"`
	Reviews_small_image_url    Reviews_small_image_url    `xml:" small_image_url,omitempty" json:"small_image_url,omitempty"`
	Reviews_text_reviews_count Reviews_text_reviews_count `xml:" text_reviews_count,omitempty" json:"text_reviews_count,omitempty"`
}

type Reviews_authors struct {
	Reviews_author Reviews_author `xml:" author,omitempty" json:"author,omitempty"`
}

type Reviews_average_rating struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_b struct {
	Reviews_br []Reviews_br `xml:" br,omitempty" json:"br,omitempty"`
	Reviews_i  []Reviews_i  `xml:" i,omitempty" json:"i,omitempty"`
	Text       string       `xml:",chardata" json:",omitempty"`
}

type Reviews_body struct {
}

type Reviews_book struct {
	Reviews_authors              Reviews_authors              `xml:" authors,omitempty" json:"authors,omitempty"`
	Reviews_average_rating       Reviews_average_rating       `xml:" average_rating,omitempty" json:"average_rating,omitempty"`
	Reviews_description          Reviews_description          `xml:" description,omitempty" json:"description,omitempty"`
	Reviews_edition_information  Reviews_edition_information  `xml:" edition_information,omitempty" json:"edition_information,omitempty"`
	Reviews_format               Reviews_format               `xml:" format,omitempty" json:"format,omitempty"`
	Reviews_id                   Reviews_id                   `xml:" id,omitempty" json:"id,omitempty"`
	Reviews_image_url            Reviews_image_url            `xml:" image_url,omitempty" json:"image_url,omitempty"`
	Reviews_isbn                 Reviews_isbn                 `xml:" isbn,omitempty" json:"isbn,omitempty"`
	Reviews_isbn13               Reviews_isbn13               `xml:" isbn13,omitempty" json:"isbn13,omitempty"`
	Reviews_large_image_url      Reviews_large_image_url      `xml:" large_image_url,omitempty" json:"large_image_url,omitempty"`
	Reviews_link                 Reviews_link                 `xml:" link,omitempty" json:"link,omitempty"`
	Reviews_num_pages            Reviews_num_pages            `xml:" num_pages,omitempty" json:"num_pages,omitempty"`
	Reviews_publication_day      Reviews_publication_day      `xml:" publication_day,omitempty" json:"publication_day,omitempty"`
	Reviews_publication_month    Reviews_publication_month    `xml:" publication_month,omitempty" json:"publication_month,omitempty"`
	Reviews_publication_year     Reviews_publication_year     `xml:" publication_year,omitempty" json:"publication_year,omitempty"`
	Reviews_published            Reviews_published            `xml:" published,omitempty" json:"published,omitempty"`
	Reviews_publisher            Reviews_publisher            `xml:" publisher,omitempty" json:"publisher,omitempty"`
	Reviews_ratings_count        Reviews_ratings_count        `xml:" ratings_count,omitempty" json:"ratings_count,omitempty"`
	Reviews_small_image_url      Reviews_small_image_url      `xml:" small_image_url,omitempty" json:"small_image_url,omitempty"`
	Reviews_text_reviews_count   Reviews_text_reviews_count   `xml:" text_reviews_count,omitempty" json:"text_reviews_count,omitempty"`
	Reviews_title                Reviews_title                `xml:" title,omitempty" json:"title,omitempty"`
	Reviews_title_without_series Reviews_title_without_series `xml:" title_without_series,omitempty" json:"title_without_series,omitempty"`
}

type Reviews_br struct {
}

type Reviews_comments_count struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_date_added struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_date_updated struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_description struct {
	Reviews_a  Reviews_a    `xml:" a,omitempty" json:"a,omitempty"`
	Reviews_b  []Reviews_b  `xml:" b,omitempty" json:"b,omitempty"`
	Reviews_br []Reviews_br `xml:" br,omitempty" json:"br,omitempty"`
	Reviews_i  []Reviews_i  `xml:" i,omitempty" json:"i,omitempty"`
	Reviews_p  []Reviews_p  `xml:" p,omitempty" json:"p,omitempty"`
	Text       string       `xml:",chardata" json:",omitempty"`
}

type Reviews_edition_information struct {
}

type Reviews_em struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_format struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_i struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_id struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Reviews_image_url struct {
	Attr_nophoto string `xml:" nophoto,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

type Reviews_isbn struct {
	Attr_nil string `xml:" nil,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

type Reviews_isbn13 struct {
	Attr_nil string `xml:" nil,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

type Reviews_key struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_large_image_url struct {
}

type Reviews_link struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_method struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_num_pages struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_owned struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_p struct {
	Reviews_em Reviews_em `xml:" em,omitempty" json:"em,omitempty"`
	Text       string     `xml:",chardata" json:",omitempty"`
}

type Reviews_publication_day struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_publication_month struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_publication_year struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_published struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_publisher struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_rating struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_ratings_count struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_read_at struct {
}

type Reviews_read_count struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_recommended_by struct {
}

type Reviews_recommended_for struct {
}

type Reviews_review struct {
	Reviews_body            Reviews_body            `xml:" body,omitempty" json:"body,omitempty"`
	Reviews_book            Reviews_book            `xml:" book,omitempty" json:"book,omitempty"`
	Reviews_comments_count  Reviews_comments_count  `xml:" comments_count,omitempty" json:"comments_count,omitempty"`
	Reviews_date_added      Reviews_date_added      `xml:" date_added,omitempty" json:"date_added,omitempty"`
	Reviews_date_updated    Reviews_date_updated    `xml:" date_updated,omitempty" json:"date_updated,omitempty"`
	Reviews_id              Reviews_id              `xml:" id,omitempty" json:"id,omitempty"`
	Reviews_link            Reviews_link            `xml:" link,omitempty" json:"link,omitempty"`
	Reviews_owned           Reviews_owned           `xml:" owned,omitempty" json:"owned,omitempty"`
	Reviews_rating          Reviews_rating          `xml:" rating,omitempty" json:"rating,omitempty"`
	Reviews_read_at         Reviews_read_at         `xml:" read_at,omitempty" json:"read_at,omitempty"`
	Reviews_read_count      Reviews_read_count      `xml:" read_count,omitempty" json:"read_count,omitempty"`
	Reviews_recommended_by  Reviews_recommended_by  `xml:" recommended_by,omitempty" json:"recommended_by,omitempty"`
	Reviews_recommended_for Reviews_recommended_for `xml:" recommended_for,omitempty" json:"recommended_for,omitempty"`
	Reviews_shelves         Reviews_shelves         `xml:" shelves,omitempty" json:"shelves,omitempty"`
	Reviews_spoiler_flag    Reviews_spoiler_flag    `xml:" spoiler_flag,omitempty" json:"spoiler_flag,omitempty"`
	Reviews_spoilers_state  Reviews_spoilers_state  `xml:" spoilers_state,omitempty" json:"spoilers_state,omitempty"`
	Reviews_started_at      Reviews_started_at      `xml:" started_at,omitempty" json:"started_at,omitempty"`
	Reviews_url             Reviews_url             `xml:" url,omitempty" json:"url,omitempty"`
	Reviews_votes           Reviews_votes           `xml:" votes,omitempty" json:"votes,omitempty"`
}

type Reviews_reviews struct {
	Attr_end       string           `xml:" end,attr"  json:",omitempty"`
	Attr_start     string           `xml:" start,attr"  json:",omitempty"`
	Attr_total     string           `xml:" total,attr"  json:",omitempty"`
	Reviews_review []Reviews_review `xml:" review,omitempty" json:"review,omitempty"`
}

type Reviews_role struct {
}

type Reviews_root struct {
	Reviews_GoodreadsResponse Reviews_GoodreadsResponse `xml:" GoodreadsResponse,omitempty" json:"GoodreadsResponse,omitempty"`
}

type Reviews_shelf struct {
	Attr_exclusive       string `xml:" exclusive,attr"  json:",omitempty"`
	Attr_name            string `xml:" name,attr"  json:",omitempty"`
	Attr_review_shelf_id string `xml:" review_shelf_id,attr"  json:",omitempty"`
	Attr_sortable        string `xml:" sortable,attr"  json:",omitempty"`
}

type Reviews_shelves struct {
	Reviews_shelf []Reviews_shelf `xml:" shelf,omitempty" json:"shelf,omitempty"`
}

type Reviews_small_image_url struct {
	Attr_nophoto string `xml:" nophoto,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

type Reviews_spoiler_flag struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_spoilers_state struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_started_at struct {
}

type Reviews_text_reviews_count struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Reviews_title struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_title_without_series struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_url struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reviews_votes struct {
	Text string `xml:",chardata" json:",omitempty"`
}
