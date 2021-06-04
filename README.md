# 文件上传客户端

## 支持上游

+ 七牛

## 快速开始

### 安装命令

#### Go version < 1.16
GO111MODULE=on go get github.com/xialeistudio/go-general-uploader
#### Go 1.16+
go install github.com/xialeistudio/go-general-uploader

### 初次执行(自动生成配置模板)

```bash
go-general-uploader
```

### 添加七牛配置

根据上一步命令输出的配置文件配置即可。模板如下：

```yaml
qiniu:
  accesskey: ""
  secretkey: ""
  bucket: "static"
  bucketurl: "https://static.ddhigh.com" # 外网资源域名
  keyprefix: "blog/" # key前缀
```
