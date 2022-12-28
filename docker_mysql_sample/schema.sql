CREATE TABLE `user`
(
  `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
  `name`     VARCHAR(20) NOT NULL COMMENT 'ユーザー名',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';
