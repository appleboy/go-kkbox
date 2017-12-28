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

	// AlbumURL retrieve information of the album with {album_id}.
	AlbumURL = "https://api.kkbox.com/v1.1/albums/%s"
	// AlbumTrackURL list of tracks of an album.
	AlbumTrackURL = AlbumURL + "/tracks"

	// ArtistURL retrieve information of the artist with {artist_id}.
	ArtistURL = "https://api.kkbox.com/v1.1/artists/%s"
	// ArtistAlbumURL List of albums of an artist.
	ArtistAlbumURL = ArtistURL + "/albums"
	// ArtistTopTrackURL List of top tracks of an artist.
	ArtistTopTrackURL = ArtistURL + "/top-tracks"
	// ArtistRelatedArtistURL Get related artists of an artist.
	ArtistRelatedArtistURL = ArtistURL + "/related-artists"

	// SharedPlaylistURL retrieve information of the shared playlist with {playlist_id}.
	SharedPlaylistURL = "https://api.kkbox.com/v1.1/shared-playlists/%s"
	// SharedPlaylistTrackURL list of songs of a shared playlist.
	SharedPlaylistTrackURL = "https://api.kkbox.com/v1.1/shared-playlists/%s/tracks"

	// FeaturedPlayListURL List of songs hand-picked and arranged by KKBOX editors.
	FeaturedPlayListURL = "https://api.kkbox.com/v1.1/featured-playlists"
	// FeaturedSinglePlayListURL retrieve information of the featured playlist with {playlist_id}.
	FeaturedSinglePlayListURL = FeaturedPlayListURL + "/%s"
)
