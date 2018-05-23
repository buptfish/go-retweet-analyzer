package twitter

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Oauth2TokenRequest struct {
	GrantType string `json:"grant_type"`
}

type Oauth2TokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func (c *Client) PostOauth2Token(p *Oauth2TokenRequest) (*Oauth2TokenResponse, error) {
	const spath = "oauth2/token"

	data := url.Values{}
	data.Set("grant_type", p.GrantType)

	req, err := c.newRequest(http.MethodPost, spath, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIKey, c.APISecret)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var oauth2Token Oauth2TokenResponse
	if err := decodeBody(resp, &oauth2Token); err != nil {
		return nil, err
	}

	return &oauth2Token, nil
}
