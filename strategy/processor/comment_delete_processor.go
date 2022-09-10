package processor

import (
	"context"

	"github.com/kaysun/code_taste/strategy/entity"
)

// CommentDeleteProcessor 评论删除处理器，实现CommentProcessor 评论处理器接口
type CommentDeleteProcessor struct {
}

// ProcessComment 处理删除评论
func (CommentDeleteProcessor) ProcessComment(ctx context.Context,
	info entity.CommentInfo) error {
	return nil
}
