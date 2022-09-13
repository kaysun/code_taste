package service

import (
	"context"
	"fmt"

	"github.com/kaysun/code_taste/003extension/entity"
)

// NewsCommentService 新闻评论应用服务，实现CommentServer 评论服务者接口
type NewsCommentService struct {
	// 嵌套默认评论应用服务，只替换添加评论
	DefaultCommentService
}

// AddComment 添加评论
func (n NewsCommentService) AddComment(ctx context.Context,
	comment entity.Comment) error {
	// 默认添加评论能力
	err := n.DefaultCommentService.AddComment(ctx, comment)
	if err != nil {
		return err
	}

	// 扩展能力调用原子能力（仓储接口）
	repo := n.DefaultCommentService.commentCommandRepo
	err = repo.AddCommentExtension(ctx, comment)
	fmt.Println("NewsCommentService AddComment")

	// 也可以全部调用原子能力，拼装
	//repo.AddCommentInfo(ctx, comment)
	//repo.AddCommentExtension(ctx, comment)
	//repo.AddCommentExtension(ctx, comment)

	return err
}
