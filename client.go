package goodreads

import (
	"errors"
	"net/http"

	"github.com/mrjones/oauth"
)

const (
	REQUEST_TOKEN_URL   = "http://www.goodreads.com/oauth/request_token"
	AUTHORIZE_TOKEN_URL = "http://www.goodreads.com/oauth/authorize"
	ACCESS_TOKEN_URL    = "http://www.goodreads.com/oauth/access_token"
)

//Client is the Goodreads API Client object
type Client struct {
	Consumer    *oauth.Consumer
	user        *oauth.AccessToken
	client      *http.Client
	consumerKey string
}

//NewClient is the constructor with only the Consumer key and secret
func NewClient(key string, secret string) *Client {
	c := Client{
		consumerKey: key,
	}
	c.SetConsumer(key, secret)
	return &c
}

// Constructor with Consumer key/secret and user token/secret
func NewClientWithToken(consumerKey, consumerSecret, token, tokenSecret string) *Client {
	c := NewClient(consumerKey, consumerSecret)
	c.SetToken(token, tokenSecret)
	return c
}

// Set Consumer credentials, invalidates any previously cached client
func (c *Client) SetConsumer(consumerKey string, consumerSecret string) {
	c.Consumer = oauth.NewConsumer(consumerKey, consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   REQUEST_TOKEN_URL,
			AuthorizeTokenUrl: AUTHORIZE_TOKEN_URL,
			AccessTokenUrl:    ACCESS_TOKEN_URL,
		})
	c.client = nil
}

// Set user credentials, invalidates any previously cached client
func (c *Client) SetToken(token string, secret string) {
	c.user = &oauth.AccessToken{
		AdditionalData: nil,
		Secret:         secret,
		Token:          token,
	}
	c.client = nil
}

// Retrieve the underlying HTTP client
func (c *Client) GetHttpClient() (*http.Client, error) {
	if c.Consumer == nil {
		return nil, errors.New("Consumer credentials are not set")

	}
	if c.user == nil {
		c.SetToken("", "")
	}
	if c.client == nil {
		c.client, _ = c.Consumer.MakeHttpClient(c.user)
		c.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	return c.client, nil
}
