package uploader

import "context"

// Uploader 上传接口
type Uploader interface {
	// Upload 上传
	Upload(ctx context.Context, filename string) (string, error)
}
