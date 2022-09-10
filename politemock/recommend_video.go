package politemock

import (
	"context"
	"errors"
)

// RecommendVideo 推荐视频，实现Recommender 推荐器接口
type RecommendVideo struct {
}

// GetRecommendList 获取用户的推荐列表
func (RecommendVideo) GetRecommendList(ctx context.Context,
	userID uint64) ([]VideoCover, error) {
	return []VideoCover{
		{
			Title:     "超级飞侠第十季",
			CoverID:   "mzc00200hhfx2nk",
			URL:       "mock url",
			ImageURL:  "mock image url",
			VideoType: VideoTypeTV,
			VideoTag:  VideoTagJuvenile,
		},
		{
			Title:     "羊羊游乐园",
			CoverID:   "mzc00200nq9aw5x",
			URL:       "mock url",
			ImageURL:  "mock image url",
			VideoType: VideoTypeTV,
			VideoTag:  VideoTagJuvenile,
		},
	}, errors.New("mock error")
}
