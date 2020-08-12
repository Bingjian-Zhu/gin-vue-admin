-- ----------------------------
-- 删除数据库和用户名
-- ----------------------------
drop database if exists `gin-vue`;
use mysql;
delete from user where user='gin' and host='localhost';
flush privileges;

-- ----------------------------
-- 新建数据库和用户名
-- ----------------------------
-- 支持emoji：需要mysql数据库参数： character_set_server=utf8mb4
create database `gin-vue` default character set utf8mb4 collate utf8mb4_unicode_ci;
use `gin-vue`;
create user `gin`@`localhost` identified by 'Gin123';
grant all privileges on `gin-vue`.* to `gin`@`localhost`;
flush privileges;

-- ----------------------------
-- 使用该数据库
-- ----------------------------
use `gin-vue`;
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for go_article
-- ----------------------------
DROP TABLE IF EXISTS `go_article`;
CREATE TABLE `go_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text COMMENT '内容',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `created_on` datetime  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` datetime  DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为草稿、1为已发布、2为删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

-- ----------------------------
-- Table structure for go_user
-- ----------------------------
DROP TABLE IF EXISTS `go_user`;
CREATE TABLE `go_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `avatar` varchar(255) DEFAULT 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG' COMMENT '头像地址',
  `user_type` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '状态 1管理员、2为普通用户',
  `deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0未进行逻辑删除、1为逻辑删除',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  `created_on` datetime  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` datetime  DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for go_tag
-- ----------------------------
DROP TABLE IF EXISTS `go_tag`;
CREATE TABLE `go_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` datetime  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` datetime  DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint(1) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';


-- ----------------------------
-- Table structure for go_role
-- ----------------------------
DROP TABLE IF EXISTS `go_role`;
CREATE TABLE `go_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `user_name` varchar(50) NOT NULL DEFAULT '' COMMENT '账号，冗余设计',
  `value` varchar(50) DEFAULT '' COMMENT 'role值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for go_menus
-- ----------------------------
DROP TABLE IF EXISTS `go_menus`;
CREATE TABLE `go_menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `url` varchar(255) DEFAULT '' COMMENT '菜单地址',
  `parent_id` int(10) unsigned NOT NULL COMMENT '上级id',
  `parent_name` varchar(50) NOT NULL DEFAULT '' COMMENT '上级菜单名称',
  `level` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '菜单等级',
  `deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0未进行逻辑删除、1为逻辑删除',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  `created_on` datetime  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` datetime  DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for go_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `go_role_menu`;
CREATE TABLE `go_role_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL COMMENT '角色ID',
  `menu_id` int(10) unsigned NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- 插入数据
-- ----------------------------
INSERT INTO `go_user`(`id`, `username`, `password`, `avatar`,`user_type`) VALUES (1, 'admin', '111111', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG',1);
INSERT INTO `go_user`(`id`, `username`, `password`, `avatar`,`user_type`) VALUES (2, 'test', '111111', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG',2);

INSERT INTO `go_tag`(`id`, `name`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `state`) VALUES (1, '1', '2019-08-18 18:56:01', 'test', NULL, '', 0, 1);
INSERT INTO `go_tag`(`id`, `name`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `state`) VALUES (2, '2', '2019-08-16 18:56:06', 'test', NULL, '', 0, 1);
INSERT INTO `go_tag`(`id`, `name`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `state`) VALUES (3, '3', '2019-08-18 18:56:09', 'test', NULL, '', 0, 1);

INSERT INTO `go_article`(`id`, `tag_id`, `title`, `desc`, `content`, `cover_image_url`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `state`) VALUES (1, 1, 'test1', 'test-desc', 'test-content', '', '2019-08-19 21:00:39', 'test-created', '2019-08-19 21:00:39', '', 0, 0);
INSERT INTO `go_article`(`id`, `tag_id`, `title`, `desc`, `content`, `cover_image_url`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `state`) VALUES (2, 1, 'test2', 'test-desc', 'test-content', '', '2019-08-19 21:00:48', 'test-created', '2019-08-19 21:00:48', '', 0, 2);
INSERT INTO `go_article`(`id`, `tag_id`, `title`, `desc`, `content`, `cover_image_url`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `state`) VALUES (3, 1, 'test3', 'test-desc', 'test-content', '', '2019-08-19 21:00:49', 'test-created', '2019-08-19 21:00:49', '', 0, 1);

INSERT INTO `go_role`(`id`, `user_id`, `user_name`, `value`) VALUES (1, 1, 'admin', 'admin');
INSERT INTO `go_role`(`id`, `user_id`, `user_name`, `value`) VALUES (2, 1, 'admin', 'test');
INSERT INTO `go_role`(`id`, `user_id`, `user_name`, `value`) VALUES (3, 2, 'test', 'test');
