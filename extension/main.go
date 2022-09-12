// Package main 多个业务类型有共同的业务逻辑，也有各自定制的业务逻辑，采用装饰者模式+策略模式进行优雅扩展
package main

import (
	"context"
	"fmt"

	"github.com/kaysun/code_taste/extension/entity"
	"github.com/kaysun/code_taste/extension/infra"
	"github.com/kaysun/code_taste/extension/service"
)

func main() {
	// 依赖注入，解决全局唯一问题，且没有全局变量和单例的坏味道
	commentCommandRepo := infra.CommentCommand{}
	defaultCommentService := service.NewCommentService(commentCommandRepo)

	// 视频评论服务，有单独的获取评论列表逻辑
	videoCommentService := service.VideoCommentService{
		DefaultCommentService: defaultCommentService,
	}

	// 新闻评论服务，有单独的添加评论逻辑
	newsCommentService := service.NewsCommentService{
		DefaultCommentService: defaultCommentService,
	}
	
	// 体育评论服务，默认的逻辑
	sportsCommentService := defaultCommentService

	// 获取评论列表，局部变量的声明贴近调用位置
	ctx := context.Background()
	var (
		articleID uint64 = 392087762433
		page             = 0
		pageSize         = 10
	)
	videoCommentService.GetComments(ctx, articleID, page, pageSize)
	println()
	newsCommentService.GetComments(ctx, articleID, page, pageSize)
	println()
	sportsCommentService.GetComments(ctx, articleID, page, pageSize)
	println()

	// 添加评论
	comment := entity.Comment{
		CommentID: 778899,
	}
	videoCommentService.AddComment(ctx, comment)
	println()
	newsCommentService.AddComment(ctx, comment)
	println()
	videoCommentService.AddComment(ctx, comment)
	println()
}

func println() {
	fmt.Println("*****************")
	fmt.Println()
}
