package processor

import "github.com/kaysun/code_taste/strategy/entity"

// commentProcessorHandler 创建评论操作器的函数对象，一等公民
type commentProcessorHandler func() CommentProcessor

// operationAndProcessorMap 操作类型与处理器映射关系，表格驱动，降低圈复杂度
var operationAndProcessorMap = map[entity.OperationType]commentProcessorHandler{
	entity.OperationTypeOrigin:    newOriginProcessor,
	entity.OperationTypeReply:     newReplyProcessor,
	entity.OperationTypeDelete:    newDeleteProcessor,
	entity.OperationTypeRecommend: newRecommendProcessor,
	// 点赞与取消点赞，都使用评论点赞处理器
	entity.OperationTypeUp:       newUpProcessor,
	entity.OperationTypeCancelUp: newUpProcessor,
	// 作者精选与取消作者精选，都使用作者精选处理器
	entity.OperationTypeAuthorPick:       newAuthorPickProcessor,
	entity.OperationTypeCancelAuthorPick: newAuthorPickProcessor,
	// 标记神评、取消标记神评、运营加赞、取消运营加赞，都使用相运营操作处理器
	entity.OperationTypeMarkGod:       newOperationProcessor,
	entity.OperationTypeCancelMarkGod: newOperationProcessor,
	entity.OperationTypeFakeUp:        newOperationProcessor,
	entity.OperationTypeCancelFakeUp:  newOperationProcessor,
}

// NewCommentProcessor 创建评论处理器接口
func NewCommentProcessor(operationType entity.OperationType) CommentProcessor {
	if handler, ok := operationAndProcessorMap[operationType]; ok {
		return handler()
	}
	return nil
}

// newOriginProcessor 创建发评处理器
func newOriginProcessor() CommentProcessor {
	return CommentOriginProcessor{}
}

// newReplyProcessor 创建回复处理器
func newReplyProcessor() CommentProcessor {
	return CommentReplyProcessor{}
}

// newDeleteProcessor 创建删评处理器
func newDeleteProcessor() CommentProcessor {
	return CommentDeleteProcessor{}
}

// newUpProcessor 创建点赞处理器
func newUpProcessor() CommentProcessor {
	return CommentUpProcessor{}
}

// newAuthorPickProcessor 创建作者精选处理器
func newAuthorPickProcessor() CommentProcessor {
	return CommentAuthorPickProcessor{}
}

// newRecommendProcessor 创建荐评处理器
func newRecommendProcessor() CommentProcessor {
	return CommentRecommendProcessor{}
}

// newOperationProcessor 创建评论运营处理器
func newOperationProcessor() CommentProcessor {
	return CommentOperationProcessor{}
}
