package kkbox

import (
	"fmt"
)

// FetchReleaseCategory List of new release categories.
func (b *Box) FetchReleaseCategory(params ...Param) (*ReleaseCategoryList, error) {
	resp := new(ReleaseCategoryList)
	url := ReleaseCategoryURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchSingleReleaseCategory retrieve information of the new release category with {category_id}.
func (b *Box) FetchSingleReleaseCategory(id string, params ...Param) (*AlbumList, error) {
	resp := new(AlbumList)
	url := fmt.Sprintf(ReleaseSingleCategoryURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// // FetchFeaturedPlayListTrack list the songs of a featured playlist.
// func (b *Box) FetchFeaturedPlayListTrack(playList string, params ...Param) (*TrackData, error) {
// 	resp := new(TrackData)
// 	url := fmt.Sprintf(FeaturedPlayListTrackURL, playList)
// 	err := b.fetchData(url, resp, params...)

// 	return resp, err
// }
