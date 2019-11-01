package qiita

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTagsService_Stocks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `[
      {
        "followers_count": 100,
        "icon_url": "https://s3-ap-northeast-1.amazonaws.com/qiita-tag-image/9de6a11d330f5694820082438f88ccf4a1b289b2/medium.jpg",
        "id": "qiita",
        "items_count": 200
      }
    ]`)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	items, err := client.Tags.List(opt)

	if err != nil {
		t.Errorf("Issues.List returned error: %v", err)
	}

	want := []Tag{
		{
			FollowersCount: 100,
			IconURL:        "https://s3-ap-northeast-1.amazonaws.com/qiita-tag-image/9de6a11d330f5694820082438f88ccf4a1b289b2/medium.jpg",
			Id:             "qiita",
			ItemsCount:     200,
		},
	}
	if !reflect.DeepEqual(items, want) {
		t.Errorf("Issues.List returned %+v, want %+v", items, want)
	}
}
