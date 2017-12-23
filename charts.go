package kkbox

import (
	"crypto/tls"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

// ChartsResp for charts list
type ChartsResp struct {
	Data []struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Images      []struct {
			Height int    `json:"height"`
			Width  int    `json:"width"`
			URL    string `json:"url"`
		} `json:"images"`
		UpdatedAt time.Time `json:"updated_at"`
		Owner     struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"owner"`
	} `json:"data"`
	Paging struct {
		Offset   int    `json:"offset"`
		Limit    int    `json:"limit"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
	} `json:"paging"`
	Summary struct {
		Total int `json:"total"`
	} `json:"summary"`
}

// Param for http get parameter
type Param struct {
	Territory string
	Page      int
	PerPage   int
}

// FetchCharts List of song rankings.
func (b *Box) FetchCharts(params ...Param) (*ChartsResp, error) {
	resp := new(ChartsResp)
	url := ChartURL
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

	req.Header("Authorization", b.Auth.TokenType+" "+b.Auth.AccessToken)

	req.Param("offset", strconv.Itoa(offset))
	req.Param("limit", strconv.Itoa(limit))
	req.Param("territory", territory)

	return req.ToJSON(res)
}
