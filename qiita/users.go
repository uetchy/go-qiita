package qiita

import (
	"fmt"
)

type UsersService struct {
	client *Client
}

type User struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	FacebookId        string `json:"facebook_id"`
	FolloweesCount    int    `json:"followees_count"`
	FollowersCount    int    `json:"followers_count"`
	GithubLoginName   string `json:"github_login_name"`
	Id                string `json:"id"`
	ItemsCount        int    `json:"items_count"`
	LinkedinId        string `json:"linkedin_id"`
	Location          string `json:"location"`
	Organization      string `json:"organization"`
	PermanentId       int    `json:"permanent_id"`
	ProfileImageURL   string `json:"profile_image_url"`
	TwitterScreenName string `json:"twitter_screen_name"`
	WebsiteURL        string `json:"website_url"`
}

// TODO: List users in newest order.
func (s *UsersService) List(opt *ItemsListOptions) ([]Item, error) {
	req, err := s.client.NewRequest("GET", "users", nil)
	if err != nil {
		return nil, err
	}
	items := new([]Item)
	_, err = s.client.Do(req, items)
	if err != nil {
		return nil, err
	}
	return *items, err
}

// TODO: Get a user.
func (s *UsersService) Get(userId string) (*Item, error) {
	req, err := s.client.NewRequest("GET", "users/"+userId, nil)
	if err != nil {
		return nil, err
	}
	item := new(Item)
	_, err = s.client.Do(req, item)
	if err != nil {
		return nil, err
	}
	return item, err
}

// TODO: List users a user is following.
func (s *UsersService) Followees(userId string) ([]User, error) {
	req, err := s.client.NewRequest("GET", "users/"+userId+"/followees", nil)
	if err != nil {
		return nil, err
	}
	users := new([]User)
	_, err = s.client.Do(req, users)
	if err != nil {
		return nil, err
	}
	return *users, err
}

// TODO: List users who are following a user.
func (s *UsersService) Followers(userId string) ([]User, error) {
	req, err := s.client.NewRequest("GET", "users/"+userId+"/followers", nil)
	if err != nil {
		return nil, err
	}
	users := new([]User)
	_, err = s.client.Do(req, users)
	if err != nil {
		return nil, err
	}
	return *users, err
}

// TODO: Check if the current user is following a user.
func (s *UsersService) IsFollowed(userId string) error {
	// req, err := s.client.NewRequest("GET", "users/"+userId+"/following", nil)
	return nil
}

// TODO: Follow a user.
func (s *UsersService) Follow(userId string) error {
	req, err := s.client.NewRequest("PUT", "users/"+userId+"/following", nil)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil || resp.StatusCode != 204 {
		return err
	}
	return nil
}

// TODO: Unfollow a user.
func (s *UsersService) Unfollow(userId string) error {
	req, err := s.client.NewRequest("DELETE", "users/"+userId+"/following", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// List a user's items in newest order.
func (s *UsersService) Items(userId string, opt *ListOptions) ([]Item, error) {
	u := fmt.Sprintf("users/%s/items", userId)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	items := new([]Item)
	_, err = s.client.Do(req, items)
	if err != nil {
		return nil, err
	}
	return *items, err
}

// List a user's stocked items in recently-stocked order.
func (s *UsersService) Stocks(userId string, opt *ListOptions) ([]Item, error) {
	u := fmt.Sprintf("users/%s/stocks", userId)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	items := new([]Item)
	_, err = s.client.Do(req, items)
	if err != nil {
		return nil, err
	}
	return *items, err
}

// TODO: List tags a user is following to in recently-tagged order.
func (s *UsersService) FolloingTags(userId string) ([]Tag, error) {
	u := fmt.Sprintf("users/%s/following_tags", userId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	tags := new([]Tag)
	_, err = s.client.Do(req, tags)
	if err != nil {
		return nil, err
	}
	return *tags, err
}
