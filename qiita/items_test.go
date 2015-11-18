package qiita

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestItemsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"query":    "query",
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `[
      {
        "rendered_body": "<h1>Example</h1>",
        "body": "# Example",
        "coediting": false,
        "id": "4bd431809afb1bb99e4f",
        "private": false,
        "tags": [],
        "title": "Example title",
        "url": "https://qiita.com/yaotti/items/4bd431809afb1bb99e4f",
        "user": {}
      }
    ]`)
	})
	items, err := client.Items.List("query", &ListOptions{Page: 1, PerPage: 2})
	if err != nil {
		t.Errorf("Issues.List returned error: %v", err)
	}
	want := []Item{
		{
			Body:         "# Example",
			RenderedBody: "<h1>Example</h1>",
			CoEditing:    false,
			Id:           "4bd431809afb1bb99e4f",
			Private:      false,
			Tags:         []AttachedTag{},
			Title:        "Example title",
			URL:          "https://qiita.com/yaotti/items/4bd431809afb1bb99e4f",
			User:         User{},
		},
	}
	if !reflect.DeepEqual(items, want) {
		t.Errorf("Issues.List returned %+v, want %+v", items, want)
	}
}
