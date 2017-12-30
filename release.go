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

// FetchReleaseCategoryAlbum list of albums of a new release category.
func (b *Box) FetchReleaseCategoryAlbum(id string, params ...Param) (*ArtistAlbumData, error) {
	resp := new(ArtistAlbumData)
	url := fmt.Sprintf(ReleaseCategoryAlbumURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}
