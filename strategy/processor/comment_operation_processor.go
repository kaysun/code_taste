package processor

import (
	"context"

	"github.com/kaysun/code_taste/strategy/entity"
)

// CommentOperationProcessor 评论运营处理器，实现CommentProcessor 评论处理器接口
type CommentOperationProcessor struct {
}

// ProcessComment 处理运营评论
func (CommentOperationProcessor) ProcessComment(ctx context.Context,
	info entity.CommentInfo) error {
	return nil
}
