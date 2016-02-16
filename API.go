package gobacklog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PrintResponseJSON is bool
var PrintResponseJSON = false

// Space returns
func (c *Client) Space() (*Space, error) {
	bytes, er := c.Get("/api/v2/space", url.Values{})

	if er != nil {
		return nil, er
	}

	var space *Space
	json.Unmarshal(bytes, &space)
	return space, nil
}

// SpaceNotification /api/v2/space/notification
func (c *Client) SpaceNotification() (*SpaceNotification, error) {
	bytes, er := c.Get("/api/v2/space/notification", url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var spaceNotification *SpaceNotification
	json.Unmarshal(bytes, &spaceNotification)
	return spaceNotification, nil
}

// DiskUsage /api/v2/space/diskUsage
func (c *Client) DiskUsage() (*DiskUsage, error) {
	bytes, er := c.Get("/api/v2/space/diskUsage", url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var diskUsage *DiskUsage
	json.Unmarshal(bytes, &diskUsage)
	return diskUsage, nil
}

// Users /api/v2/users
func (c *Client) Users() (UserSlice, error) {
	bytes, er := c.Get("/api/v2/users", url.Values{})

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

// User /api/v2/users/:userId
func (c *Client) User(userID int) (*User, error) {
	endpoint := fmt.Sprintf("/api/v2/users/%d", userID)
	bytes, er := c.Get(endpoint, url.Values{})

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

// Myself returns
func (c *Client) Myself() (*User, error) {
	bytes, er := c.Get("/api/v2/users/myself", url.Values{})

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

// ProjectsWithOption returns project information.
// /api/v2/projects
func (c *Client) ProjectsWithOption(option *ProjectsOption) (ProjectSlice, error) {
	params, er := option.Values()
	if er != nil {
		return nil, er
	}

	bytes, er := c.Get("/api/v2/projects", params)

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var projs ProjectSlice
	json.Unmarshal(bytes, &projs)

	return projs, nil
}

// ProjectWithID returns project information.
// /api/v2/projects/:projectIdOrKey
func (c *Client) ProjectWithID(projectID int) (*Project, error) {
	endpoint := fmt.Sprintf("/api/v2/projects/%d", projectID)
	bytes, er := c.Get(endpoint, url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var proj *Project
	json.Unmarshal(bytes, &proj)

	return proj, nil
}

// ProjectWithKey returns project information.
// /api/v2/projects/:projectIdOrKey
func (c *Client) ProjectWithKey(projectKey string) (*Project, error) {
	if len(projectKey) == 0 {
		return nil, errors.New("invalid arguments")
	}

	endpoint := fmt.Sprintf("/api/v2/projects/%s", projectKey)
	bytes, er := c.Get(endpoint, url.Values{})

	if er != nil {
		return nil, er
	}

	if PrintResponseJSON {
		fmt.Println(string(bytes))
	}

	var proj *Project
	json.Unmarshal(bytes, &proj)

	return proj, nil
}

// Issues is
func (c *Client) Issues() (IssueSlice, error) {
	bytes, er := c.Get("/api/v2/issues", url.Values{})

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
