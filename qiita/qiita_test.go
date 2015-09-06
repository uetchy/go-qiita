package qiita

import (
  "testing"
)

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
