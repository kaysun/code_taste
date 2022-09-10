package politemock

// VideoCover 视频专辑
type VideoCover struct {
	// CoverID 专辑id
	CoverID string
	// Title 标题
	Title string
	// URL 播放链接
	URL string
	// ImageURL 封面图
	ImageURL string
	// VideoType 视频类型
	VideoType VideoType
	// VideoTag 视频标签
	VideoTag VideoTag
}

// VideoType 视频类型
type VideoType uint8

const (
	// VideoTypeTV 电视剧
	VideoTypeTV VideoType = 1
	// VideoTypeFilm 电影
	VideoTypeFilm VideoType = 2
	// VideoTypeVariety 综艺
	VideoTypeVariety VideoType = 3

	// ...
)

// VideoTag 视频标签
type VideoTag uint8

const (
	// VideoTagSuspense 悬疑
	VideoTagSuspense VideoTag = 1
	// VideoTagWarfare 战争
	VideoTagWarfare VideoTag = 2
	// VideoTagLove 爱情
	VideoTagLove VideoTag = 3
	// VideoTagPalaceOpera 宫廷剧
	VideoTagPalaceOpera VideoTag = 4
	// VideoTagJuvenile 少儿
	VideoTagJuvenile VideoTag = 5

	// ...
)
