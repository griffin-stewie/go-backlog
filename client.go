package gobacklog

import (
	"bytes"
	"encoding/json"
	"path"
	// "fmt"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Verbose is bool
var Verbose = false

// BacklogErrorResponse is error model
type BacklogErrorResponse struct {
	Errors BacklogErrorSlice `json:"errors"`
}

// BacklogError is error model
// +gen * slice:"Where,Count,SortBy,GroupBy[string],Select[string]"
type BacklogError struct {
	Message  string `json:"message,omitempty"`
	Code     int    `json:"code,omitempty"`
	MoreInfo string `json:"moreInfo,omitempty"`
}

// error returns Errors
func (e *BacklogErrorResponse) error() error {
	if len(e.Errors) == 0 {
		return nil
	}

	s := e.Errors.SelectString(func(b *BacklogError) string {
		return b.Message
	})

	return errors.New(strings.Join(s, ", "))
}

// HTTP interface of HTTP METHODS's methods
type HTTP interface {
	Get()
	Post()
	Put()
	Delete()
}

// Client is
type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	APIKey     string
}

// NewClient returns Backlog HTTP Client
func NewClient(baseURL *url.URL, APIKey string) *Client {
	s := &Client{
		BaseURL: baseURL,
		APIKey:  APIKey,
	}

	return s
}

// Get GET method
func (c *Client) Get(endpoint string, params url.Values) ([]byte, error) {
	return c.execute("GET", endpoint, params)
}

// Post POST method
func (c *Client) Post(endpoint string, params url.Values) ([]byte, error) {
	return c.execute("POST", endpoint, params)
}

// Put PUT method
func (c *Client) Put(endpoint string, params url.Values) ([]byte, error) {
	return c.execute("PUT", endpoint, params)
}

// Delete DELETE method
func (c *Client) Delete(endpoint string, params url.Values) ([]byte, error) {
	return c.execute("DELETE", endpoint, params)
}

func (c *Client) composeURL(pathStr string, params url.Values) string {
	copiedURL := *c.BaseURL
	copiedURL.Path = path.Join(copiedURL.Path, pathStr)
	params.Set("apiKey", c.APIKey)
	copiedURL.RawQuery = params.Encode()
	return copiedURL.String()
}

func (c *Client) parseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if Verbose {
			log.Println(err)
		}

		return []byte(``), err
	}

	if Verbose {
		log.Printf("[DEBUG] resp: %#+v", resp)
		log.Printf("[DEBUG] body: %v", string(body))
	}

	if resp.StatusCode != 200 {
		var er BacklogErrorResponse
		json.Unmarshal(body, &er)
		return []byte(``), er.error()
	}

	return body, nil
}

func (c *Client) execute(method, endpoint string, params url.Values) ([]byte, error) {
	resp, err := c.executeReturnsResponse(method, endpoint, params)

	if err != nil {
		return []byte(``), err
	}

	return c.parseBody(resp)
}

func (c *Client) executeReturnsResponse(method, endpoint string, params url.Values) (resp *http.Response, err error) {
	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}

	var (
		req        *http.Request
		requestErr error
	)

	if method != "GET" {
		req, requestErr = http.NewRequest(method,
			c.composeURL(endpoint, url.Values{}),
			bytes.NewBufferString(params.Encode()),
		)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, requestErr = http.NewRequest(method,
			c.composeURL(endpoint, params),
			nil,
		)
	}

	if requestErr != nil {
		panic(requestErr)
	}

	return c.HTTPClient.Do(req)
}
