package kkbox

import (
// "fmt"
)

// FetchFeaturedCategory List of featured playlist categories.
func (b *Box) FetchFeaturedCategory(params ...Param) (*GroupListData, error) {
	resp := new(GroupListData)
	url := FeaturedCategoryURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}
