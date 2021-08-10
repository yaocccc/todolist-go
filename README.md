# TODOLIST-GO

基于golang写的一个todolist后端demo项目

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
