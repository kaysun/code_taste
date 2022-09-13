package infra

import (
	"context"
	"fmt"

	"github.com/kaysun/code_taste/004multilevelstorage/entity"
	"github.com/kaysun/code_taste/004multilevelstorage/repo"
)

// CommentRedisAdapter 评论远程缓存适配器，实现CommentRepo 评论仓储接口
type CommentRedisAdapter struct {
	// redisProxy redis代理
	// redisProxy XXXX

	// CommentRepo 评论仓储接口对象。注意，这里是接口对象，对于redis来说，不知道上层是谁的，只知道上层有这些方法能力。
	repo.CommentRepo
}

// AddComment 添加评论
func (c CommentRedisAdapter) AddComment(ctx context.Context,
	comment entity.Comment) error {
	// 写redis
	fmt.Println("CommentRedisAdapter AddComment")
	// 写上层缓存
	c.CommentRepo.AddComment(ctx, comment)
	return nil
}

// GetComments 获取评论列表
func (c CommentRedisAdapter) GetComments(ctx context.Context,
	articleID uint64, page, pageSize int) ([]entity.Comment, error) {
	// 优先读上层缓存
	comments, err := c.CommentRepo.GetComments(ctx, articleID, page, pageSize)
	if len(comments) > 0 && err == nil {
		return comments, err
	}
	// 上层缓存没有数据，读redis
	// c.redisProxy...
	fmt.Println("CommentRedisAdapter GetComments")
	return nil, nil
}
