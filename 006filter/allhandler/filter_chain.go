package allhandler

// ContentFilterChain 内容过滤器链
type ContentFilterChain struct {
	// Filters 过滤器列表
	Filters []ContentFilter
}

// AddFilter 添加过滤器
func (c *ContentFilterChain) AddFilter(filter ContentFilter) {
	c.Filters = append(c.Filters, filter)
}

// FilterComments 过滤评论
func (c *ContentFilterChain) FilterComments(comments []Comment) []Comment {
	for _, filter := range c.Filters {
		comments = filter.DoFilter(comments)
	}
	return comments
}
