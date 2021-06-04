package main

import (
	"context"
	"fmt"
	"github.com/xialeistudio/go-general-uploader/config"
	"github.com/xialeistudio/go-general-uploader/uploader"
	"github.com/xialeistudio/go-general-uploader/uploader/qiniu"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	cfg, err := config.New("go-general-uploader.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	if len(os.Args) < 2 {
		log.Println("Usage: go-general-uploader 文件1 文件2 ... 文件N")
		os.Exit(0)
	}

	var client uploader.Uploader
	if cfg.Qiniu != nil {
		client = qiniu.New(cfg.Qiniu)
	}
	if client == nil {
		log.Fatalln("无可用上传配置")
	}

	filenames := os.Args[1:]
	result := make([]string, len(filenames))
	wg := &sync.WaitGroup{}
	wg.Add(len(filenames))

	for i := range filenames {
		i := i
		go func() {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			url, err := client.Upload(ctx, filenames[i])
			if err != nil {
				log.Printf("upload(%v) error(%v)", filenames[i], err)
				return
			}
			result[i] = url
		}()
	}
	wg.Wait()
	fmt.Println(strings.Join(result, "\n"))
}
