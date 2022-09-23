package onehandler

// ContentFilter 内容过滤器
type ContentFilter interface {
	// DoFilter 过滤内容
	DoFilter(comment Comment) bool
}

// SensitiveWordFilter 敏感词过滤器，实现ContentFilter 内容过滤器
type SensitiveWordFilter struct {
}

// DoFilter 过滤内容
func (s SensitiveWordFilter) DoFilter(comment Comment) bool {
	var isValid bool
	// ...
	return isValid
}

// AbuseFilter 谩骂过滤器，实现ContentFilter 内容过滤器
type AbuseFilter struct {
}

// DoFilter 过滤内容
func (s AbuseFilter) DoFilter(comment Comment) bool {
	var isValid bool
	// ...
	return isValid
}

// CheatPraiseFilter 骗赞过滤器，实现ContentFilter 内容过滤器
type CheatPraiseFilter struct {
}

// DoFilter 过滤内容
func (s CheatPraiseFilter) DoFilter(comment Comment) bool {
	var isValid bool
	// ...
	isValid = true
	return isValid
}
