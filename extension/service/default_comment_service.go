package service

import (
	"context"
	"fmt"

	"github.com/kaysun/code_taste/extension/entity"
	"github.com/kaysun/code_taste/extension/repo"
)

// DefaultCommentService 默认的评论应用服务，实现CommentServer 评论服务者接口
type DefaultCommentService struct {
	// commentCommandRepo 评论命令仓储接口
	commentCommandRepo repo.CommentCommandRepo
}

// NewCommentService 创建评论应用服务
func NewCommentService(
	commentCommand repo.CommentCommandRepo) DefaultCommentService {
	return DefaultCommentService{
		commentCommandRepo: commentCommand,
	}
}

// GetComments 获取评论列表
func (d DefaultCommentService) GetComments(ctx context.Context, articleID uint64,
	page, pageSize int) ([]entity.Comment, error) {
	fmt.Println("DefaultCommentService GetComments")
	return nil, nil
}

// AddComment 添加评论
func (d DefaultCommentService) AddComment(ctx context.Context,
	comment entity.Comment) error {
	fmt.Println("DefaultCommentService AddComment")
	return nil
}
