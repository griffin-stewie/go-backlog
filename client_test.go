package gobacklog

import (
	"testing"
)

func TestNewClientAdjustedBaseURL(t *testing.T) {
	tab := []struct {
		BaseURL     string
		AdjustedURL string
	}{
		{
			BaseURL:     "http://example.com/",
			AdjustedURL: "http://example.com",
		},
		{
			BaseURL:     "http://example.com",
			AdjustedURL: "http://example.com",
		},
		{
			BaseURL:     "",
			AdjustedURL: "",
		},
	}
	for _, v := range tab {
		c := NewClient(v.BaseURL, "")
		if c.BaseURL != v.AdjustedURL {
			t.Errorf(`NewClient(%q, ""): BaseURL = %q; want %q`, v.BaseURL, c.BaseURL, v.AdjustedURL)
		}
	}
}
