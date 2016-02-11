package gobacklog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PrintResponseJSON is bool
var PrintResponseJSON = false

// Client is
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	APIKey     string
}

// NewClient returns Backlog HTTP Client
func NewClient(baseURL, APIKey string) *Client {
	s := &Client{
		BaseURL:    baseURL,
		APIKey:     APIKey,
		HTTPClient: http.DefaultClient,
	}

	return s
}

func (c *Client) appendAPIKey(URL string) string {
	return URL + "?apiKey=" + c.APIKey
}

// Space returns
func (c *Client) Space() (*Space, error) {
	url := c.appendAPIKey(c.BaseURL + "/api/v2/space")

	req, _ := http.NewRequest("GET", url, nil)
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bytes, er := ioutil.ReadAll(res.Body)

	if er != nil {
		return nil, er
	}

	var space *Space
	json.Unmarshal(bytes, &space)
	return space, nil
}

// SpaceNotification /api/v2/space/notification
func (c *Client) SpaceNotification() (*SpaceNotification, error) {

	url := c.appendAPIKey(c.BaseURL + "/api/v2/space/notification")

	req, _ := http.NewRequest("GET", url, nil)
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bytes, er := ioutil.ReadAll(res.Body)

	if er != nil {
		return nil, er
	}

	fmt.Println(string(bytes))
	var spaceNotification *SpaceNotification
	json.Unmarshal(bytes, &spaceNotification)
	return spaceNotification, nil
}

// DiskUsage /api/v2/space/diskUsage
func (c *Client) DiskUsage() (*DiskUsage, error) {

	url := c.appendAPIKey(c.BaseURL + "/api/v2/space/diskUsage")

	req, _ := http.NewRequest("GET", url, nil)
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bytes, er := ioutil.ReadAll(res.Body)

	if er != nil {
		return nil, er
	}

	fmt.Println(string(bytes))
	var diskUsage *DiskUsage
	json.Unmarshal(bytes, &diskUsage)
	return diskUsage, nil
}

// Myself returns
func (c *Client) Myself() (*User, error) {
	url := c.appendAPIKey(c.BaseURL + "/api/v2/users/myself")

	req, _ := http.NewRequest("GET", url, nil)
	res, err := c.HTTPClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	bytes, er := ioutil.ReadAll(res.Body)

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var user *User
	json.Unmarshal(bytes, &user)

	return user, nil
}

// Issues is
func (c *Client) Issues() (IssueSlice, error) {
	url := c.appendAPIKey(c.BaseURL + "/api/v2/issues")

	req, _ := http.NewRequest("GET", url, nil)
	res, err := c.HTTPClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	bytes, er := ioutil.ReadAll(res.Body)

	if er != nil {
		return nil, er
	}

	// fmt.Println(string(bytes))
	var issues IssueSlice
	json.Unmarshal(bytes, &issues)

	return issues, nil
}

// IssuesWithOption is
func (c *Client) IssuesWithOption(opt *IssuesOption) (IssueSlice, error) {
	url := c.appendAPIKey(c.BaseURL + "/api/v2/issues")

	query, err := opt.ParamString()
	if err != nil {
		return nil, err
	}

	url = url + "&" + query

	req, _ := http.NewRequest("GET", url, nil)
	res, err := c.HTTPClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	bytes, er := ioutil.ReadAll(res.Body)

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var issues IssueSlice
	json.Unmarshal(bytes, &issues)

	return issues, nil
}

// IssueWithKey is
func (c *Client) IssueWithKey(issueIDOrKey string) (*Issue, error) {
	url := c.appendAPIKey(c.BaseURL + "/api/v2/issues/" + issueIDOrKey)

	req, _ := http.NewRequest("GET", url, nil)
	res, err := c.HTTPClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	bytes, er := ioutil.ReadAll(res.Body)

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	issue := &Issue{}
	json.Unmarshal(bytes, issue)

	return issue, nil
}
