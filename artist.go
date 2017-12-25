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

// FetchArtistAlbum list of albums of an artist.
func (b *Box) FetchArtistAlbum(id string, params ...Param) (*ArtistAlbumData, error) {
	resp := new(ArtistAlbumData)
	url := fmt.Sprintf(ArtistAlbumURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}
