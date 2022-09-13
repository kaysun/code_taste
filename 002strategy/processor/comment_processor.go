package processor

import (
	"context"

	"github.com/kaysun/code_taste/002strategy/entity"
)

// CommentProcessor 评论处理器接口
type CommentProcessor interface {
	// ProcessComment 处理评论
	ProcessComment(ctx context.Context, info entity.CommentInfo) error
}
