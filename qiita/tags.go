package qiita

import (
	// "fmt"
	"net/http"
)

type TagsService struct {
	client *Client
}

type AttachedTag struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

type Tag struct {
	FollowersCount int    `json:"followers_count"`
	IconURL        string `json:"icon_url"`
	Id             string `json:"id"`
	ItemsCount     int    `json:"items_count"`
}

func (s *TagsService) List(opt *ListOptions) ([]Tag, *http.Response, error) {
	u, err := addOptions("/tags", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	tags := new([]Tag)
	resp, err := s.client.Do(req, tags)
	if err != nil {
		return nil, nil, err
	}

	return *tags, resp, err
}
