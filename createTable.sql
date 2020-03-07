CREATE TABLE `tbl_file`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
    `file_md5` char(40) NOT NULL DEFAULT '' COMMNET '文件的MD5',
    `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
    `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
    `file_addr` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '文件存储地址',
    `create_at` DATETIME DEFAULT NOW() COMMENT '创建日期',
    `update_at` DATETIME DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP() COMMENT '更新时间'，
    `status` INT(11) NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除)',
    `ext1` INT(11) DEFAULT '0' COMMENT '备用字段1',
    `ext2` text COMMENT '备用字段2',
    PRIMARY KEY('id'),
    UNIQUE KEY `idx_file_hash` (`file_md5`),
    KEY `idx_status`(`status`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;