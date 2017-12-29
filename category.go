package kkbox

import (
	"fmt"
)

// FetchFeaturedCategory List of featured playlist categories.
func (b *Box) FetchFeaturedCategory(params ...Param) (*GroupListData, error) {
	resp := new(GroupListData)
	url := FeaturedCategoryURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchSingleFeaturedCategory retrieve information of the featured playlist with {playlist_id}.
func (b *Box) FetchSingleFeaturedCategory(id string, params ...Param) (*CategoryListData, error) {
	resp := new(CategoryListData)
	url := fmt.Sprintf(FeaturedSingleCategoryURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}
