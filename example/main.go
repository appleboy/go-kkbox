package main

import (
	"fmt"
	"log"

	"github.com/appleboy/go-kkbox"
)

var clientID string
var clientSecret string

func main() {
	k, err := kkbox.New(clientID, clientSecret)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", k)

	// fetch charts
	resp, err := k.FetchCharts()
	fmt.Printf("%#v\n", resp)
	fmt.Printf("%#v\n", err)
}
