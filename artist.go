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

// FetchArtistTopTrack List of top tracks of an artist.
func (b *Box) FetchArtistTopTrack(id string, params ...Param) (*AlbumTrackData, error) {
	resp := new(AlbumTrackData)
	url := fmt.Sprintf(ArtistTopTrackURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchArtistRelated List of top tracks of an artist.
func (b *Box) FetchArtistRelated(id string, params ...Param) (*ArtistAlbumData, error) {
	resp := new(ArtistAlbumData)
	url := fmt.Sprintf(ArtistRelatedArtistURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}
