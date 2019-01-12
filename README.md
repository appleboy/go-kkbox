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
* [x] Shared Playlists
  - [x] /shared-playlists/{playlist_id}
  - [x] /shared-playlists/{playlist_id}/tracks
* [x] Featured Playlists
  - [x] /featured-playlists
  - [x] /featured-playlists/{playlist_id}
  - [x] /featured-playlists/{playlist_id}/tracks
* [x] Featured Playlist Categories
  - [x] /featured-playlist-categories
  - [x] /featured-playlist-categories/{category_id}
  - [x] /featured-playlist-categories/{category_id}/playlists
* [x] Mood Stations
  - [x] /mood-stations
  - [x] /mood-stations/{station_id}
* [x] Genre Stations
  - [x] /genre-stations
  - [x] /genre-stations/{station_id}
* [x] New Release Categories
  - [x] /new-release-categories
  - [x] /new-release-categories/{category_id}
  - [x] /new-release-categories/{category_id}/albums
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

## Install

Install SDK library

```
$ go get -u github.com/appleboy/go-kkbox
```

Inital the KKBox client:

```go
package main

import (
	"log"
	"os"

	"github.com/appleboy/go-kkbox"
)

var clientID = os.Getenv("CLIENT_ID")
var clientSecret = os.Getenv("CLIENT_SECRET")

func main() {
	if clientID == "" || clientSecret == "" {
		log.Fatal("missing client id or secret")
	}
	k, err := kkbox.New(clientID, clientSecret)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("====== kkbox client ======")
	spew.Dump(k)
	fmt.Println("====== kkbox end ======")

	// fetch charts
	charts, err := k.FetchCharts()
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	fmt.Printf("%#v\n", charts)

	ranks, err := k.FetchChartPlayList("4nUZM-TY2aVxZ2xaA-")
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	spew.Dump(ranks)

	tracks, err := k.FetchChartPlayListTrack("4nUZM-TY2aVxZ2xaA-", kkbox.Param{
		PerPage:   1,
		Page:      2,
		Territory: "HK",
	})
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	spew.Dump(tracks)
	log.Println("length: ", len(tracks.Data))
}
```

Run program:

```sh
$ CLIENT_ID=xxx CLIENT_SECRET=xxx go run main.go
```
