package gobacklog

import (
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
