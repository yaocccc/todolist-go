# TODOLIST-GO

基于golang写的一个todolist后端demo项目

## 阶段目标

> - [x] 基础WEB服务功能
> - [x] JSON解析相关
> - [x] 参数校验
> - [x] 请求中间件
> - [x] 请求追踪
> - [x] ORM增删改查功能
> - [x] ORM结合WEB服务
> - [ ] 追踪ORM日志
> - [ ] 全局错误处理
> - [ ] SWAGGER swaggo
> - [ ] REDIS能力
> - [ ] MONGO能力
> - [ ] KAFKA能力
> - [ ] GRPC能力
> - [ ] 包管理
> - [ ] 部署

## 项目结构

```plaintext
  .
  ├── apis
  │   ├── apis.go
  │   └── user.go
  ├── config
  │   └── config.go
  ├── middlewares
  │   ├── auth.go
  │   ├── processingCount.go
  │   └── traceLogger.go
  ├── models
  │   ├── article.go
  │   ├── models.go
  │   ├── ref.go
  │   └── tag.go
  ├── routers
  │   └── routers.go
  ├── sql
  │   └── tables.sql
  ├── types
  │   └── pagination
  ├── utils
  │   ├── logger.go
  │   └── time.go
  ├── go.mod
  ├── go.sum
  └── main.go
```

## 项目配置

- 修改 config/config.go 中的mysql配置

## 项目初始化

- go get
- 手动执行 sql/tables.sql 对应的sql语句初始化数据库表结构

## 项目启动

go run main.go

## MYSQL表说明

- articles 文章表
- tags 标签表
- article_tag_refs 文章标签关联表(一个文章可以关联多个标签)
