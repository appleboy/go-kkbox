package kkbox

import (
	"fmt"
)

// FetchSharedPlayList retrieve information of the shared playlist with {playlist_id}.
func (b *Box) FetchSharedPlayList(id string, params ...Param) (*PlayListData, error) {
	resp := new(PlayListData)
	url := fmt.Sprintf(SharedPlaylistURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchSharedPlayListTrack list of songs of a shared playlist..
func (b *Box) FetchSharedPlayListTrack(id string, params ...Param) (*TrackData, error) {
	resp := new(TrackData)
	url := fmt.Sprintf(SharedPlaylistTrackURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}
