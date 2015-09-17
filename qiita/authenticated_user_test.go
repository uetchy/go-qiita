package qiita

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAuthenticatedUserService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/authenticated_user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
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
      "website_url": "http://yaotti.hatenablog.com",
      "image_monthly_upload_limit": 1048576,
      "image_monthly_upload_remaining": 524288,
      "team_only": false
    }`)
	})

	items, err := client.AuthenticatedUser.Get()

	if err != nil {
		t.Errorf("Issues.List returned error: %v", err)
	}

	want := &User{
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
	}
	if !reflect.DeepEqual(items, want) {
		t.Errorf("Issues.List returned %+v, want %+v", items, want)
	}
}

func TestAuthenticatedUserService_Items(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/authenticated_user/items", func(w http.ResponseWriter, r *http.Request) {
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
	items, _, err := client.AuthenticatedUser.Items(opt)

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
			Title: "Example title",
			URL:   "https://qiita.com/yaotti/items/4bd431809afb1bb99e4f",
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
