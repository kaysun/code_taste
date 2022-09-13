package repo

import (
	"context"

	"github.com/kaysun/code_taste/003extension/entity"
)

// CommentCommandRepo 评论命令仓储接口
type CommentCommandRepo interface {
	// AddCommentInfo 添加评论信息
	AddCommentInfo(ctx context.Context, comment entity.Comment) error
	// AddCommentAttr 添加评论属性
	AddCommentAttr(ctx context.Context, comment entity.Comment) error
	// AddCommentExtension 添加评论扩展信息
	AddCommentExtension(ctx context.Context, comment entity.Comment) error
}
