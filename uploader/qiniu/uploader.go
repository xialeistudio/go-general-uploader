package qiniu

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"path"
	"time"
)

// Config 七牛配置
type Config struct {
	AccessKey string
	SecretKey string
	Bucket    string
	BucketURL string
	KeyPrefix string
}

type client struct {
	config *Config
	mac    *qbox.Mac
}

// New 构造方法
func New(config *Config) *client {
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	return &client{
		config: config,
		mac:    mac,
	}
}

// Upload 上传
func (u client) Upload(ctx context.Context, filename string) (string, error) {
	if u.config.AccessKey == "" || u.config.SecretKey == "" || u.config.BucketURL == "" || u.config.Bucket == "" {
		return "", errors.New("七牛配置无效，请检查")
	}
	policy := storage.PutPolicy{
		Scope: u.config.Bucket,
	}
	token := policy.UploadToken(u.mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseHTTPS:      true,
		UseCdnDomains: true,
	}
	objectKey := fmt.Sprintf(
		"%s/%s/%d%s",
		u.config.KeyPrefix,
		time.Now().Format("2006/01/02"),
		time.Now().UnixNano(),
		path.Ext(filename),
	)
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := uploader.PutFile(ctx, &ret, token, objectKey, filename, nil)
	if err != nil {
		return "", err
	}
	return u.config.BucketURL + "/" + ret.Key, nil
}
