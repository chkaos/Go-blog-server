<p align="center">
  <a href="https://Go-blog-server" target="blank">
    <img src="https://raw.githubusercontent.com/surmon-china/nodepress/master/logo.png" height="90" alt="go-blog-server Logo" />
  </a>
</p>

# Go-blog-server

## Introduction

**RESTful API service for [astella.me](https://astella.me) blog, powered by [Golang](https://github.com/golang/go), required [Mysql](https://www.mysql.com/cn/) & [Redis](https://redis.io/).** 

**适用于 项目[astella.me](https://astella.me) 的 RESTful API 服务；基于 golang；
需安装 [Mysql](https://www.mysql.com/cn/) 和 [Redis](https://redis.io/) 方可完整运行；**

**项目结构参考[golang-standards/project-layout](https://github.com/golang-standards/project-layout)； **

## 安装依赖 windows
$env:GO111MODULE = "on"
$env:GOPROXY = "https://goproxy.io"
go mod vendor
## 使用方法
数据库迁移
```
bee generate migration tags -driver=mysql -fields="name:string,description:string,icon:string,article_num:int"
bee generate migration category -driver=mysql -fields="name:string,description:string,article_num:int"
bee generate migration mapping-articles-tags -driver=mysql -fields="article_id:int,tag_id:int"
bee generate migration notifications -driver=mysql -fields="type:int,user_id:int,target_id:int,is_viewed:int,content:string,created_at:string"
bee generate migration comments -driver=mysql -fields="state:int,user_id:int,article_id:int,reply_comment_id:int,is_spam:int,content:string,created_at:string,likes_num:int"
bee generate migration articles -driver=mysql -fields="title:string,description:string,keywords:string,category_id:int,content:string,rendered_content:string,created_at:string,published_at:string,updated_at:string,type:int,reproduce_url:string,thumb:string,pvs_num:int,likes_num:int,comments_num:int"
bee generate migration users -driver=mysql -fields="name:string,email:string,password:string,avatar:string,url:string,role:enum,from:enum,created_at:date,update_at:date,mute:int"
bee migrate -driver=mysql -conn="root:@tcp(127.0.0.1:3306)/go-blog-server"

数据迁移失败报错
FATAL ▶ 0005 Column migration.timestamp type mismatch: TYPE: timestamp, DEFAULT: DEFAULT CURRENT_TIMESTAMP()
https://blog.csdn.net/weixin_43671322/article/details/89537182
``` 
 go run ./cmd/main.go

 todo
 项目结构优化
 提示信息前缀（添加，修改。。。）
 提取重复代码