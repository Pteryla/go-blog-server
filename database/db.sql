CREATE DATABASE
   IF NOT EXISTS go_server_mysqldb DEFAULT CHARACTER
  SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

USE go_server_mysqldb;
# 创建文章标签表
CREATE TABLE `blog_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(100)        DEFAULT '' COMMENT '标签名称',
    # 公共部分 start
    `created_on`  int(10) unsigned    DEFAULT '0' COMMENT '创建时间',
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned    DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100)        DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned    DEFAULT '0' COMMENT '删除时间',
    `deleted_by`  varchar(100)        DEFAULT '' COMMENT '删除人',
    `is_del`      tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
    # 公共部分 end
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用 1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='标签管理';

# 创建文章栏目表
CREATE TABLE `blog_collection`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`      varchar(100)        DEFAULT '' COMMENT '栏目名称',
    # 公共部分 start
    `created_on`  int(10) unsigned    DEFAULT '0' COMMENT '创建时间',
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned    DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100)        DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned    DEFAULT '0' COMMENT '删除时间',
    `deleted_by`  varchar(100)        DEFAULT '' COMMENT '删除人',
    `is_del`      tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
    # 公共部分 end
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用 1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='栏目管理';

# 创建文章表
CREATE TABLE `blog_article`
(
    `id`              int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title`           varchar(100)        DEFAULT '' COMMENT '文章标题',
    `desc`            varchar(300)        DEFAULT '' COMMENT '文章简述',
    `cover_image_url` varchar(255)        DEFAULT '' COMMENT '封面图片地址',
    `content_type`    tinyint(3) unsigned DEFAULT '1' COMMENT '文章内容类型 0为md文件类型、1为db txt类型',
    `content`         longtext COMMENT '文章内容',
    `content_url`     varchar(255)        DEFAULT '' COMMENT '文章内容地址',
    # 公共部分 start
    `created_on`      int(10) unsigned    DEFAULT '0' COMMENT '创建时间',
    `created_by`      varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on`     int(10) unsigned    DEFAULT '0' COMMENT '修改时间',
    `modified_by`     varchar(100)        DEFAULT '' COMMENT '修改人',
    `deleted_on`      int(10) unsigned    DEFAULT '0' COMMENT '删除时间',
    `deleted_by`      varchar(100)        DEFAULT '' COMMENT '删除人',
    `is_del`          tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
    # 公共部分 end
    `state`           tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='文章管理';

# 文章标签关联表
CREATE TABLE `blog_article_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id`  int(10) unsigned NOT NULL COMMENT '文章id',
    `tag_id`      int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签id',
    # 公共部分 start
    `created_on`  int(10) unsigned          DEFAULT '0' COMMENT '创建时间',
    `created_by`  varchar(100)              DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned          DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100)              DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned          DEFAULT '0' COMMENT '删除时间',
    `deleted_by`  varchar(100)              DEFAULT '' COMMENT '删除人',
    `is_del`      tinyint(3) unsigned       DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
    # 公共部分 end
    `state`       tinyint(3) unsigned       DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='文章标签关联管理';

# 创建文章栏目关联表
CREATE TABLE `blog_article_collection`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id`    int(10) unsigned NOT NULL COMMENT '文章id',
    `collection_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签id',
    # 公共部分 start
    `created_on`    int(10) unsigned          DEFAULT '0' COMMENT '创建时间',
    `created_by`    varchar(100)              DEFAULT '' COMMENT '创建人',
    `modified_on`   int(10) unsigned          DEFAULT '0' COMMENT '修改时间',
    `modified_by`   varchar(100)              DEFAULT '' COMMENT '修改人',
    `deleted_on`    int(10) unsigned          DEFAULT '0' COMMENT '删除时间',
    `deleted_by`    varchar(100)              DEFAULT '' COMMENT '删除人',
    `is_del`        tinyint(3) unsigned       DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
    # 公共部分 end
    `state`         tinyint(3) unsigned       DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='文章栏目关联管理';

# blog用户信息表
CREATE TABLE `blog_user`
(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`             varchar(100)              DEFAULT '' COMMENT '用户昵称',
    `brief`            varchar(200)              DEFAULT '' COMMENT '用户个性签名',
    `avatar_url`       varchar(255)              DEFAULT '' COMMENT '用户头像',
    `uid`              varchar(100)     NOT NULL DEFAULT '' COMMENT '用户uid',
    `account`          varchar(15)      NOT NULL DEFAULT '' COMMENT '用户账号',
    `wx_account`       varchar(200)              DEFAULT '' COMMENT '微信账号',
    `github_address`   varchar(200)              DEFAULT '' COMMENT 'GitHub地址',
    `weibo_address`    varchar(200)              DEFAULT '' COMMENT '微博地址',
    `bilibili_address` varchar(200)              DEFAULT '' COMMENT 'B站地址',
    `password_salt`    varchar(15)               DEFAULT '' COMMENT '密码盐',
    `password`         varchar(18)      NOT NULL DEFAULT '' COMMENT '密码',
    `wx_id`            varchar(200)              DEFAULT '' COMMENT '用户微信id',
    `created_on`       int(10) unsigned          DEFAULT '0' COMMENT '创建时间',
    `modified_on`      int(10) unsigned          DEFAULT '0' COMMENT '修改时间',
    `deleted_on`       int(10) unsigned          DEFAULT '0' COMMENT '注销时间',
    `state`            tinyint(3) unsigned       DEFAULT '1' COMMENT '状态 0为禁用、1为启用、2为注销',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户信息管理';
