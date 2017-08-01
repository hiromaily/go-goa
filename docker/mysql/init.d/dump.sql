-- MySQL dump 10.13  Distrib 5.7.12, for osx10.11 (x86_64)
--
-- Host: 127.0.0.1    Database: hiromaily
-- ------------------------------------------------------
-- Server version	5.7.12

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


--
-- DATABASE hiromaily
--
DROP DATABASE IF EXISTS `hiromaily`;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `hiromaily` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `hiromaily`;


--
-- Table structure for table `t_users`
--

DROP TABLE IF EXISTS `t_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `user_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT 'User Name',
  `email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'E-Mail Address',
  `password` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Password',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Users Table';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_users`
--

LOCK TABLES `t_users` WRITE;
/*!40000 ALTER TABLE `t_users` DISABLE KEYS */;
INSERT INTO `t_users` VALUES
  (1,'harry1','aaaa1@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (2,'harry2','aaaa2@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (3,'harry3','aaaa3@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (4,'harry4','aaaa4@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (5,'harry5','aaaa5@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (6,'harry6','aaaa6@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (7,'harry7','aaaa7@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (8,'harry8','aaaa8@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (9,'harry9','aaaa9@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (10,'harry10','aaaa10@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (11,'harry11','aaaa11@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15'),
  (12,'harry12','aaaa12@test.jp','02aaa55939a894316cfc3427234bf201','0','2017-07-31 21:43:15','2017-07-31 21:43:15');
/*!40000 ALTER TABLE `t_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-09-03 16:52:59
