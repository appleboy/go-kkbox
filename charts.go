package kkbox

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/httplib"
)

// FetchCharts List of song rankings.
func (b *Box) FetchCharts(params ...Param) (*GroupListData, error) {
	resp := new(GroupListData)
	url := ChartURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchChartPlayList to retrieve information of the song ranking with {playlist_id}.
func (b *Box) FetchChartPlayList(platListID string, params ...Param) (*PlayListData, error) {
	resp := new(PlayListData)
	url := fmt.Sprintf(ChartPlayListURL, platListID)
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchChartPlayListTrack to retrieve information of the song ranking with {playlist_id}.
func (b *Box) FetchChartPlayListTrack(platListID string, params ...Param) (*TrackData, error) {
	resp := new(TrackData)
	url := fmt.Sprintf(ChartPlayListTrackURL, platListID)
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

	fmt.Println(string(req.DumpRequest()))

	return req.ToJSON(res)
}
