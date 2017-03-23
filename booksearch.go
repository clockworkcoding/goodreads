package goodreads

type Search_GoodreadsResponse struct {
	Search_Request Search_Request `xml:" Request,omitempty" json:"Request,omitempty"`
	Search_search  Search_search  `xml:" search,omitempty" json:"search,omitempty"`
}

type Search_Request struct {
	Search_authentication Search_authentication `xml:" authentication,omitempty" json:"authentication,omitempty"`
	Search_key            Search_key            `xml:" key,omitempty" json:"key,omitempty"`
	Search_method         Search_method         `xml:" method,omitempty" json:"method,omitempty"`
}

type Search_authentication struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_author struct {
	Search_id   Search_id   `xml:" id,omitempty" json:"id,omitempty"`
	Search_name Search_name `xml:" name,omitempty" json:"name,omitempty"`
}

type Search_average_rating struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_best_book struct {
	Attr_type              string                 `xml:" type,attr"  json:",omitempty"`
	Search_author          Search_author          `xml:" author,omitempty" json:"author,omitempty"`
	Search_id              Search_id              `xml:" id,omitempty" json:"id,omitempty"`
	Search_image_url       Search_image_url       `xml:" image_url,omitempty" json:"image_url,omitempty"`
	Search_small_image_url Search_small_image_url `xml:" small_image_url,omitempty" json:"small_image_url,omitempty"`
	Search_title           Search_title           `xml:" title,omitempty" json:"title,omitempty"`
}

type Search_books_count struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Search_id struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Search_image_url struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_key struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_method struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_original_publication_day struct {
	Attr_nil  string `xml:" nil,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Search_original_publication_month struct {
	Attr_nil  string `xml:" nil,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Search_original_publication_year struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Search_query struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_query_time_seconds struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_ratings_count struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Search_results struct {
	Search_work []Search_work `xml:" work,omitempty" json:"work,omitempty"`
}

type Search_results_end struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_results_start struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_root struct {
	Search_GoodreadsResponse Search_GoodreadsResponse `xml:" GoodreadsResponse,omitempty" json:"GoodreadsResponse,omitempty"`
}

type Search_search struct {
	Search_query              Search_query              `xml:" query,omitempty" json:"query,omitempty"`
	Search_query_time_seconds Search_query_time_seconds `xml:" query-time-seconds,omitempty" json:"query-time-seconds,omitempty"`
	Search_results            Search_results            `xml:" results,omitempty" json:"results,omitempty"`
	Search_results_end        Search_results_end        `xml:" results-end,omitempty" json:"results-end,omitempty"`
	Search_results_start      Search_results_start      `xml:" results-start,omitempty" json:"results-start,omitempty"`
	Search_source             Search_source             `xml:" source,omitempty" json:"source,omitempty"`
	Search_total_results      Search_total_results      `xml:" total-results,omitempty" json:"total-results,omitempty"`
}

type Search_small_image_url struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_source struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_text_reviews_count struct {
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Search_title struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_total_results struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Search_work struct {
	Search_average_rating             Search_average_rating             `xml:" average_rating,omitempty" json:"average_rating,omitempty"`
	Search_best_book                  Search_best_book                  `xml:" best_book,omitempty" json:"best_book,omitempty"`
	Search_books_count                Search_books_count                `xml:" books_count,omitempty" json:"books_count,omitempty"`
	Search_id                         Search_id                         `xml:" id,omitempty" json:"id,omitempty"`
	Search_original_publication_day   Search_original_publication_day   `xml:" original_publication_day,omitempty" json:"original_publication_day,omitempty"`
	Search_original_publication_month Search_original_publication_month `xml:" original_publication_month,omitempty" json:"original_publication_month,omitempty"`
	Search_original_publication_year  Search_original_publication_year  `xml:" original_publication_year,omitempty" json:"original_publication_year,omitempty"`
	Search_ratings_count              Search_ratings_count              `xml:" ratings_count,omitempty" json:"ratings_count,omitempty"`
	Search_text_reviews_count         Search_text_reviews_count         `xml:" text_reviews_count,omitempty" json:"text_reviews_count,omitempty"`
}
