package kkbox

import (
	"fmt"
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
