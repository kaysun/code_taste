package main

import (
	"fmt"

	"github.com/kaysun/code_taste/006filter/allhandler"
	"github.com/kaysun/code_taste/006filter/onehandler"
)

func main() {
	oneHandler()
	allHandler()
}

func oneHandler() {
	chain := onehandler.ContentFilterChain{}
	chain.AddFilter(onehandler.SensitiveWordFilter{})
	chain.AddFilter(onehandler.AbuseFilter{})
	chain.AddFilter(onehandler.CheatPraiseFilter{})
	fmt.Println(fmt.Sprintf("onehandler：是否已过滤=%v",
		chain.WhetherFiltered(onehandler.Comment{
			CommentID: 1234567,
			Content:   "点520个赞，哥们就去表白",
		})))
}

func allHandler() {
	chain := allhandler.ContentFilterChain{}
	chain.AddFilter(allhandler.SensitiveWordFilter{})
	chain.AddFilter(allhandler.AbuseFilter{})
	chain.AddFilter(allhandler.CheatPraiseFilter{})

	comments := []allhandler.Comment{
		{
			CommentID: 12345678,
			Content:   "点520个赞，哥们就去表白",
		},
		{
			CommentID: 12345679,
			Content:   "甘宇被救了，太好了！",
		},
	}
	fmt.Println(fmt.Sprintf("allhandler： 过滤前comments=%+v", comments))
	comments = chain.FilterComments(comments)
	fmt.Println(fmt.Sprintf("allhandler： 过滤后comments=%+v", comments))
}
