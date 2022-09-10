package processor

import (
	"context"

	"github.com/kaysun/code_taste/strategy/entity"
)

// CommentReplyProcessor 评论回复处理器，实现CommentProcessor 评论处理器接口
type CommentReplyProcessor struct {
}

// ProcessComment 处理回复评论
func (CommentReplyProcessor) ProcessComment(ctx context.Context,
	info entity.CommentInfo) error {
	return nil
}
