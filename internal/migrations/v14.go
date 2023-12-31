package migrations

import (
	"context"
	"time"
	"xorm.io/xorm/schemas"

	"xorm.io/xorm"
)

func updateTheLengthOfRevisionContent(ctx context.Context, x *xorm.Engine) (err error) {
	sess := x.Context(ctx)
	if x.Dialect().URI().DBType == schemas.MYSQL {
		_, err = sess.Exec("ALTER TABLE `revision` CHANGE `content` `content` MEDIUMTEXT NOT NULL;")
	}
	return err
}

type RevisionV14 struct {
	ID           string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt    time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt    time.Time `xorm:"updated TIMESTAMP updated_at"`
	UserID       string    `xorm:"not null default 0 BIGINT(20) user_id"`
	ObjectType   int       `xorm:"not null default 0 INT(11) object_type"`
	ObjectID     string    `xorm:"not null default 0 BIGINT(20) INDEX object_id"`
	Title        string    `xorm:"not null default '' VARCHAR(255) title"`
	Content      string    `xorm:"not null MEDIUMTEXT content"`
	Log          string    `xorm:"VARCHAR(255) log"`
	Status       int       `xorm:"not null default 1 INT(11) status"`
	ReviewUserID int64     `xorm:"not null default 0 BIGINT(20) review_user_id"`
}

func (RevisionV14) TableName() string {
	return "revision"
}
