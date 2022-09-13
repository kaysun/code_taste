// Package main 处理多种类型，使用简单工厂+策略模式，保证符合开闭原则，修改和添加都单一
package main

import (
	"context"

	"github.com/kaysun/code_taste/002strategy/entity"
	"github.com/kaysun/code_taste/002strategy/processor"
)

func main() {
	comment := entity.CommentInfo{
		CommentID:     1256852,
		Content:       "迪丽热巴真美啊",
		CheckStatus:   entity.CheckStatusApproved,
		CheckType:     entity.CheckTypeReviewBeforeIssuing,
		OperationType: entity.OperationTypeOrigin,
	}
	// 如果想要全局唯一对象，可以采用main函数依赖注入的方式，比全局函数和单例函数优雅得多
	commentProcessor := processor.NewCommentProcessor(comment.OperationType)
	commentProcessor.ProcessComment(context.Background(), comment)
}
