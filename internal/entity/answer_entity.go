package entity

import "time"

const (
	AnswerSearchOrderByDefault = "default"
	AnswerSearchOrderByTime    = "updated"
	AnswerSearchOrderByVote    = "vote"

	AnswerStatusAvailable = 1
	AnswerStatusDeleted   = 10
)

var AdminAnswerSearchStatus = map[string]int{
	"available": AnswerStatusAvailable,
	"deleted":   AnswerStatusDeleted,
}

// Answer answer
type Answer struct {
	ID             string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt      time.Time `xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt      time.Time `xorm:"updated_at TIMESTAMP"`
	QuestionID     string    `xorm:"not null default 0 BIGINT(20) question_id"`
	UserID         string    `xorm:"not null default 0 BIGINT(20) INDEX user_id"`
	LastEditUserID string    `xorm:"not null default 0 BIGINT(20) last_edit_user_id"`
	OriginalText   string    `xorm:"not null MEDIUMTEXT original_text"`
	ParsedText     string    `xorm:"not null MEDIUMTEXT parsed_text"`
	Status         int       `xorm:"not null default 1 INT(11) status"`
	Accepted       int       `xorm:"not null default 1 INT(11) adopted"`
	CommentCount   int       `xorm:"not null default 0 INT(11) comment_count"`
	VoteCount      int       `xorm:"not null default 0 INT(11) vote_count"`
	RevisionID     string    `xorm:"not null default 0 BIGINT(20) revision_id"`
}

type AnswerSearch struct {
	Answer
	IncludeDeleted bool   `json:"include_deleted"`
	LoginUserID    string `json:"login_user_id"`
	Order          string `json:"order_by"`                   // default or updated
	Page           int    `json:"page" form:"page"`           // Query number of pages
	PageSize       int    `json:"page_size" form:"page_size"` // Search page size
}

// TableName answer table name
func (Answer) TableName() string {
	return "answer"
}
