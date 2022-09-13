package infra

import (
	"context"

	"github.com/kaysun/code_taste/003extension/entity"
)

// CommentCommand 评论命令仓储实现，实现CommentCommandRepo 评论命令仓储接口
type CommentCommand struct {
}

// AddCommentInfo 添加评论信息
func (CommentCommand) AddCommentInfo(ctx context.Context,
	comment entity.Comment) error {
	return nil
}

// AddCommentAttr 添加评论属性
func (CommentCommand) AddCommentAttr(ctx context.Context,
	comment entity.Comment) error {
	return nil
}

// AddCommentExtension 添加评论扩展信息
func (CommentCommand) AddCommentExtension(ctx context.Context,
	comment entity.Comment) error {
	return nil
}
