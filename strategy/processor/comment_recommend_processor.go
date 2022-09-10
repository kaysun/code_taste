package processor

import (
	"context"
	"github.com/kaysun/code_taste/strategy/entity"
)

// CommentRecommendProcessor 荐评处理器，实现CommentProcessor 评论处理器接口
type CommentRecommendProcessor struct {
}

// ProcessComment 处理用户推荐评论
func (CommentRecommendProcessor) ProcessComment(ctx context.Context,
	info entity.CommentInfo) error {
	return nil
}
