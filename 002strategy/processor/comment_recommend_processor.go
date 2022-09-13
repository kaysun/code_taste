package processor

import (
	"context"
	"fmt"

	"github.com/kaysun/code_taste/002strategy/entity"
)

// CommentRecommendProcessor 荐评处理器，实现CommentProcessor 评论处理器接口
type CommentRecommendProcessor struct {
}

// ProcessComment 处理用户推荐评论
func (CommentRecommendProcessor) ProcessComment(ctx context.Context,
	info entity.CommentInfo) error {
	fmt.Println("处理用户推荐评论")
	return nil
}
