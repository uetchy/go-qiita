package qiita

type ExpandedTemplate struct {
	ExpandedBody  string        `json:"expanded_body"`
	ExpandedTags  []AttachedTag `json:"expanded_tags"`
	ExpandedTitle string        `json:"expanded_title"`
}

type ExpandedTemplateService struct {
	client *Client
}

// TODO: Get a template where its variables are expanded.
func (s *ExpandedTemplateService) Create(body string, tags []Tag, title string) (*ExpandedTemplate, error) {
	req, err := s.client.NewRequest("POST", "expanded_templates", nil)
	if err != nil {
		return nil, err
	}
	expandedTemplate := new(ExpandedTemplate)
	_, err = s.client.Do(req, expandedTemplate)
	if err != nil {
		return nil, err
	}
	return expandedTemplate, nil
}
