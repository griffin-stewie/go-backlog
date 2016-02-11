package gobacklog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PrintResponseJSON is bool
var PrintResponseJSON = false

// Space returns
func (c *Client) Space() (*Space, error) {
	bytes, er := c.Get("/api/v2/space", map[string]string{})

	if er != nil {
		return nil, er
	}

	var space *Space
	json.Unmarshal(bytes, &space)
	return space, nil
}

// SpaceNotification /api/v2/space/notification
func (c *Client) SpaceNotification() (*SpaceNotification, error) {
	bytes, er := c.Get("/api/v2/space/notification", map[string]string{})

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
	bytes, er := c.Get("/api/v2/space/diskUsage", map[string]string{})

	if er != nil {
		return nil, er
	}

	fmt.Println(string(bytes))
	var diskUsage *DiskUsage
	json.Unmarshal(bytes, &diskUsage)
	return diskUsage, nil
}

// Users /api/v2/users
func (c *Client) Users() (UserSlice, error) {
	bytes, er := c.Get("/api/v2/users", map[string]string{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var users UserSlice
	json.Unmarshal(bytes, &users)

	return users, nil
}

// Myself returns
func (c *Client) Myself() (*User, error) {
	bytes, er := c.Get("/api/v2/users/myself", map[string]string{})

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
	bytes, er := c.Get("/api/v2/issues", map[string]string{})

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

	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}

	query, err := opt.ParamString()
	if err != nil {
		return nil, err
	}

	url = url + "&" + query

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

	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}

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

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	issue := &Issue{}
	json.Unmarshal(bytes, issue)

	return issue, nil
}
