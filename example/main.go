package main

import (
	"fmt"
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

	fmt.Printf("%#v\n", k)

	// fetch charts
	charts, err := k.FetchCharts()
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	fmt.Printf("%#v\n", charts)

	tracks, err := k.FetchChartPlayList("4nUZM-TY2aVxZ2xaA-")
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	fmt.Printf("%#v\n", tracks)
}
