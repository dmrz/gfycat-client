package main

import (
	"flag"
	"fmt"
	"net/http"
)

var clientId, clientSecret, reactionsTagName string

func init() {
	flag.StringVar(&clientId, "clientId", "", "GFYCAT client id")
	flag.StringVar(&clientSecret, "clientSecret", "", "GFYCAT client secret")
	flag.StringVar(&reactionsTagName, "tagName", "funny", "GFYCAT reactions tag name")
	flag.Parse()
}

func main() {
	api := ApiClient{
		clientId,
		clientSecret,
		"https://api.gfycat.com/v1",
		&http.Client{},
		OauthToken{},
	}
	// Print first reaction gif url
	reactions := api.getReactions(reactionsTagName, 1)
	fmt.Println(reactions.Gfycats[0].GifUrl)
}
