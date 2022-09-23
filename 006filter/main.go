package main

import (
	"fmt"
	"github.com/kaysun/code_taste/006filter/onehandler"
)

func main() {
	chain := onehandler.ContentFilterChain{}
	chain.AddFilter(onehandler.SensitiveWordFilter{})
	chain.AddFilter(onehandler.AbuseFilter{})
	chain.AddFilter(onehandler.CheatPraiseFilter{})
	fmt.Println(fmt.Sprintf("onehandler是否已过滤=%v",
		chain.WhetherFiltered(onehandler.Comment{
			CommentID: 1234567,
			Content:   "点520个赞，哥们就去表白",
		})))
}
