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
> - [x] SWAGGER
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

## swagger文档

```plaintext
  重新生成: swag init
  访问: http://localhost:9527/swagger/index.html#/
```

## MYSQL表说明

- [表结构](./sql/tables.sql)
- articles 文章表
- tags 标签表
- article_tag_refs 文章标签关联表(一个文章可以关联多个标签)

## APIS

```plaintext
  GetArticles: 批量获取文章
  CreateArticles: 批量创建文章
  UpdateArticles: 批量更新文章
  DeleteArticles: 批量删除文章
  GetTags: 批量获取标签
  CreateTags: 批量创建标签
  UpdateTags: 批量更新标签
  DeleteTags: 批量删除标签
```

## 状况

```plaintext
  实现文章和标签相关的API和MYSQL DB操作 并能实现对应业务逻辑
  swagger文档实现情况ok
  无mongo redis kafka逻辑
  无用户体系
  无登录鉴权相关逻辑

  API出入参 参数校验模糊
  db操作不成规范和体系
  较多的重复代码操作
```
