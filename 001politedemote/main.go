// Package main 解决降级｜兜底场景，或者业务前期需要mock数据场景，采用装饰者模式，优雅兜底与mock
package main

import (
	// import分组：go源码包、本地包、第三方包
	"context"
	"fmt"
	"github.com/kaysun/code_taste/001politedemote/recommender"
	"time"
)

func main() {
	// 需要mock数据时
	recommender := recommender.DefaultRecommendVideo{
		Recommender: recommender.RecommendVideo{},
	}
	// 不需要mock数据时
	//recommender := recommender.RecommendVideo{}

	ctx, _ := context.WithTimeout(context.Background(), 200*time.Millisecond)
	covers, err := recommender.GetRecommendList(ctx, 25985276)
	// 超120列折行
	fmt.Println(fmt.Sprintf("recommender.GetRecommendList "+
		"covers=%+v, err=%+v", covers, err))
}
