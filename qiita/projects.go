package qiita

import (
	"time"
)

type ProjectsService struct {
	client *Client
}

type Project struct {
	RenderedBody string    `json:"rendered_body"`
	Archived     bool      `json:"archived"`
	Body         string    `json:"body"`
	CreatedAt    time.Time `json:"created_at"`
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TODO: Create a new project.
func (s *ProjectsService) Create() error {
	req, err := s.client.NewRequest("POST", "projects", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: List projects in newest order.
func (s *ProjectsService) List(opt *ListOptions) ([]Project, error) {
	req, err := s.client.NewRequest("GET", "projects", nil)
	if err != nil {
		return nil, err
	}
	projects := new([]Project)
	_, err = s.client.Do(req, projects)
	if err != nil {
		return nil, err
	}
	return *projects, nil
}

// TODO: Get a project.
// https://qiita.com/api/v2/docs#get-apiv2projectsproject_id
func (s *ProjectsService) Get(projectId string) (*Project, error) {
	req, err := s.client.NewRequest("GET", "projects/"+projectId, nil)
	if err != nil {
		return nil, err
	}
	project := new(Project)
	_, err = s.client.Do(req, project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// TODO: Update a project.
// https://qiita.com/api/v2/docs#patch-apiv2projectsproject_id
func (s *ProjectsService) Edit(projectId string) error {
	req, err := s.client.NewRequest("PATCH", "projectss/"+projectId, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Delete a project.
// https://qiita.com/api/v2/docs#delete-apiv2projectsproject_id
func (s *ProjectsService) Delete(projectId string) error {
	req, err := s.client.NewRequest("DELETE", "projectss/"+projectId, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}
