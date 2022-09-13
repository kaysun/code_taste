package recommender

import (
	"context"
	"errors"
	"github.com/kaysun/code_taste/001politedemote/entity"
)

// RecommendVideo 推荐视频，实现Recommender 推荐器接口
type RecommendVideo struct {
}

// GetRecommendList 获取用户的推荐列表
func (RecommendVideo) GetRecommendList(ctx context.Context,
	userID uint64) ([]entity.VideoCover, error) {
	return []entity.VideoCover{
		{
			Title:     "超级飞侠第十季",
			CoverID:   "mzc00200hhfx2nk",
			URL:       "mock url",
			ImageURL:  "mock image url",
			VideoType: entity.VideoTypeTV,
			VideoTag:  entity.VideoTagJuvenile,
		},
		{
			Title:     "羊羊游乐园",
			CoverID:   "mzc00200nq9aw5x",
			URL:       "mock url",
			ImageURL:  "mock image url",
			VideoType: entity.VideoTypeTV,
			VideoTag:  entity.VideoTagJuvenile,
		},
	}, errors.New("mock error")
}
