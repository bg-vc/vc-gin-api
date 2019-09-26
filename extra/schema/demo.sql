
CREATE DATABASE `demo` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户',
  `salt` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'salt',
  `pwd_crypt` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码加密',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0:无效 1:有效',
  `is_admin` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'is_admin 0:否 1:是',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='账号表';

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名字',
  `age` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
  `address` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地址',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';