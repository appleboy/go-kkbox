# go-kkbox

[![GoDoc](https://godoc.org/github.com/appleboy/go-kkbox?status.svg)](https://godoc.org/github.com/appleboy/go-kkbox)
[![Build Status](http://drone.wu-boy.com/api/badges/appleboy/go-kkbox/status.svg)](http://drone.wu-boy.com/appleboy/go-kkbox)
[![codecov](https://codecov.io/gh/appleboy/go-kkbox/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/go-kkbox)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/go-kkbox)](https://goreportcard.com/report/github.com/appleboy/go-kkbox)

[KKBOX Open API](https://docs-en.kkbox.codes/) SDK for Golang.

## Features

* [x] Tracks
  - [x] /tracks/{track_id}
* [x] Albums
  - [x] /albums/{album_id}
  - [x] /albums/{album_id}/tracks
* [x] Artists
  - [x] /artists/{artist_id}
  - [x] /artists/{artist_id}/albums
  - [x] /artists/{artist_id}/top-tracks
  - [x] /artists/{artist_id}/related-artists
* [ ] Shared Playlists
  - [ ] /shared-playlists/{playlist_id}
  - [ ] /shared-playlists/{playlist_id}/tracks
* [ ] Featured Playlists
  - [ ] /featured-playlists
  - [ ] /featured-playlists/{playlist_id}
  - [ ] /featured-playlists/{playlist_id}/tracks
* [ ] Featured Playlist Categories
  - [ ] /featured-playlist-categories
  - [ ] /featured-playlist-categories/{category_id}
  - [ ] /featured-playlist-categories/{category_id}/playlists
* [ ] Mood Stations
  - [ ] /mood-stations
  - [ ] /mood-stations/{station_id}
* [ ] Genre Stations
  - [ ] /genre-stations
  - [ ] /genre-stations/{station_id}
* [ ] New Release Categories
  - [ ] /new-release-categories
  - [ ] /new-release-categories/{category_id}
  - [ ] /new-release-categories/{category_id}/albums
* [x] Search
  - [x] /search
* [x] New Hits Playlists
  - [x] /new-hits-playlists
  - [x] /new-hits-playlists/{playlist_id}
  - [x] /new-hits-playlists/{playlist_id}/tracks
* [x] Charts
  - [x] /charts
  - [x] /charts/{playlist_id}
  - [x] /charts/{playlist_id}/tracks
