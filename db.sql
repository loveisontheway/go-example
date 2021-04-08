CREATE DATABASE IF NOT EXISTS `test` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `test`;

SET FOREIGN_KEY_CHECKS=0;

CREATE TABLE `member` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `money` decimal(10,2) unsigned NOT NULL,
  `birthday` date DEFAULT NULL,
  `created_at` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

INSERT INTO `test`.`member` (`id`, `username`, `money`, `birthday`, `created_at`) VALUES ('1', 'mrtwenty', '130.00', '2019-06-19', '1223');
INSERT INTO `test`.`member` (`id`, `username`, `money`, `birthday`, `created_at`) VALUES ('2', 'zhao', '80.00', NULL, '2323');