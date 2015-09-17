package qiita

type TeamsService struct {
	client *Client
}

type Team struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	Active bool   `json:"active"`
}

// TODO: List teams the user belongs to in newest order.
func (s *TeamsService) List() ([]Team, error) {
	req, err := s.client.NewRequest("GET", "teams", nil)
	if err != nil {
		return nil, err
	}
	teams := new([]Team)
	_, err = s.client.Do(req, teams)
	if err != nil {
		return nil, err
	}
	return *teams, err
}
