package kkbox

import (
	"fmt"
)

// FetchArtist retrieve information of the artist with {artist_id}..
func (b *Box) FetchArtist(id string) (*Artist, error) {
	resp := new(Artist)
	url := fmt.Sprintf(ArtistURL, id)
	err := b.fetchData(url, resp)

	return resp, err
}
