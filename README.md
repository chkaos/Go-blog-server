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

## 根据操作系统进行golang环境配置并安装依赖 windows
$env:GO111MODULE = "on"
$env:GOPROXY = "https://goproxy.io"
go mod

## npm安装提交规范相关工具
npm i
npm run cz  // commit 规范

## 使用方法

go run cmd/main.go 或者 npm run dev

 todo
 项目结构优化
 提示信息前缀（添加，修改。。。）
 提取重复代码
 获取项目配置方法优化

## 参考项目
[go-gin-example](https://github.com/eddycjy/go-gin-example)
[golang-gin-realworld-example-app](https://github.com/gothinkster/golang-gin-realworld-example-app)
