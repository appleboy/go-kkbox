package kkbox

import (
	"time"
)

// Summary page data
type Summary struct {
	Total int `json:"total"`
}

// Owner data
type Owner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Image for music
type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

// Paging data
type Paging struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

// Artist struct
type Artist struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	URL    string  `json:"url"`
	Images []Image `json:"images"`
}

// Album data
type Album struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	URL                  string   `json:"url"`
	Explicitness         bool     `json:"explicitness"`
	AvailableTerritories []string `json:"available_territories"`
	ReleaseDate          string   `json:"release_date"`
	Images               []Image  `json:"images"`
	Artist               Artist   `json:"artist"`
}

// Track data
type Track struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Duration             int      `json:"duration"`
	URL                  string   `json:"url"`
	TrackNumber          int      `json:"track_number"`
	Explicitness         bool     `json:"explicitness"`
	AvailableTerritories []string `json:"available_territories"`
	Album                Album    `json:"album"`
}

// Song for play list
type Song struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Images      []Image   `json:"images"`
	UpdatedAt   time.Time `json:"updated_at"`
	Owner       Owner     `json:"owner"`
}

// GroupListData for song list
type GroupListData struct {
	Data    []Song  `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// PlayListData to retrieve information of playlist
type PlayListData struct {
	Tracks struct {
		Data    []Track `json:"data"`
		Paging  Paging  `json:"paging"`
		Summary Summary `json:"summary"`
	} `json:"tracks"`
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Images      []Image   `json:"images"`
	UpdatedAt   time.Time `json:"updated_at"`
	Owner       Owner     `json:"owner"`
}

// TrackData List tracks of a chart playlist.
type TrackData struct {
	Data    []Track `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// Param for http get parameter
type Param struct {
	// The search keywords, URL encoded.
	Q string
	// The content type. Content type could be track, album, artist, or playlist.
	Type string
	// Territory code, i.e. TW, HK, SG, MY, JP, of search content.
	Territory string
	Page      int
	// The number of items to return per page, not to exceed 50.
	PerPage int
}

// SearchData for search results
type SearchData struct {
	Artists struct {
		Data    []Artist `json:"data"`
		Paging  Paging   `json:"paging"`
		Summary Summary  `json:"summary"`
	} `json:"artists"`
	Tracks struct {
		Data    []Track `json:"data"`
		Paging  Paging  `json:"paging"`
		Summary Summary `json:"summary"`
	} `json:"tracks"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// AlbumTrackData list of tracks of an album.
type AlbumTrackData struct {
	Data    []Track `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// ArtistAlbumData list of albums of an artist.
type ArtistAlbumData struct {
	Data    []Album `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// CategoryListData for category
type CategoryListData struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	Images    []Image       `json:"images"`
	Playlists GroupListData `json:"playlists"`
}

// MoodListData mood stations.
type MoodListData struct {
	Data []struct {
		ID     string  `json:"id"`
		Name   string  `json:"name"`
		Images []Image `json:"images"`
	} `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// MoodData retrieve information of the mood station with {station_id}.
type MoodData struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Tracks struct {
		Data    []Track `json:"data"`
		Paging  Paging  `json:"paging"`
		Summary Summary `json:"summary"`
	} `json:"tracks"`
}

// GenreList List of genre stations.
type GenreList struct {
	Data []struct {
		ID       string `json:"id"`
		Category string `json:"category"`
		Name     string `json:"name"`
	} `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// GenreData retrieve information of the genre station with {station_id}.
type GenreData struct {
	Category string `json:"category"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Tracks   struct {
		Data    []Track `json:"data"`
		Paging  Paging  `json:"paging"`
		Summary Summary `json:"summary"`
	} `json:"tracks"`
}

// ReleaseCategoryList List of new release categories.
type ReleaseCategoryList struct {
	Data []struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	} `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// AlbumList retrieve information of the new release category with {category_id}.
type AlbumList struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Albums struct {
		Data    []Album `json:"data"`
		Paging  Paging  `json:"paging"`
		Summary Summary `json:"summary"`
	} `json:"albums"`
}
