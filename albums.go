package kkbox

import (
	"fmt"
)

// FetchAlbum retrieve information of the album with {album_id}.
func (b *Box) FetchAlbum(id string) (*Album, error) {
	resp := new(Album)
	url := fmt.Sprintf(AlbumURL, id)
	err := b.fetchData(url, resp)

	return resp, err
}
