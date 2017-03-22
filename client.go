package goodreads

import(
	"github.com/dghubble/oauth1"
	"golang.org/x/net/context"
	"net/http"
)

// The Goodreads API Client object
type Client struct {
	consumer *oauth1.Config
	user *oauth1.Token
	client *http.Client
}

// Constructor with only the consumer key and secret
func NewClient(consumerKey string, consumerSecret string) *Client {
	c := Client{}
	c.SetConsumer(consumerKey, consumerSecret)
	return &c
}

// Constructor with consumer key/secret and user token/secret
func NewClientWithToken(consumerKey string, consumerSecret string, token string, tokenSecret string) *Client {
	c := NewClient(consumerKey, consumerSecret)
	c.SetToken(token, tokenSecret)
	return c
}


// Set consumer credentials, invalidates any previously cached client
func (c *Client) SetConsumer(consumerKey string, consumerSecret string) {
	c.consumer = oauth1.NewConfig(consumerKey, consumerSecret)
	c.client = nil
}

// Set user credentials, invalidates any previously cached client
func (c *Client) SetToken(token string, tokenSecret string) {
	c.user = oauth1.NewToken(token, tokenSecret)
	c.client = nil
}


// Retrieve the underlying HTTP client
func (c *Client) GetHttpClient() *http.Client {
	if c.consumer == nil {
		panic("Consumer credentials are not set")
	}
	if c.user == nil {
		c.SetToken("", "")
	}
	if c.client == nil {
		c.client = c.consumer.Client(context.TODO(), c.user)
		c.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	return c.client
}