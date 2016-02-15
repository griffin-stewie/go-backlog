package gobacklog

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

// IssueStatus enum
type IssueStatus int

// IssueStatus enum
const (
	Open IssueStatus = iota + 1
	InProgress
	Resolved
	Closed
)

// IssuesOption represents
type IssuesOption struct {
	IDs         []int         `url:"id[],omitempty"`
	ProjectIDs  []int         `url:"projectId[],omitempty"`
	AssigneeIds []int         `url:"assigneeId[],omitempty"`
	Statuses    []IssueStatus `url:"statusId[],omitempty"`
}

// ParamString returns
func (c *IssuesOption) ParamString() (string, error) {
	values, err := query.Values(c)
	if err != nil {
		return "", err
	}
	return values.Encode(), nil
}

// Values returns
func (c *IssuesOption) Values() (url.Values, error) {
	return query.Values(c)
}

// ProjectsOption represents
type ProjectsOption struct {
	Archived bool `url:"archived,omitempty"`
	All      bool `url:"all,omitempty"`
}

// Values returns
func (c *ProjectsOption) Values() (url.Values, error) {
	return query.Values(c)
}
