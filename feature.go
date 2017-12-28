package kkbox

import (
	"fmt"
)

// FetchFeatured List of songs hand-picked and arranged by KKBOX editors.
func (b *Box) FetchFeatured(params ...Param) (*GroupListData, error) {
	resp := new(GroupListData)
	url := FeaturedPlayListURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchFeaturedPlayList retrieve information of the featured playlist with {playlist_id}.
func (b *Box) FetchFeaturedPlayList(playList string, params ...Param) (*PlayListData, error) {
	resp := new(PlayListData)
	url := fmt.Sprintf(FeaturedSinglePlayListURL, playList)
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// // FetchHitPlayListTrack list of tracks of a new hits playlist.
// func (b *Box) FetchHitPlayListTrack(playList string, params ...Param) (*TrackData, error) {
// 	resp := new(TrackData)
// 	url := fmt.Sprintf(NewHitPlayListTrackURL, playList)
// 	err := b.fetchData(url, resp, params...)

// 	return resp, err
// }
