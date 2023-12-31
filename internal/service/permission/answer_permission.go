package permission

import (
	"context"
	"github.com/answerdev/answer/internal/entity"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/schema"
)

// GetAnswerPermission get answer permission
func GetAnswerPermission(ctx context.Context, userID, creatorUserID string,
	status int, canEdit, canDelete, canRecover bool) (
	actions []*schema.PermissionMemberAction) {
	lang := handler.GetLangByCtx(ctx)
	actions = make([]*schema.PermissionMemberAction, 0)
	if len(userID) > 0 {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "report",
			Name:   translator.Tr(lang, reportActionName),
			Type:   "reason",
		})
	}
	if canEdit || userID == creatorUserID {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "edit",
			Name:   translator.Tr(lang, editActionName),
			Type:   "edit",
		})
	}

	if (canDelete || userID == creatorUserID) && status != entity.AnswerStatusDeleted {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "delete",
			Name:   translator.Tr(lang, deleteActionName),
			Type:   "confirm",
		})
	}

	if canRecover && status == entity.AnswerStatusDeleted {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "undelete",
			Name:   translator.Tr(lang, undeleteActionName),
			Type:   "confirm",
		})
	}
	return actions
}
