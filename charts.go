package kkbox

import (
	"crypto/tls"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

// Summary page data
type Summary struct {
	Total int `json:"total"`
}

type Owner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Image for music
type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

// Paging data
type Paging struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

// ChartDatas for charts list
type ChartDatas struct {
	Data []struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		Images      []Image   `json:"images"`
		UpdatedAt   time.Time `json:"updated_at"`
		Owner       Owner     `json:"owner"`
	} `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// Artist struct
type Artist struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	URL    string  `json:"url"`
	Images []Image `json:"images"`
}

// Album data
type Album struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	URL                  string   `json:"url"`
	Explicitness         bool     `json:"explicitness"`
	AvailableTerritories []string `json:"available_territories"`
	ReleaseDate          string   `json:"release_date"`
	Images               []Image  `json:"images"`
	Artist               Artist   `json:"artist"`
}

// Track data
type Track struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Duration             int      `json:"duration"`
	URL                  string   `json:"url"`
	TrackNumber          int      `json:"track_number"`
	Explicitness         bool     `json:"explicitness"`
	AvailableTerritories []string `json:"available_territories"`
	Album                Album    `json:"album"`
}

// TrackDatas to retrieve information of the song ranking
type TrackDatas struct {
	Tracks struct {
		Data    []Track `json:"data"`
		Paging  Paging  `json:"paging"`
		Summary Summary `json:"summary"`
	} `json:"tracks"`
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Images      []Image   `json:"images"`
	UpdatedAt   time.Time `json:"updated_at"`
	Owner       Owner     `json:"owner"`
}

// Param for http get parameter
type Param struct {
	Territory string
	Page      int
	PerPage   int
}

// FetchCharts List of song rankings.
func (b *Box) FetchCharts(params ...Param) (*ChartDatas, error) {
	resp := new(ChartDatas)
	url := ChartURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchChartPlayList to retrieve information of the song ranking with {playlist_id}.
func (b *Box) FetchChartPlayList(platListID string, params ...Param) (*TrackDatas, error) {
	resp := new(TrackDatas)
	url := fmt.Sprintf(ChartPlayListURL, platListID)
	err := b.fetchData(url, resp, params...)

	return resp, err
}

func (b *Box) fetchData(url string, res interface{}, params ...Param) error {
	req := httplib.Get(url).Debug(b.Debug)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	offset := 0
	limit := 10
	territory := "TW"
	if len(params) > 0 {
		param := params[0]
		if param.PerPage > 0 {
			limit = param.PerPage
		}
		if param.Page > 1 {
			offset = (param.Page - 1) * limit
		}
		if param.Territory != "" {
			territory = param.Territory
		}
	}

	// Add authorization header
	authorization := b.Auth.TokenType + " " + b.Auth.AccessToken
	req.Header("Authorization", authorization)

	req.Param("offset", strconv.Itoa(offset))
	req.Param("limit", strconv.Itoa(limit))
	req.Param("territory", territory)

	return req.ToJSON(res)
}
