package infra

import (
	// import分组：go源码包、本地包、第三方包
	"context"
	"fmt"

	"github.com/kaysun/code_taste/004multilevelstorage/entity"

	"github.com/allegro/bigcache"
)

// CommentLocalCacheAdapter 评论本地缓存适配器，实现CommentRepo 评论仓储接口
type CommentLocalCacheAdapter struct {
	// bigCache 内存缓存组件，该适配器屏蔽了具体使用的内存缓存
	bigCache bigcache.BigCache
}

// AddComment 添加评论
func (CommentLocalCacheAdapter) AddComment(ctx context.Context,
	comment entity.Comment) error {
	fmt.Println("CommentLocalCacheAdapter AddComment")
	return nil
}

// GetComments 获取评论列表
func (c CommentLocalCacheAdapter) GetComments(ctx context.Context,
	articleID uint64, page, pageSize int) ([]entity.Comment, error) {
	// c.BigCache...
	fmt.Println("CommentLocalCacheAdapter GetComments")
	return nil, nil
}
