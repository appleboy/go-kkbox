package kkbox

import (
	"fmt"
)

// FetchMoodStationList List of mood stations.
func (b *Box) FetchMoodStationList(params ...Param) (*MoodListData, error) {
	resp := new(MoodListData)
	url := MoodStationURL
	err := b.fetchData(url, resp, params...)

	return resp, err
}

// FetchMoodStation retrieve information of the mood station with {station_id}.
func (b *Box) FetchMoodStation(id string, params ...Param) (*MoodData, error) {
	resp := new(MoodData)
	url := fmt.Sprintf(MoodSingleStationURL, id)
	err := b.fetchData(url, resp, params...)

	return resp, err
}
