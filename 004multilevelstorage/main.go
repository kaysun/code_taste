// Package main 多级存储，装饰者模式，优雅读写
package main

import (
	"context"
	"fmt"

	"github.com/kaysun/code_taste/004multilevelstorage/entity"
	"github.com/kaysun/code_taste/004multilevelstorage/infra"
)

func main() {
	// 只有main函数才知道存储分级是怎样的，整个业务代码都不需要关注
	commentAdapter := infra.CommentDBAdapter{
		CommentRepo: infra.CommentRedisAdapter{
			CommentRepo: infra.CommentLocalCacheAdapter{},
		},
	}

	// 添加评论，db->redis->local cache
	ctx := context.Background()
	commentAdapter.AddComment(ctx, entity.Comment{
		CommentID: 1234567890,
	})

	fmt.Println()

	// 读评论，local cache->redis->db
	commentAdapter.GetComments(ctx, 48938547583, 0, 20)
}
