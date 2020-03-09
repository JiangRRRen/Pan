CREATE TABLE `tbl_user`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL DEFAULF '' COMMENT '用户名',
    `user_pwd` varchar(256) NOT NULL DEFAULT '',
    `email` varchar(64) DEFAULT '',
    `phone` varchar(128) DEFAULT '',
    `email_validated` tinyint(1) DEFAULT 0,
    `phone_validated` TINYINT(1) DEFAULT 0,
    `signup_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `last_active` DEFAULT CURRENT_TIMESTAMP,
    `profile` text COMMENT '用户属性'
    `status` int(11) NOT NULL DEFAULT '0' COMMENT '账户状态：启用、禁止、锁定、标记删除',
    PRIMARY KEY(`id`),
    UNIQUE KEY `idx_phone`(`phone`),
    KEY `idx_status`(`status`)
)ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=uft8mb4;