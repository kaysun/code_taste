package service

import (
	"context"

	"github.com/kaysun/code_taste/extension/entity"
)

// CommentServer 评论服务者接口
type CommentServer interface {
	// GetComments 获取评论列表
	GetComments(ctx context.Context, articleID uint64,
		page, pageSize int) ([]entity.Comment, error)

	// AddComment 添加评论
	AddComment(ctx context.Context, comment entity.Comment) error
}
