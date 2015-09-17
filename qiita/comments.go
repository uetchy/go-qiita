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

func (s *CommentsService) Get() (*Comment, error) {
	req, err := s.client.NewRequest("GET", "comments/1", nil)
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
