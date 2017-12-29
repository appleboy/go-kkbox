package kkbox

import (
	"fmt"
)

// FetchGenreStationList List of genre stations.
func (b *Box) FetchGenreStationList(params ...Param) (*GenreList, error) {
	resp := new(GenreList)
	url := GenreStationURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchGenreStation retrieve information of the genre station with {station_id}.
func (b *Box) FetchGenreStation(id string, params ...Param) (*GenreData, error) {
	resp := new(GenreData)
	url := fmt.Sprintf(GenreSingleStationURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}
