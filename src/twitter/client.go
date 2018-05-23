package twitter

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

type Client struct {
	URL                            *url.URL
	HTTPClient                     *http.Client
	APIKey, APISecret, BearerToken string
	Logger                         *log.Logger
}

func NewClient(key, secret string, logger *log.Logger) (*Client, error) {
	if len(key) == 0 {
		return nil, errors.New("missing key")
	}
	if len(secret) == 0 {
		return nil, errors.New("missing secret")
	}
	if logger == nil {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	u, _ := url.ParseRequestURI("https://api.twitter.com")
	c := &Client{URL: u, HTTPClient: &http.Client{}, APIKey: key, APISecret: secret, BearerToken: "", Logger: logger}

	resp, err := c.PostOauth2Token(&Oauth2TokenRequest{GrantType: "client_credentials"})
	if err != nil {
		return nil, err
	}
	c.BearerToken = resp.AccessToken

	return c, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (c *Client) newRequest(method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if c.BearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.BearerToken)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	return req, nil
}
