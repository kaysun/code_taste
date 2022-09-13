package infra

import (
	"context"
	"fmt"
	"github.com/kaysun/code_taste/004multilevelstorage/entity"
	"github.com/kaysun/code_taste/004multilevelstorage/repo"
)

// CommentDBAdapter 评论数据库适配器，实现CommentRepo 评论仓储接口
type CommentDBAdapter struct {
	// mysqlProxy MySQL代理
	// mysqlProxy XXXX

	// CommentRepo 评论仓储接口对象。注意，这里是接口对象，对于mysql来说，不知道上层是谁的，只知道上层有这些方法能力。
	repo.CommentRepo
}

// AddComment 添加评论
func (c CommentDBAdapter) AddComment(ctx context.Context,
	comment entity.Comment) error {
	// 写db
	fmt.Println("CommentDBAdapter AddComment")
	// 写上层缓存
	c.CommentRepo.AddComment(ctx, comment)
	return nil
}

// GetComments 获取评论列表
func (c CommentDBAdapter) GetComments(ctx context.Context,
	articleID uint64, page, pageSize int) ([]entity.Comment, error) {
	// 优先读上层缓存
	comments, err := c.CommentRepo.GetComments(ctx, articleID, page, pageSize)
	if len(comments) > 0 && err == nil {
		return comments, err
	}
	// 上层缓存没有数据，读db
	// c.mysqlProxy...
	fmt.Println("CommentDBAdapter GetComments")
	return nil, nil
}
