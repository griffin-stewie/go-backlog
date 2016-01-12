package gobacklog

import (
	"testing"
)

func TestBuildQueryPramString(t *testing.T) {
	input := &IssuesOption{
		AssigneeIds: []int{999999999},
		Statuses:    []IssueStatus{1, 2, 3},
	}

	expected := "assigneeId%5B%5D=999999999&statusId%5B%5D=1&statusId%5B%5D=2&statusId%5B%5D=3"
	actual, err := input.ParamString()

	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("\n resutl: %v\nexpected: %v", actual, expected)
	}
}
