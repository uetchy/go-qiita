package qiita

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

// List tags in newest order.
func (s *TagsService) List(opt *ListOptions) ([]Tag, error) {
	u, err := addOptions("/tags", opt)
	if err != nil {
		return nil, err
	}
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

// TODO: Get a tag.
func (s *TagsService) Get(tagId string) (*Tag, error) {
	req, err := s.client.NewRequest("GET", "tags/"+tagId, nil)
	if err != nil {
		return nil, err
	}
	tag := new(Tag)
	_, err = s.client.Do(req, tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// TODO: List tagged items in recently-tagged order.
func (s *TagsService) TaggedItems(tagId string) ([]Item, error) {
	req, err := s.client.NewRequest("GET", "tags/"+tagId+"/items", nil)
	if err != nil {
		return nil, err
	}
	items := new([]Item)
	_, err = s.client.Do(req, items)
	if err != nil {
		return nil, err
	}
	return *items, nil
}

// TODO: Check if you are following a tag or not.
func (s *TagsService) IsFollowed(tagId string) error {
	req, err := s.client.NewRequest("GET", "tags/"+tagId+"/following", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Follow a tag.
func (s *TagsService) Follow(tagId string) error {
	req, err := s.client.NewRequest("PUT", "tags/"+tagId+"/following", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Unfollow a tag.
func (s *TagsService) Unfollow(tagId string) error {
	req, err := s.client.NewRequest("DELETE", "tags/"+tagId+"/following", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}
