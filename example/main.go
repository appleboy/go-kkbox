package main

import (
	"fmt"
	"log"
	"os"

	"github.com/appleboy/go-kkbox"
	"github.com/davecgh/go-spew/spew"
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

	// // fetch charts
	// charts, err := k.FetchCharts()
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }

	// fmt.Printf("%#v\n", charts)

	// ranks, err := k.FetchChartPlayList("4nUZM-TY2aVxZ2xaA-")
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }

	// spew.Dump(ranks)

	// tracks, err := k.FetchChartPlayListTrack("4nUZM-TY2aVxZ2xaA-", kkbox.Param{
	// 	PerPage: 1,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }

	// spew.Dump(tracks)

	// fetch hits
	hits, err := k.FetchHits(kkbox.Param{
		PerPage: 2,
	})
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	fmt.Println("hit length:", len(hits.Data))
	// spew.Dump(hits)

	list, err := k.FetchHitPlayList("DZrC8m29ciOFY2JAm3", kkbox.Param{
		Page:    2,
		PerPage: 2,
	})
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	fmt.Println("list length:", len(list.Tracks.Data))
	// spew.Dump(list)

	tracks, err := k.FetchHitPlayListTrack("DZrC8m29ciOFY2JAm3", kkbox.Param{
		PerPage: 3,
	})
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	fmt.Println("track length:", len(tracks.Data))
	spew.Dump(tracks)
}
