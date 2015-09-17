package qiita

type ProjectService struct {
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

func (s *ProjectService) List(opt *ListOptions) ([]Project, error) {

}

func (s *ProjectService) Create() error {

}

func (s *ProjectService) Delete() error {

}

func (s *ProjectService) Get() error {

}

func (s *ProjectService) Edit() error {

}
