package qiita

import (
	"time"
)

type CommentsService struct {
	client *Client
}

type Comment struct {
	RenderedBody string    `json:"rendered_body"`
	Body         string    `json:"body"`
	CreatedAt    time.Time `json:"created_at"`
	Id           string    `json:"id"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `json:"user"`
}

// TODO: Post a comment on an item.
func (s *CommentsService) Create(itemId string, body string) error {
	req, err := s.client.NewRequest("POST", "items/"+itemId+"/comments", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: List comments on an item in newest order.
func (s *CommentsService) List(itemId string) ([]Comment, error) {
	req, err := s.client.NewRequest("GET", "items/"+itemId+"/comments", nil)
	if err != nil {
		return nil, err
	}
	comments := new([]Comment)
	_, err = s.client.Do(req, comments)
	if err != nil {
		return nil, err
	}
	return *comments, err
}

// TODO: Get a comment.
func (s *CommentsService) Get(commentId string) (*Comment, error) {
	req, err := s.client.NewRequest("GET", "comments/"+commentId, nil)
	if err != nil {
		return nil, err
	}
	comment := new(Comment)
	_, err = s.client.Do(req, comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// TODO: Update a comment.
func (s *CommentsService) Edit(commentId string, comment *Comment) error {
	req, err := s.client.NewRequest("PATCH", "comments/"+commentId, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Delete a comment.
func (s *CommentsService) Delete(commentId string) error {
	req, err := s.client.NewRequest("DELETE", "comments/"+commentId, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Send a thank to a comment.
func (s *CommentsService) AddThank(id string) error {
	req, err := s.client.NewRequest("PUT", "comments/"+id+"/thank", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Delete a thank from a comment.
func (s *CommentsService) DeleteThank(id string) error {
	req, err := s.client.NewRequest("DELETE", "comments/"+id+"/thank", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}
