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
	"strings"
)

type Oauth2TokenRequest struct {
	GrantType string `json:"grant_type"`
}

type Oauth2TokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

type UsersShowRequest struct {
	UserID     int    `json:"user_id,omitempty"`
	ScreenName string `json:"screen_name,omitempty"`
}

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

func (c *Client) newRequest(method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIKey, c.APISecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (c *Client) PostOauth2Token(p *Oauth2TokenRequest) (*Oauth2TokenResponse, error) {
	const spath = "oauth2/token"

	var data url.Values
	data.Set("grant_type", p.GrantType)

	req, err := c.newRequest(http.MethodPost, spath, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var oauth2Token Oauth2TokenResponse
	if err := decodeBody(resp, &oauth2Token); err != nil {
		return nil, err
	}

	return &oauth2Token, nil
}
