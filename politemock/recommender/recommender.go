package recommender

import (
	"context"
	"github.com/kaysun/code_taste/politemock/entity"
)

// Recommender 推荐器接口
type Recommender interface {
	// GetRecommendList 获取用户的推荐列表
	GetRecommendList(ctx context.Context, userID uint64) ([]entity.VideoCover, error)
}
