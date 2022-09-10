package processor

import (
	"context"
	"fmt"

	"github.com/kaysun/code_taste/strategy/entity"
)

// CommentAuthorPickProcessor 评论作者精选处理器，实现CommentProcessor 评论处理器接口
type CommentAuthorPickProcessor struct {
}

// ProcessComment 处理作者精选评论
func (CommentAuthorPickProcessor) ProcessComment(ctx context.Context,
	info entity.CommentInfo) error {
	fmt.Println("处理作者精选评论")
	return nil
}
