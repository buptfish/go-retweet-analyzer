package twitter

import (
	"errors"
	"net/http"
	"net/url"
)

type UsersShowRequest struct {
	UserID          int    `json:"user_id,omitempty"`
	ScreenName      string `json:"screen_name,omitempty"`
	IncludeEntities bool   `json:"include_entities,omitempty"`
}

type UsersShowResponse struct {
	ID              int         `json:"id"`
	IDStr           string      `json:"id_str"`
	Name            string      `json:"name"`
	ScreenName      string      `json:"screen_name"`
	Location        string      `json:"location"`
	ProfileLocation interface{} `json:"profile_location"`
	Description     string      `json:"description"`
	URL             string      `json:"url"`
	Entities        struct {
		URL struct {
			Urls []struct {
				URL         string `json:"url"`
				ExpandedURL string `json:"expanded_url"`
				DisplayURL  string `json:"display_url"`
				Indices     []int  `json:"indices"`
			} `json:"urls"`
		} `json:"url"`
		Description struct {
			Urls []struct {
				URL         string `json:"url"`
				ExpandedURL string `json:"expanded_url"`
				DisplayURL  string `json:"display_url"`
				Indices     []int  `json:"indices"`
			} `json:"urls"`
		} `json:"description"`
	} `json:"entities"`
	Protected       bool        `json:"protected"`
	FollowersCount  int         `json:"followers_count"`
	FriendsCount    int         `json:"friends_count"`
	ListedCount     int         `json:"listed_count"`
	CreatedAt       string      `json:"created_at"`
	FavouritesCount int         `json:"favourites_count"`
	UtcOffset       interface{} `json:"utc_offset"`
	TimeZone        interface{} `json:"time_zone"`
	GeoEnabled      bool        `json:"geo_enabled"`
	Verified        bool        `json:"verified"`
	StatusesCount   int         `json:"statuses_count"`
	Lang            string      `json:"lang"`
	Status          struct {
		CreatedAt string `json:"created_at"`
		ID        int64  `json:"id"`
		IDStr     string `json:"id_str"`
		Text      string `json:"text"`
		Truncated bool   `json:"truncated"`
		Entities  struct {
			Hashtags     []interface{} `json:"hashtags"`
			Symbols      []interface{} `json:"symbols"`
			UserMentions []struct {
				ScreenName string `json:"screen_name"`
				Name       string `json:"name"`
				ID         int    `json:"id"`
				IDStr      string `json:"id_str"`
				Indices    []int  `json:"indices"`
			} `json:"user_mentions"`
			Urls []interface{} `json:"urls"`
		} `json:"entities"`
		Source               string      `json:"source"`
		InReplyToStatusID    int64       `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr string      `json:"in_reply_to_status_id_str"`
		InReplyToUserID      int         `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   string      `json:"in_reply_to_user_id_str"`
		InReplyToScreenName  string      `json:"in_reply_to_screen_name"`
		Geo                  interface{} `json:"geo"`
		Coordinates          interface{} `json:"coordinates"`
		Place                interface{} `json:"place"`
		Contributors         interface{} `json:"contributors"`
		IsQuoteStatus        bool        `json:"is_quote_status"`
		RetweetCount         int         `json:"retweet_count"`
		FavoriteCount        int         `json:"favorite_count"`
		Favorited            bool        `json:"favorited"`
		Retweeted            bool        `json:"retweeted"`
		Lang                 string      `json:"lang"`
	} `json:"status"`
	ContributorsEnabled            bool        `json:"contributors_enabled"`
	IsTranslator                   bool        `json:"is_translator"`
	IsTranslationEnabled           bool        `json:"is_translation_enabled"`
	ProfileBackgroundColor         string      `json:"profile_background_color"`
	ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool        `json:"profile_background_tile"`
	ProfileImageURL                string      `json:"profile_image_url"`
	ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
	ProfileBannerURL               string      `json:"profile_banner_url"`
	ProfileLinkColor               string      `json:"profile_link_color"`
	ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string      `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
	HasExtendedProfile             bool        `json:"has_extended_profile"`
	DefaultProfile                 bool        `json:"default_profile"`
	DefaultProfileImage            bool        `json:"default_profile_image"`
	Following                      interface{} `json:"following"`
	FollowRequestSent              interface{} `json:"follow_request_sent"`
	Notifications                  interface{} `json:"notifications"`
	TranslatorType                 string      `json:"translator_type"`
}

func (c *Client) GetUsersShow(p *UsersShowRequest) (*UsersShowResponse, error) {
	const spath = "1.1/users/show.json"

	data := url.Values{}
	switch {
	case p.UserID != 0:
		data.Set("user_id", string(p.UserID))
	case p.ScreenName != "":
		data.Set("screen_name", p.ScreenName)
	default:
		return nil, errors.New("missing parameter")
	}

	req, err := c.newRequest(http.MethodGet, spath, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = data.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var usersShow UsersShowResponse
	if err := decodeBody(resp, &usersShow); err != nil {
		return nil, err
	}

	return &usersShow, nil
}
