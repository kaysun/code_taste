package onehandler

// ContentFilterChain 内容过滤器链
type ContentFilterChain struct {
	// Filters 过滤器列表
	Filters []ContentFilter
}

// AddFilter 添加过滤器
func (c *ContentFilterChain) AddFilter(filter ContentFilter) {
	c.Filters = append(c.Filters, filter)
}

// WhetherFiltered 是否已过滤
func (c *ContentFilterChain) WhetherFiltered(comment Comment) bool {
	for _, filter := range c.Filters {
		if filter.DoFilter(comment) {
			return true
		}
	}
	return false
}
