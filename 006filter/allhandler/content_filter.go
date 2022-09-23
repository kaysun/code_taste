package allhandler

// ContentFilter 内容过滤器
type ContentFilter interface {
	// DoFilter 过滤内容
	DoFilter(comments []Comment) []Comment
}

// SensitiveWordFilter 敏感词过滤器，实现ContentFilter 内容过滤器
type SensitiveWordFilter struct {
}

// DoFilter 过滤内容
func (s SensitiveWordFilter) DoFilter(comments []Comment) []Comment {
	// 模拟过滤
	return comments
}

// AbuseFilter 谩骂过滤器，实现ContentFilter 内容过滤器
type AbuseFilter struct {
}

// DoFilter 过滤内容
func (s AbuseFilter) DoFilter(comments []Comment) []Comment {
	// 模拟过滤
	return comments
}

// CheatPraiseFilter 骗赞过滤器，实现ContentFilter 内容过滤器
type CheatPraiseFilter struct {
}

// DoFilter 过滤内容
func (s CheatPraiseFilter) DoFilter(comments []Comment) []Comment {
	var retained []Comment
	for _, comment := range comments {
		if comment.CommentID != 12345678 {
			retained = append(retained, comment)
		}
	}
	return retained
}
