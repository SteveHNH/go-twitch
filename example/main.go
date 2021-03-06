package main

import (
	"log"
	"os"

	"github.com/stevehnh/go-twitch"
)

var clientID string

func init() {
  clientID = os.Getenv("CLIENT_ID")
  bearerToken = os.Getenv("BEARER_TOKEN")
}

func main() {
  twitchSession, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID, BearerToken: bearerToken })
  if err != nil {
    log.Fatalln(err)
  }

  searchChannelsInput := twitch.SearchChannelsInputType{
    Query: "knspriggs",   // see https://github.com/justintv/Twitch-API/blob/master/v3_resources/search.md for query syntax
    Limit: 2,             //optional
    Offset: 0,            //optional
  }

  resp, err := twitchSession.SearchChannels(&searchChannelsInput)
  if err != nil {
    log.Fatalln(err)
  }
  log.Printf("Resp: \n%#v", resp)
}
