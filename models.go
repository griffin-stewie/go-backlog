package gobacklog

import (
	"time"
)

// Issue represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string]"
type Issue struct {
	ID          *int      `json:"id,omitempty"`
	ProjectID   *int      `json:"projectId,omitempty"`
	IssueKey    *string   `json:"issueKey,omitempty"`
	KeyID       *int      `json:"keyId,omitempty"`
	IssueType   IssueType `json:"issueType,omitempty"`
	Summary     *string   `json:"summary,omitempty"`
	Description *string   `json:"description,omitempty"`
	// Resolution interface{} `json:"resolution,omitempty"`
	// Priority interface{} `json:"priority,omitempty"`
	// Status interface{} `json:"status,omitempty"`
	// Assignee interface{} `json:"assignee,omitempty"`
	// Category interface{} `json:"category,omitempty"`
	// Versions interface{} `json:"versions,omitempty"`
	// Milestone interface{} `json:"milestone,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	DueDate   *time.Time `json:"dueDate,omitempty"`
	// EstimatedHours interface{} `json:"estimatedHours,omitempty"`
	// ActualHours interface{} `json:"actualHours,omitempty"`
	ParentIssueID *int `json:"parentIssueId,omitempty"`
	// CreatedUser interface{} `json:"createdUser,omitempty"`
	Created *time.Time `json:"created,omitempty"`

	// UpdatedUser interface{} `json:"updatedUser,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	// CustomFields []interface{} `json:"customFields,omitempty"`
	// Attachments interface{} `json:"attachments,omitempty"`
	// SharedFiles interface{} `json:"sharedFiles,omitempty"`
	// Stars interface{} `json:"stars,omitempty"`
}

// Issues represents
type Issues struct {
	Issues IssueSlice
}

// IssueType represents
type IssueType struct {
	ID           *int    `json:"id,omitempty"`
	ProjectID    *int    `json:"projectId,omitempty"`
	Name         *string `json:"name,omitempty"`
	Color        *string `json:"color,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

// User represents
type User struct {
	ID          *int    `json:"id,omitempty"`
	UserID      *string `json:"userId,omitempty"`
	Name        *string `json:"name,omitempty"`
	RoleType    *int    `json:"roleType,omitempty"`
	Lang        *string `json:"lang,omitempty"`
	MailAddress *string `json:"mailAddress,omitempty"`
}

// Space represents
type Space struct {
	SpaceKey           *string    `json:"spaceKey,omitempty"`
	Name               *string    `json:"name,omitempty"`
	OwnerID            *int       `json:"ownerId,omitempty"`
	Lang               *string    `json:"lang,omitempty"`
	Timezone           *string    `json:"timezone,omitempty"`
	ReportSendTime     *string    `json:"reportSendTime,omitempty"`
	TextFormattingRule *string    `json:"textFormattingRule,omitempty"`
	Created            *time.Time `json:"created,omitempty"`
	Updated            *time.Time `json:"updated,omitempty"`
}

// SpaceNotification represents
type SpaceNotification struct {
	Content *string    `json:"content,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}

// DiskUsage represents
type DiskUsage struct {
	Capacity    *int                 `json:"capacity,omitempty"`
	Issue       *int                 `json:"issue,omitempty"`
	Wiki        *int                 `json:"wiki,omitempty"`
	File        *int                 `json:"file,omitempty"`
	Subversion  *int                 `json:"subversion,omitempty"`
	Git         *int                 `json:"git,omitempty"`
	PullRequest *int                 `json:"pullRequest,omitempty"`
	Details     DiskUsageDetailSlice `json:"details,omitempty"`
}

// DiskUsageDetail represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string]"
type DiskUsageDetail struct {
	ProjectID   *int `json:"projectId,omitempty"`
	Issue       *int `json:"issue,omitempty"`
	Wiki        *int `json:"wiki,omitempty"`
	File        *int `json:"file,omitempty"`
	Subversion  *int `json:"subversion,omitempty"`
	Git         *int `json:"git,omitempty"`
	PullRequest *int `json:"pullRequest,omitempty"`
}

// Total returns total disk usage, byte unit.
func (d *DiskUsageDetail) Total() int {
	total := *d.Issue +
		*d.Wiki +
		*d.File +
		*d.Subversion +
		*d.Git +
		*d.PullRequest
	return total
}
