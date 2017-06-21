package gobacklog

import (
	"fmt"
	"net/http"
	"net/url"
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

func TestResolvingURL(t *testing.T) {
	apiKey := "apikey"
	tab := []struct {
		Endpoint    string
		BaseURL     string
		Params      url.Values
		AdjustedURL string
	}{
		{
			Endpoint:    "/api/v2/space",
			BaseURL:     "http://example.com/",
			Params:      url.Values{},
			AdjustedURL: "http://example.com/api/v2/space?apiKey=apikey",
		},
		{
			Endpoint:    "/api/v2/space",
			BaseURL:     "http://example.com",
			Params:      url.Values{},
			AdjustedURL: "http://example.com/api/v2/space?apiKey=apikey",
		},
	}
	for _, v := range tab {
		c := NewClient(v.BaseURL, apiKey)
		result := c.buildURLWithValues(v.BaseURL, v.Endpoint, v.Params)
		if result != v.AdjustedURL {
			req, requestErr := http.NewRequest("GET",
				result,
				nil,
			)
			if requestErr != nil {
				t.Error(requestErr)
			}
			fmt.Println(req.URL)
			t.Errorf(`Result = %q; BaseURL = %q; want %q`, result, c.BaseURL, v.AdjustedURL)
		}
	}
}
