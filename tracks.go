package kkbox

import (
	"fmt"
)

// FetchTrack retrieve information of the song with {track_id}.
func (b *Box) FetchTrack(id string) (*Track, error) {
	resp := new(Track)
	url := fmt.Sprintf(TrackURL, id)
	err := b.fetchData(url, resp)

	return resp, err
}
