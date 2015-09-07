package qiita

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func TestAddOptions(t *testing.T) {
	want := "items?page=2&per_page=100&query=test"
	got, err := addOptions("items", &ItemsListOptions{"test", ListOptions{Page: 2, PerPage: 100}})
	if got != want || err != nil {
		t.Fatalf("%v: %v != %v", err, got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)
	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Fatalf("%v != %v", got, want)
	}
}
