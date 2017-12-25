package kkbox

import (
	"fmt"
)

// FetchHits List of new hits playlists.
func (b *Box) FetchHits(params ...Param) (*GroupListData, error) {
	resp := new(GroupListData)
	url := ChartURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchHitPlayList retrieve information of the new hits playlist with {playlist_id}.
func (b *Box) FetchHitPlayList(playList string, params ...Param) (*PlayListData, error) {
	resp := new(PlayListData)
	url := fmt.Sprintf(NewHitPlayListURL, playList)
	err := b.fetchData(url, resp, params...)

	return resp, err
}
