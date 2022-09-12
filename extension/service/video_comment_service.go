package service

import (
	"context"
	"fmt"

	"github.com/kaysun/code_taste/extension/entity"
)

// VideoCommentService 视频评论应用服务，实现CommentServer 评论服务者接口
type VideoCommentService struct {
	// 嵌套默认评论应用服务，只替换获取评论列表
	DefaultCommentService
}

// GetComments 获取评论列表
func (d VideoCommentService) GetComments(ctx context.Context, articleID uint64,
	page, pageSize int) ([]entity.Comment, error) {
	comments, err := d.DefaultCommentService.GetComments(
		ctx, articleID, page, pageSize)
	// 除了基础数据，还需要单独处理，调用方法，或者微服务，再进行数据merge
	fmt.Println("VideoCommentService GetComments")
	return comments, err
}
