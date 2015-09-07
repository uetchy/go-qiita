package qiita

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
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
        "created_at": "2000-01-01T00:00:00+00:00",
        "id": "4bd431809afb1bb99e4f",
        "private": false,
        "tags": [
          {
            "name": "Ruby",
            "versions": [
              "0.0.1"
            ]
          }
        ],
        "title": "Example title",
        "updated_at": "2000-01-01T00:00:00+00:00",
        "url": "https://qiita.com/yaotti/items/4bd431809afb1bb99e4f",
        "user": {
          "description": "Hello, world.",
          "facebook_id": "yaotti",
          "followees_count": 100,
          "followers_count": 200,
          "github_login_name": "yaotti",
          "id": "yaotti",
          "items_count": 300,
          "linkedin_id": "yaotti",
          "location": "Tokyo, Japan",
          "name": "Hiroshige Umino",
          "organization": "Increments Inc",
          "permanent_id": 1,
          "profile_image_url": "https://si0.twimg.com/profile_images/2309761038/1ijg13pfs0dg84sk2y0h_normal.jpeg",
          "twitter_screen_name": "yaotti",
          "website_url": "http://yaotti.hatenablog.com"
        }
      }
    ]`)
	})

	opt := &ItemsListOptions{
		"query",
		ListOptions{Page: 1, PerPage: 2},
	}
	items, _, err := client.Items.List(opt)

	if err != nil {
		t.Errorf("Issues.List returned error: %v", err)
	}

	want := []Item{
		{
			Body:         "# Example",
			RenderedBody: "<h1>Example</h1>",
			CoEditing:    false,
			CreatedAt:    time.Date(2000, time.January, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
			Id:           "4bd431809afb1bb99e4f",
			Private:      false,
			Tags: []AttachedTag{{
				"Ruby",
				[]string{
					"0.0.1",
				},
			}},
			Title:     "Example title",
			UpdatedAt: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
			URL:       "https://qiita.com/yaotti/items/4bd431809afb1bb99e4f",
			User: User{
				Description:       "Hello, world.",
				FacebookId:        "yaotti",
				FolloweesCount:    100,
				FollowersCount:    200,
				GithubLoginName:   "yaotti",
				Id:                "yaotti",
				ItemsCount:        300,
				LinkedinId:        "yaotti",
				Location:          "Tokyo, Japan",
				Name:              "Hiroshige Umino",
				Organization:      "Increments Inc",
				PermanentId:       1,
				ProfileImageURL:   "https://si0.twimg.com/profile_images/2309761038/1ijg13pfs0dg84sk2y0h_normal.jpeg",
				TwitterScreenName: "yaotti",
				WebsiteURL:        "http://yaotti.hatenablog.com",
			},
		},
	}
	if !reflect.DeepEqual(items, want) {
		t.Errorf("Issues.List returned %+v, want %+v", items, want)
	}
}