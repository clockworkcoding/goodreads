package goodreads

type Book_GoodreadsResponse struct {
	Book_Request Book_Request `xml:" Request,omitempty" json:"Request,omitempty"`
	Book_book    []Book_book  `xml:" book,omitempty" json:"book,omitempty"`
}

type Book_Request struct {
	Book_authentication Book_authentication `xml:" authentication,omitempty" json:"authentication,omitempty"`
	Book_key            Book_key            `xml:" key,omitempty" json:"key,omitempty"`
	Book_method         Book_method         `xml:" method,omitempty" json:"method,omitempty"`
}

type Book_asin struct {
}

type Book_authentication struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_author struct {
	Book_average_rating     []Book_average_rating   `xml:" average_rating,omitempty" json:"average_rating,omitempty"`
	Book_id                 []Book_id               `xml:" id,omitempty" json:"id,omitempty"`
	Book_image_url          []Book_image_url        `xml:" image_url,omitempty" json:"image_url,omitempty"`
	Book_link               []Book_link             `xml:" link,omitempty" json:"link,omitempty"`
	Book_name               Book_name               `xml:" name,omitempty" json:"name,omitempty"`
	Book_ratings_count      []Book_ratings_count    `xml:" ratings_count,omitempty" json:"ratings_count,omitempty"`
	Book_role               Book_role               `xml:" role,omitempty" json:"role,omitempty"`
	Book_small_image_url    []Book_small_image_url  `xml:" small_image_url,omitempty" json:"small_image_url,omitempty"`
	Book_text_reviews_count Book_text_reviews_count `xml:" text_reviews_count,omitempty" json:"text_reviews_count,omitempty"`
}

type Book_authors struct {
	Book_author []Book_author `xml:" author,omitempty" json:"author,omitempty"`
}

type Book_average_rating struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_best_book_id struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_book struct {
	Book_asin                 Book_asin                 `xml:" asin,omitempty" json:"asin,omitempty"`
	Book_authors              []Book_authors            `xml:" authors,omitempty" json:"authors,omitempty"`
	Book_average_rating       []Book_average_rating     `xml:" average_rating,omitempty" json:"average_rating,omitempty"`
	Book_book_links           Book_book_links           `xml:" book_links,omitempty" json:"book_links,omitempty"`
	Book_buy_links            Book_buy_links            `xml:" buy_links,omitempty" json:"buy_links,omitempty"`
	Book_country_code         Book_country_code         `xml:" country_code,omitempty" json:"country_code,omitempty"`
	Book_description          Book_description          `xml:" description,omitempty" json:"description,omitempty"`
	Book_edition_information  Book_edition_information  `xml:" edition_information,omitempty" json:"edition_information,omitempty"`
	Book_format               Book_format               `xml:" format,omitempty" json:"format,omitempty"`
	Book_id                   []Book_id                 `xml:" id,omitempty" json:"id,omitempty"`
	Book_image_url            []Book_image_url          `xml:" image_url,omitempty" json:"image_url,omitempty"`
	Book_is_ebook             Book_is_ebook             `xml:" is_ebook,omitempty" json:"is_ebook,omitempty"`
	Book_isbn                 []Book_isbn               `xml:" isbn,omitempty" json:"isbn,omitempty"`
	Book_isbn13               []Book_isbn13             `xml:" isbn13,omitempty" json:"isbn13,omitempty"`
	Book_kindle_asin          Book_kindle_asin          `xml:" kindle_asin,omitempty" json:"kindle_asin,omitempty"`
	Book_language_code        Book_language_code        `xml:" language_code,omitempty" json:"language_code,omitempty"`
	Book_link                 []Book_link               `xml:" link,omitempty" json:"link,omitempty"`
	Book_marketplace_id       Book_marketplace_id       `xml:" marketplace_id,omitempty" json:"marketplace_id,omitempty"`
	Book_num_pages            []Book_num_pages          `xml:" num_pages,omitempty" json:"num_pages,omitempty"`
	Book_popular_shelves      Book_popular_shelves      `xml:" popular_shelves,omitempty" json:"popular_shelves,omitempty"`
	Book_publication_day      []Book_publication_day    `xml:" publication_day,omitempty" json:"publication_day,omitempty"`
	Book_publication_month    []Book_publication_month  `xml:" publication_month,omitempty" json:"publication_month,omitempty"`
	Book_publication_year     []Book_publication_year   `xml:" publication_year,omitempty" json:"publication_year,omitempty"`
	Book_publisher            Book_publisher            `xml:" publisher,omitempty" json:"publisher,omitempty"`
	Book_ratings_count        []Book_ratings_count      `xml:" ratings_count,omitempty" json:"ratings_count,omitempty"`
	Book_reviews_widget       Book_reviews_widget       `xml:" reviews_widget,omitempty" json:"reviews_widget,omitempty"`
	Book_series_works         Book_series_works         `xml:" series_works,omitempty" json:"series_works,omitempty"`
	Book_similar_books        Book_similar_books        `xml:" similar_books,omitempty" json:"similar_books,omitempty"`
	Book_small_image_url      []Book_small_image_url    `xml:" small_image_url,omitempty" json:"small_image_url,omitempty"`
	Book_text_reviews_count   Book_text_reviews_count   `xml:" text_reviews_count,omitempty" json:"text_reviews_count,omitempty"`
	Book_title                []Book_title              `xml:" title,omitempty" json:"title,omitempty"`
	Book_title_without_series Book_title_without_series `xml:" title_without_series,omitempty" json:"title_without_series,omitempty"`
	Book_url                  Book_url                  `xml:" url,omitempty" json:"url,omitempty"`
	Book_work                 []Book_work               `xml:" work,omitempty" json:"work,omitempty"`
}

type Book_book_link struct {
	Book_id   []Book_id   `xml:" id,omitempty" json:"id,omitempty"`
	Book_link []Book_link `xml:" link,omitempty" json:"link,omitempty"`
	Book_name Book_name   `xml:" name,omitempty" json:"name,omitempty"`
}

type Book_book_links struct {
	Book_book_link Book_book_link `xml:" book_link,omitempty" json:"book_link,omitempty"`
}

type Book_books_count struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_buy_link struct {
	Book_id   []Book_id   `xml:" id,omitempty" json:"id,omitempty"`
	Book_link []Book_link `xml:" link,omitempty" json:"link,omitempty"`
	Book_name Book_name   `xml:" name,omitempty" json:"name,omitempty"`
}

type Book_buy_links struct {
	Book_buy_link []Book_buy_link `xml:" buy_link,omitempty" json:"buy_link,omitempty"`
}

type Book_country_code struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_default_chaptering_book_id struct {
	Attr_nil  string `xml:" nil,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
}

type Book_default_description_language_code struct {
	Attr_nil string `xml:" nil,attr"  json:",omitempty"`
}

type Book_desc_user_id struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_edition_information struct {
}

type Book_format struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_id struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_image_url struct {
	Attr_nophoto string `xml:" nophoto,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

type Book_is_ebook struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_isbn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_isbn13 struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_key struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_kindle_asin struct {
}

type Book_language_code struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_link struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_marketplace_id struct {
}

type Book_media_type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_method struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_note struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_num_pages struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_numbered struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_original_language_id struct {
	Attr_nil  string `xml:" nil,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
}

type Book_original_publication_day struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_original_publication_month struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_original_publication_year struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_original_title struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_popular_shelves struct {
	Book_shelf []Book_shelf `xml:" shelf,omitempty" json:"shelf,omitempty"`
}

type Book_primary_work_count struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_publication_day struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_publication_month struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_publication_year struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_publisher struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_rating_dist struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_ratings_count struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_ratings_sum struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_reviews_count struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_reviews_widget struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_role struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_root struct {
	Book_GoodreadsResponse Book_GoodreadsResponse `xml:" GoodreadsResponse,omitempty" json:"GoodreadsResponse,omitempty"`
}

type Book_series struct {
	Book_description        Book_description        `xml:" description,omitempty" json:"description,omitempty"`
	Book_id                 []Book_id               `xml:" id,omitempty" json:"id,omitempty"`
	Book_note               Book_note               `xml:" note,omitempty" json:"note,omitempty"`
	Book_numbered           Book_numbered           `xml:" numbered,omitempty" json:"numbered,omitempty"`
	Book_primary_work_count Book_primary_work_count `xml:" primary_work_count,omitempty" json:"primary_work_count,omitempty"`
	Book_series_works_count Book_series_works_count `xml:" series_works_count,omitempty" json:"series_works_count,omitempty"`
	Book_title              []Book_title            `xml:" title,omitempty" json:"title,omitempty"`
}

type Book_series_work struct {
	Book_id            []Book_id          `xml:" id,omitempty" json:"id,omitempty"`
	Book_series        Book_series        `xml:" series,omitempty" json:"series,omitempty"`
	Book_user_position Book_user_position `xml:" user_position,omitempty" json:"user_position,omitempty"`
}

type Book_series_works struct {
	Book_series_work []Book_series_work `xml:" series_work,omitempty" json:"series_work,omitempty"`
}

type Book_series_works_count struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_shelf struct {
	Attr_count string `xml:" count,attr"  json:",omitempty"`
	Attr_name  string `xml:" name,attr"  json:",omitempty"`
}

type Book_similar_books struct {
	Book_book []Book_book `xml:" book,omitempty" json:"book,omitempty"`
}

type Book_small_image_url struct {
	Attr_nophoto string `xml:" nophoto,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

type Book_text_reviews_count struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Book_title struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_title_without_series struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_url struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_user_position struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Book_work struct {
	Book_best_book_id                      Book_best_book_id                      `xml:" best_book_id,omitempty" json:"best_book_id,omitempty"`
	Book_books_count                       Book_books_count                       `xml:" books_count,omitempty" json:"books_count,omitempty"`
	Book_default_chaptering_book_id        Book_default_chaptering_book_id        `xml:" default_chaptering_book_id,omitempty" json:"default_chaptering_book_id,omitempty"`
	Book_default_description_language_code Book_default_description_language_code `xml:" default_description_language_code,omitempty" json:"default_description_language_code,omitempty"`
	Book_desc_user_id                      Book_desc_user_id                      `xml:" desc_user_id,omitempty" json:"desc_user_id,omitempty"`
	Book_id                                []Book_id                              `xml:" id,omitempty" json:"id,omitempty"`
	Book_media_type                        Book_media_type                        `xml:" media_type,omitempty" json:"media_type,omitempty"`
	Book_original_language_id              Book_original_language_id              `xml:" original_language_id,omitempty" json:"original_language_id,omitempty"`
	Book_original_publication_day          Book_original_publication_day          `xml:" original_publication_day,omitempty" json:"original_publication_day,omitempty"`
	Book_original_publication_month        Book_original_publication_month        `xml:" original_publication_month,omitempty" json:"original_publication_month,omitempty"`
	Book_original_publication_year         Book_original_publication_year         `xml:" original_publication_year,omitempty" json:"original_publication_year,omitempty"`
	Book_original_title                    Book_original_title                    `xml:" original_title,omitempty" json:"original_title,omitempty"`
	Book_rating_dist                       Book_rating_dist                       `xml:" rating_dist,omitempty" json:"rating_dist,omitempty"`
	Book_ratings_count                     []Book_ratings_count                   `xml:" ratings_count,omitempty" json:"ratings_count,omitempty"`
	Book_ratings_sum                       Book_ratings_sum                       `xml:" ratings_sum,omitempty" json:"ratings_sum,omitempty"`
	Book_reviews_count                     Book_reviews_count                     `xml:" reviews_count,omitempty" json:"reviews_count,omitempty"`
	Book_text_reviews_count                Book_text_reviews_count                `xml:" text_reviews_count,omitempty" json:"text_reviews_count,omitempty"`
}
