package kkbox

import (
// "fmt"
)

// FetchMoodStationList List of mood stations.
func (b *Box) FetchMoodStationList(params ...Param) (*MoodListData, error) {
	resp := new(MoodListData)
	url := MoodStationURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// // FetchHitPlayList retrieve information of the new hits playlist with {playlist_id}.
// func (b *Box) FetchHitPlayList(playList string, params ...Param) (*PlayListData, error) {
// 	resp := new(PlayListData)
// 	url := fmt.Sprintf(NewHitPlayListURL, playList)
// 	err := b.fetchData(url, resp, params...)

// 	return resp, err
// }
