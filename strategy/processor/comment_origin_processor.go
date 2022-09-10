package processor

import (
	"context"

	"github.com/kaysun/code_taste/strategy/entity"
)

// CommentOriginProcessor 评论添加处理器，实现CommentProcessor 评论处理器接口
type CommentOriginProcessor struct {
}

// ProcessComment 处理添加评论
func (CommentOriginProcessor) ProcessComment(ctx context.Context,
	info entity.CommentInfo) error {
	return nil
}
