package processor

import (
	"context"
	"fmt"

	"github.com/kaysun/code_taste/strategy/entity"
)

// CommentUpProcessor 评论点赞处理器，实现CommentProcessor 评论处理器接口
type CommentUpProcessor struct {
}

// ProcessComment 处理点赞评论
func (CommentUpProcessor) ProcessComment(ctx context.Context,
	info entity.CommentInfo) error {
	fmt.Println("处理点赞评论")
	return nil
}
