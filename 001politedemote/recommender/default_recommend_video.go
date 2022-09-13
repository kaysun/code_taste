// Package 001politedemote 优雅的mock和兜底
package recommender

import (
	"context"
	"github.com/kaysun/code_taste/001politedemote/entity"
)

/*
DefaultRecommendVideo 默认的推荐视频，实现Recommender 推荐器接口
装饰者模式，作用是兜底，当正常链路失败时，返回兜底数据
*/
type DefaultRecommendVideo struct {
	// Recommender 嵌套推荐器接口对象
	Recommender
}

// GetRecommendList 获取用户的推荐列表
func (d DefaultRecommendVideo) GetRecommendList(ctx context.Context,
	userID uint64) ([]entity.VideoCover, error) {
	recommendCovers, err := d.Recommender.GetRecommendList(ctx, userID)
	if err == nil {
		return recommendCovers, nil
	}
	return []entity.VideoCover{
		{
			Title:     "汽车世界之小小修理工",
			CoverID:   "mzc00200ujklpeq",
			URL:       "mock url",
			ImageURL:  "mock image url",
			VideoType: entity.VideoTypeTV,
			VideoTag:  entity.VideoTagJuvenile,
		},
		{
			Title:     "小猪佩奇全集",
			CoverID:   "bzfkv5se8qaqel2",
			URL:       "mock url",
			ImageURL:  "mock image url",
			VideoType: entity.VideoTypeTV,
			VideoTag:  entity.VideoTagJuvenile,
		},
	}, nil
}
