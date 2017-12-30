package kkbox

import (
// "fmt"
)

// FetchReleaseCategory List of new release categories.
func (b *Box) FetchReleaseCategory(params ...Param) (*ReleaseCategoryList, error) {
	resp := new(ReleaseCategoryList)
	url := ReleaseCategoryURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// // FetchFeaturedPlayList retrieve information of the featured playlist with {playlist_id}.
// func (b *Box) FetchFeaturedPlayList(playList string, params ...Param) (*PlayListData, error) {
// 	resp := new(PlayListData)
// 	url := fmt.Sprintf(FeaturedSinglePlayListURL, playList)
// 	err := b.fetchData(url, resp, params...)

// 	return resp, err
// }

// // FetchFeaturedPlayListTrack list the songs of a featured playlist.
// func (b *Box) FetchFeaturedPlayListTrack(playList string, params ...Param) (*TrackData, error) {
// 	resp := new(TrackData)
// 	url := fmt.Sprintf(FeaturedPlayListTrackURL, playList)
// 	err := b.fetchData(url, resp, params...)

// 	return resp, err
// }
