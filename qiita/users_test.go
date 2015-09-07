package qiita

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	// "time"
)

func TestUsersService_Stocks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/yaotti/stocks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
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
        "tags": [
          {
            "name": "Ruby",
            "versions": [
              "0.0.1"
            ]
          }
        ],
        "title": "Example title",
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

	opt := &ListOptions{Page: 1, PerPage: 2}
	items, _, err := client.Users.Stocks("yaotti", opt)

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
			Tags: []AttachedTag{{
				"Ruby",
				[]string{
					"0.0.1",
				},
			}},
			Title:     "Example title",
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
