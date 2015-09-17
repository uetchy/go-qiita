package qiita

type ExpandedTemplate struct {
	ExpandedBody  string        `json:"expanded_body"`
	ExpandedTags  []AttachedTag `json:"expanded_tags"`
	ExpandedTitle string        `json:"expanded_title"`
}

type ExpandedTemplateService struct {
	client *Client
}

func (s *ExpandedTemplateService) Create() (ExpandedTemplate, error) {

}
