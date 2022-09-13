package repo

import (
	"context"

	"github.com/kaysun/code_taste/004multilevelstorage/entity"
)

// CommentRepo 评论仓储接口
type CommentRepo interface {
	// CommentCommandRepo 评论命令仓储接口
	CommentCommandRepo
	// CommentQueryRepo 评论查询仓储接口
	CommentQueryRepo
}

// CommentCommandRepo 评论命令仓储接口
type CommentCommandRepo interface {
	// AddComment 添加评论
	AddComment(ctx context.Context, comment entity.Comment) error
}

// CommentQueryRepo 评论查询仓储接口
type CommentQueryRepo interface {
	// GetComments 获取评论列表
	GetComments(ctx context.Context, articleID uint64,
		page, pageSize int) ([]entity.Comment, error)
}
