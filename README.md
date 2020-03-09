<p align="center">
  <a href="https://Go-blog-server" target="blank">
    <img src="https://raw.githubusercontent.com/surmon-china/nodepress/master/logo.png" height="90" alt="go-blog-server Logo" />
  </a>
</p>

# Go-blog-server

## Introduction

**RESTful API service for [chkaos.me](https://chkaos.me) blog, powered by [Golang](https://github.com/golang/go), required [Mysql](https://www.mysql.com/cn/) & [Redis](https://redis.io/).** 

**适用于 项目[chkaos.me](https://chkaos.me) 的 RESTful API 服务；基于 golang；
需安装 [Mysql](https://www.mysql.com/cn/) 和 [Redis](https://redis.io/) 方可完整运行；**

**项目结构参考[golang-standards/project-layout](https://github.com/golang-standards/project-layout)； **

## 安装依赖 windows
$env:GO111MODULE = "on"
$env:GOPROXY = "https://goproxy.io"
go mod vendor
## 使用方法

 go run cmd/main.go

 todo
 项目结构优化
 提示信息前缀（添加，修改。。。）
 提取重复代码
 获取项目配置方法优化

 项目参考
 https://segmentfault.com/a/1190000013297705

swag init -g ./internal/routers/router.go