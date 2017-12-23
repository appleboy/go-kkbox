package kkbox

const (
	// OAuthTokenURL for auth token url
	OAuthTokenURL = "https://account.kkbox.com/oauth2/token"

	// ChartURL List of song rankings.
	ChartURL = "https://api.kkbox.com/v1.1/charts"
	// ChartPlayList retrieve information of the song ranking
	// with {playlist_id}.
	ChartPlayList = ChartURL + "/%s"
	// ChartPlayListTrack list tracks of a chart playlist.
	ChartPlayListTrack = ChartURL + "/%s/track"
)
