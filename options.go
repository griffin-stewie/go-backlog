package gobacklog

import (
	"net/url"

	"github.com/google/go-querystring/query"
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

// SortOrder enum
type SortOrder string

// SortOrder enum
const (
	Ascending  SortOrder = "asc"
	Descending           = "desc"
)

// IssuesOption represents
type IssuesOption struct {
	IDs         []int         `url:"id[],omitempty"`
	ProjectIDs  []int         `url:"projectId[],omitempty"`
	AssigneeIDs []int         `url:"assigneeId[],omitempty"`
	Statuses    []IssueStatus `url:"statusId[],omitempty"`
	SharedFile  bool          `url:"sharedFile,omitempty"`
	Count       int           `url:"count,omitempty"`
	Offset      int           `url:"offset,omitempty"`
	Keyword     string        `url:"keyword,omitempty"`
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

// ActivitiesOption represents
type ActivitiesOption struct {
	ActivityTypeIDs []int  `url:"activityTypeId[],omitempty"`
	MinID           int    `url:"minId,omitempty"`
	MaxID           int    `url:"maxId,omitempty"`
	Count           int    `url:"count,omitempty"`
	Order           string `url:"order,omitempty"`
}

// ParamString returns
func (c *ActivitiesOption) ParamString() (string, error) {
	values, err := query.Values(c)
	if err != nil {
		return "", err
	}
	return values.Encode(), nil
}

// Values returns
func (c *ActivitiesOption) Values() (url.Values, error) {
	return query.Values(c)
}
