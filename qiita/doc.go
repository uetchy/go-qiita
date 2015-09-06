/*
The package 'qiita' provides a client using Qiita API v2.

You can also use http.Client compatible client object like oauth2. Here are simple examples.

  import (
    "github.com/uetchy/go-qiita/qiita"
    "golang.org/x/oauth2"
    "fmt"
  )

  func main() {
    // Create OAuth2 client
    ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "personal access token"})
    tc := oauth2.NewClient(oauth2.NoContext, ts)

    // Create Qiita client using OAuth2 adapter
    client := qiita.NewClient(tc)

    // Fetch articles and print them
    items, _, _ := client.Items.List(&qiita.ItemsListOptions{Query: "Alfred"})
    fmt.Println(items)
  }

*/
package qiita
