package main

import (
  "fmt"
  "github.com/yanatan16/golang-instagram/instagram"
  "net/url"

//  "net/url"
)
// app is in sandbox mode and Client access bellow is for testing only.

// Client ID 881def4c577c4224b776034a7f4053e6

// Client Secret e58d56c2b8d34ecabb6b9b446320ad8c  RESET SECRET

// https://api.instagram.com/oauth/authorize/?client_id=881def4c577c4224b776034a7f4053e6&redirect_uri=e58d56c2b8d34ecabb6b9b446320ad8c&response_type=code
var AccessToken string = "c6a415df72224c3793af8b205894089d"
var ClientID string = "881def4c577c4224b776034a7f4053e6"
var ClientSecret string = "e58d56c2b8d34ecabb6b9b446320ad8c"

func main() {

fmt.Println("AccessToken: ",AccessToken)
fmt.Println("ClientID: ", ClientID)
fmt.Println("ClientSecret: ", ClientSecret)

ExampleApi_GetUserRecentMedia()
}

// ExampleApi_GetUserRecentMedia : Get the most recent media published by the owner
func ExampleApi_GetUserRecentMedia() {

  // *** or ***
  api := instagram.New(ClientID, ClientSecret, AccessToken, true)

  params := url.Values{}
  params.Set("count", "3") // 4 images in this set
  params.Set("max_timestamp", "1466809870")
  params.Set("min_timestamp", "1396751898")
  mediasResponse, err := api.GetUserRecentMedia(ccistulli_id, params)

  if err != nil {
    panic(err)
  }

  for _, media := range mediasResponse.Medias {
    processMedia(&media)
  }
}

