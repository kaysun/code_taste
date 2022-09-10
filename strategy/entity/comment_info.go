package entity

// CommentInfo 评论信息
type CommentInfo struct {
	// CommentID 评论id
	CommentID uint64
	// Content 评论内容
	Content string
	// CheckStatus 评论状态
	CheckStatus CheckStatus
	// CheckType 审核类型
	CheckType CheckType
	// OperationType 操作类型
	OperationType OperationType
}

// CanDisplay 是否可外显
func (c CommentInfo) CanDisplay() bool {
	return c.CanDisplayWhileFirstServedAndThenReviewed() ||
		c.CanDisplayWhileReviewBeforeIssuing()
}

// CanDisplayWhileFirstServedAndThenReviewed 先发后审是否可外显
func (c CommentInfo) CanDisplayWhileFirstServedAndThenReviewed() bool {
	return c.CheckType == CheckTypeFirstServedAndThenReviewed &&
		c.CheckStatus != CheckStatusReviewFailed
}

// CanDisplayWhileReviewBeforeIssuing 先审后发是否可外显
func (c CommentInfo) CanDisplayWhileReviewBeforeIssuing() bool {
	return c.CheckType == CheckTypeReviewBeforeIssuing &&
		c.CheckStatus == CheckStatusApproved
}

// CheckStatus 评论状态
type CheckStatus uint8

const (
	// CheckStatusToBeReviewed 待审核
	CheckStatusToBeReviewed CheckStatus = 1
	// CheckStatusApproved 审核已通过
	CheckStatusApproved CheckStatus = 2
	// CheckStatusReviewFailed 审核未通过
	CheckStatusReviewFailed CheckStatus = 3
)

// CheckType 审核类型
type CheckType uint8

const (
	// CheckTypeFirstServedAndThenReviewed 先发后审
	CheckTypeFirstServedAndThenReviewed CheckType = 1
	// CheckTypeReviewBeforeIssuing 先审后发
	CheckTypeReviewBeforeIssuing CheckType = 2
)

// OperationType 操作类型
type OperationType uint8

const (
	// OperationTypeOrigin 发布原始评论
	OperationTypeOrigin OperationType = 1
	// OperationTypeReply 回复评论
	OperationTypeReply OperationType = 2
	// OperationTypeDelete 删除评论
	OperationTypeDelete OperationType = 3
	// OperationTypeUp 点赞评论
	OperationTypeUp OperationType = 4
	// OperationTypeCancelUp 取消点赞评论
	OperationTypeCancelUp OperationType = 5
	// OperationTypeAuthorPick 作者精选
	OperationTypeAuthorPick OperationType = 6
	// OperationTypeCancelAuthorPick 取消作者精选
	OperationTypeCancelAuthorPick OperationType = 7
	// OperationTypeRecommend 荐评
	OperationTypeRecommend OperationType = 8
	// OperationTypeMarkGod 标记神评
	OperationTypeMarkGod OperationType = 9
	// OperationTypeCancelMarkGod 取消标记神评
	OperationTypeCancelMarkGod OperationType = 10
	// OperationTypeFakeUp 运营加赞
	OperationTypeFakeUp OperationType = 11
	// OperationTypeCancelFakeUp 取消运营加赞
	OperationTypeCancelFakeUp OperationType = 12
)
