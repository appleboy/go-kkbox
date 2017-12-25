package kkbox

const (
	// OAuthTokenURL for auth token url
	OAuthTokenURL = "https://account.kkbox.com/oauth2/token"

	// ChartURL List of song rankings.
	ChartURL = "https://api.kkbox.com/v1.1/charts"
	// ChartPlayListURL retrieve information of the song ranking
	// with {playlist_id}.
	ChartPlayListURL = ChartURL + "/%s"
	// ChartPlayListTrackURL list tracks of a chart playlist.
	ChartPlayListTrackURL = ChartURL + "/%s/tracks"

	// NewHitURL List of new hits playlists.
	NewHitURL = "https://api.kkbox.com/v1.1/new-hits-playlists"
	// NewHitPlayListURL retrieve information of the new hits playlist
	// with {playlist_id}.
	NewHitPlayListURL = NewHitURL + "/%s"
	// NewHitPlayListTrackURL List of tracks of a new hits playlist.
	NewHitPlayListTrackURL = NewHitURL + "/%s/tracks"

	// SearchURL search for various objects.
	SearchURL = "https://api.kkbox.com/v1.1/search"

	// TrackURL retrieve information of the song with {track_id}.
	TrackURL = "https://api.kkbox.com/v1.1/tracks/%s"
)
