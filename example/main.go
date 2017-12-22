package main

import (
	"fmt"

	"github.com/appleboy/go-kkbox"
)

var clientID = "1234"
var clientSecret = "1234"

func main() {
	k, err := kkbox.New(clientID, clientSecret)

	fmt.Printf("%#v\n", k)
	fmt.Printf("%#v\n", err)
}
