package qiita

import (
	"net/http"
)

type AuthenticatedUserService struct {
	client *Client
}

// Get a user associated to the current access token.
func (s *AuthenticatedUserService) Show() (*User, error) {
	req, err := s.client.NewRequest("GET", "authenticated_user", nil)
	if err != nil {
		return nil, err
	}
	user := new(User)
	_, err = s.client.Do(req, user)
	if err != nil {
		return nil, err
	}

	return user, err
}

// List the authenticated user's items in newest order
func (s *AuthenticatedUserService) Items(opt *ListOptions) ([]Item, *http.Response, error) {
	u, err := addOptions("authenticated_user/items", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	items := new([]Item)
	resp, err := s.client.Do(req, items)
	if err != nil {
		return nil, nil, err
	}

	return *items, resp, err
}
