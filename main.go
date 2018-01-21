package main

import (
  "fmt"
  "github.com/yanatan16/golang-instagram/instagram"
  "net/url"
)
// app is in sandbox mode and Client access bellow is for testing only.

// Client ID 881def4c577c4224b776034a7f4053e6

// Client Secret e58d56c2b8d34ecabb6b9b446320ad8c  RESET SECRET

// https://api.instagram.com/oauth/authorize/?client_id=881def4c577c4224b776034a7f4053e6&redirect_uri=e58d56c2b8d34ecabb6b9b446320ad8c&response_type=code
func main() {

  authenticatedApi := &instagram.Api{
    AccessToken: "c6a415df72224c3793af8b205894089d",
  }
  anotherAuthenticatedApi := instagram.New("881def4c577c4224b776034a7f4053e6", "e58d56c2b8d34ecabb6b9b446320ad8c", "c6a415df72224c3793af8b205894089d", false)


// - if enforceSigned true

func DoSomeInstagramApiStuff(accessToken string) {
  api := New("", accessToken)

  if ok, err := api.VerifyCredentials(); !ok {
    panic(err)
  }

  var myId string

  // Get yourself!
  if resp, err := api.GetSelf(); err != nil {
    panic(err)
  } else {
    // A response has two fields: Meta which you shouldn't really care about
    // And whatever your getting, in this case, a User
    me := resp.User
    fmt.Printf("My userid is %s and I have %d followers\n", me.Id, me.Counts.FollowedBy)
  }

  params := url.Values{}
  params.Set("count", "1")
  if resp, err := api.GetUserRecentMedia("self" /* this works :) */, params); err != nil {
    panic(err)
  } else {
    if len(resp.Medias) == 0 { // [sic]
      panic("I should have some sort of media posted on instagram!")
    }
    media := resp.Medias[0]
    fmt.Println("My last media was a %s with %d comments and %d likes. (url: %s)\n", media.Type, media.Comments.Count, media.Like.Count, media.Link)
  }
}

}