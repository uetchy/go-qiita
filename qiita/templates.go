package qiita

type TemplatesService struct {
	client *Client
}

type Template struct {
	Name          string        `json:"name"`
	Body          string        `json:"body"`
	Id            string        `json:"id"`
	Title         string        `json:"title"`
	Tags          []Tag         `json:"tags"`
	ExpandedBody  string        `json:"expanded_body"`
	ExpandedTags  []AttachedTag `json:"expanded_tags"`
	ExpandedTitle string        `json:"expanded_title"`
}

// TODO: Create a new template.
func (s *TemplatesService) Create(templateId string) error {
	req, err := s.client.NewRequest("POST", "templates", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: List templates in a team.
func (s *TemplatesService) List(opt *ListOptions) ([]Team, error) {
	req, err := s.client.NewRequest("GET", "templates", nil)
	if err != nil {
		return nil, err
	}
	teams := new([]Team)
	_, err = s.client.Do(req, teams)
	if err != nil {
		return nil, err
	}
	return *teams, nil
}

// TODO: Get a template.
func (s *TemplatesService) Get(templateId string) (*Team, error) {
	req, err := s.client.NewRequest("GET", "templates/"+templateId, nil)
	if err != nil {
		return nil, err
	}
	team := new(Team)
	_, err = s.client.Do(req, team)
	if err != nil {
		return nil, err
	}
	return team, nil
}

// TODO: Update a template.
func (s *TemplatesService) Edit(templateId string) error {
	req, err := s.client.NewRequest("PATCH", "templates/"+templateId, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Delete a template.
func (s *TemplatesService) Delete(templateId string) error {
	req, err := s.client.NewRequest("DELETE", "templates/"+templateId, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}
