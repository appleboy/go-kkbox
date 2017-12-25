package kkbox

import (
	"time"
)

// HitDatas List of new hits playlists.
type HitDatas struct {
	Data []struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		Images      []Image   `json:"images"`
		UpdatedAt   time.Time `json:"updated_at"`
		Owner       Owner     `json:"owner"`
	} `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// FetchHits List of new hits playlists.
func (b *Box) FetchHits(params ...Param) (*HitDatas, error) {
	resp := new(HitDatas)
	url := ChartURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}
