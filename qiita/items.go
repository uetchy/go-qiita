package qiita

import (
	"time"
)

type ItemsService struct {
	client *Client
}

type Item struct {
	RenderedBody string        `json:"rendered_body"`
	Body         string        `json:"body"`
	CoEditing    bool          `json:"coediting"`
	CreatedAt    time.Time     `json:"created_at"`
	Id           string        `json:"id"`
	Private      bool          `json:"private"`
	Tags         []AttachedTag `json:"tags"`
	Title        string        `json:"title"`
	UpdatedAt    time.Time     `json:"updated_at"`
	URL          string        `json:"url"`
	User         User          `json:"user"`
}

type ItemsListOptions struct {
	Query string `url:"query,omitempty"`
	ListOptions
}

// TODO: Create item
func (s *ItemsService) Create(item *Item) error {
	req, err := s.client.NewRequest("POST", "items", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: List items.
func (s *ItemsService) List(query string, opt *ListOptions) ([]Item, error) {
	u, _ := addOptions("items", &ItemsListOptions{query, *opt})
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

// TODO: Get items
func (s *ItemsService) Get(itemId string) (*Item, error) {
	req, err := s.client.NewRequest("GET", "items/"+itemId, nil)
	if err != nil {
		return nil, err
	}
	item := new(Item)
	_, err = s.client.Do(req, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// TODO: Update items
func (s *ItemsService) Edit(itemId string) error {
	req, err := s.client.NewRequest("PATCH", "items/"+itemId, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Delete item
func (s *ItemsService) Delete(itemId string) error {
	req, err := s.client.NewRequest("POST", "items/"+itemId, nil)
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
func (s *ItemsService) Comments(itemId string) ([]Comment, error) {
	req, err := s.client.NewRequest("GET", "items/"+itemId+"/comments", nil)
	if err != nil {
		return nil, err
	}
	comments := new([]Comment)
	_, err = s.client.Do(req, comments)
	if err != nil {
		return nil, err
	}
	return *comments, nil
}

// TODO: Post a comment on an item.
func (s *ItemsService) AddComment(itemId string) error {
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

// TODO: List users who stocked an item in recent-stocked order.
func (s *ItemsService) Stockers(itemId string, opt *ListOptions) ([]User, error) {
	req, err := s.client.NewRequest("GET", "items/"+itemId+"/stockers", nil)
	if err != nil {
		return nil, err
	}
	stockers := new([]User)
	_, err = s.client.Do(req, stockers)
	if err != nil {
		return nil, err
	}
	return *stockers, nil
}

// TODO: Like an item
func (s *ItemsService) Like(itemId string) error {
	req, err := s.client.NewRequest("PUT", "items/"+itemId+"/like", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Unlike an item
func (s *ItemsService) Unlike(itemId string) error {
	req, err := s.client.NewRequest("DELETE", "items/"+itemId+"/like", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Check if stocked an item
func (s *ItemsService) IsStocked(itemId string) (*bool, error) {
	req, err := s.client.NewRequest("GET", "items/"+itemId+"/stock", nil)
	if err != nil {
		return nil, err
	}
	isStocked := new(bool)
	_, err = s.client.Do(req, isStocked)
	if err != nil {
		return nil, err
	}
	return isStocked, nil
}

// TODO: Stock an item
func (s *ItemsService) Stock(itemId string) error {
	req, err := s.client.NewRequest("PUT", "items/"+itemId+"/stock", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Unstock an item
func (s *ItemsService) Unstock(itemId string) error {
	req, err := s.client.NewRequest("DELETE", "items/"+itemId+"/stock", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Add a tag to an item (only available on Qiita:Team)
func (s *ItemsService) AddTag(itemId string, tagName string) error {
	req, err := s.client.NewRequest("POST", "items/"+itemId+"/taggings", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Remove a tag from an item (only available on Qiita:Team)
func (s *ItemsService) RemoveTag(itemId string, tagId string) error {
	req, err := s.client.NewRequest("DELETE", "items/"+itemId+"/taggings/"+tagId, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}
