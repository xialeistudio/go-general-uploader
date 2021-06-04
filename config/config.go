package config

import (
	"github.com/xialeistudio/go-general-uploader/uploader/qiniu"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

type Config struct {
	Qiniu *qiniu.Config
}

// New 读取配置文件
func New(filename string) (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filename = path.Join(home, filename)
	log.Println("读取配置文件:", filename)
	if !isFileExists(filename) {
		err = writeDefaultConfigFile(filename)
		if err != nil {
			return nil, err
		}
	}
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	config := Config{}
	err = yaml.NewDecoder(fp).Decode(&config)
	return &config, nil
}

// 检查文件是否存在
func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// 保存默认配置文件
func writeDefaultConfigFile(filename string) error {
	log.Println("配置文件不存在，写入默认配置:", filename)
	config := &Config{Qiniu: &qiniu.Config{}}
	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fp.Close()
	err = yaml.NewEncoder(fp).Encode(config)
	return err
}
