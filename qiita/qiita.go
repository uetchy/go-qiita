package qiita

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

const (
	defaultBaseURL = "https://qiita.com/api/v2/"
	userAgent      = "go-qiita"
)

type ListOptions struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

type Response struct {
	*http.Response

	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int
}

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

	AuthenticatedUser *AuthenticatedUserService
	Comments          *CommentsService
	ExpandedTemplate  *ExpandedTemplateService
	Items             *ItemsService
	Projects          *ProjectsService
	Tags              *TagsService
	Teams             *TeamsService
	Templates         *TemplatesService
	Users             *UsersService
}

// Create http query string from map
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// Return new client instance
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.AuthenticatedUser = &AuthenticatedUserService{client: c}
	c.Comments = &CommentsService{client: c}
	c.ExpandedTemplate = &ExpandedTemplateService{client: c}
	c.Items = &ItemsService{client: c}
	c.Projects = &ProjectsService{client: c}
	c.Tags = &TagsService{client: c}
	c.Teams = &TeamsService{client: c}
	c.Templates = &TemplatesService{client: c}
	c.Users = &UsersService{client: c}
	return c
}

// Return new http request
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do http request
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err
}
