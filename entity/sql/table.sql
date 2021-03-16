CREATE TABLE `file_meta`
(
    `id`           int(11) NOT NULL AUTO_INCREMENT COMMENT '文件id，主键',
    `sha1`         varchar(48)  NOT NULL COMMENT '文件sha1值',
    `name`         varchar(256) NOT NULL COMMENT '加密后的文件名',
    `size`         bigint(20) DEFAULT 0 COMMENT '文件大小',
    `location`     varchar(256) NOT NULL COMMENT '文件存储位置',
    `created_time` datetime DEFAULT NOW() COMMENT '文件上传时间',
    `updated_time` datetime DEFAULT NOW() COMMENT '文件更新时间',
    `delete_flag`  tinyint(8) NOT NULL DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_sha1` (`sha1`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;