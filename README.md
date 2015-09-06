# go-qiita
Go binding for Qiita API.

## Usage
This is simple example of go-qiita using OAuth2 Personal Access Token.

```go
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
```

## Document

See [godoc](http://godoc.org/github.com/uetchy/go-qiita) for further information and instructions.

## Build

```
$ go get github.com/uetchy/go-qiita
$ cd $GOPATH/github.com/uetchy/go-qiita
$ go install
```
