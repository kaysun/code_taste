package main

import (
	// import分组：go源码包、本地包、第三方包
	"context"
	"fmt"
	"time"

	"github.com/kaysun/code_taste/politemock"
)

func main() {
	// 需要mock数据时
	recommender := politemock.DefaultRecommendVideo{
		Recommender: politemock.RecommendVideo{},
	}
	// 不需要mock数据时
	//recommender := politemock.RecommendVideo{}

	ctx, _ := context.WithTimeout(context.Background(), 200*time.Millisecond)
	covers, err := recommender.GetRecommendList(ctx, 25985276)
	// 超120列折行
	fmt.Println(fmt.Sprintf("recommender.GetRecommendList "+
		"covers=%+v, err=%+v", covers, err))
}
