create database todo;

CREATE TABLE `todo`.`articles` (
    id                   INT               NOT NULL AUTO_INCREMENT,
    type                 INT               DEFAULT 1 COMMENT '用户类型 1普通文本 1代办事项',
    title                VARCHAR(1000)     NOT NULL COMMENT '标题',
    content              TEXT              NOT NULL COMMENT '内容',
    status               INT               NOT NULL COMMENT '状态 0未完成 1已完成',
    completed_time       INT               NOT NULL DEFAULT 0,
    deleted_time         INT               NOT NULL DEFAULT 0,
    created_time         INT               NOT NULL DEFAULT 0,
    updated_time         INT               NOT NULL DEFAULT 0,
    is_deleted           INT               NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    INDEX (type, status),
    INDEX (title),
    INDEX (created_time)
) ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '文本表';

CREATE TABLE `todo`.`tags` (
    id                   INT               NOT NULL AUTO_INCREMENT,
    name                 VARCHAR(255)      NOT NULL COMMENT '名称',
    description          VARCHAR(255)      NOT NULL COMMENT '描述',
    PRIMARY KEY (id),
    INDEX (name)
) ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '标签表';

CREATE TABLE `todo`.`article_tag_refs` (
    id                   INT               NOT NULL AUTO_INCREMENT,
    article_id           INT               NOT NULL,
    tag_id               INT               NOT NULL,
    PRIMARY KEY (id),
    INDEX (article_id, tag_id)
) ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '文章标签关联表';
