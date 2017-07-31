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
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `first_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT 'First Name',
  `last_name` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Last Name',
  `email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'E-Mail Address',
  `password` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Password',
  `oauth2_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'oauth_flg flg',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
  `create_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
  `update_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Users Table';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_users`
--

LOCK TABLES `t_users` WRITE;
/*!40000 ALTER TABLE `t_users` DISABLE KEYS */;
INSERT INTO `t_users` VALUES (1,'harry','suzuki','aaaa@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-04-29 21:43:15','2016-04-29 21:43:15'),(2,'taro2','yamada2','aaaa2@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-05-07 11:41:23','2016-05-07 11:41:23'),(3,'taro3','yamada3','aaaa3@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-05-07 19:21:09','2016-05-07 19:21:09'),(4,'taro4','yamada4','aaaa4@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:19:49','2016-06-13 20:19:49'),(5,'taro5','yamada5','aaaa5@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:23','2016-06-13 20:20:23'),(6,'taro6','yamada6','aaaa6@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:24','2016-06-13 20:20:24'),(7,'taro7','yamada7','aaaa7@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:24','2016-06-13 20:20:24'),(8,'taro8','yamada8','aaaa8@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:24','2016-06-13 20:20:24'),(9,'taro9','yamada9','aaaa9@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:24','2016-06-13 20:20:24'),(10,'taro10','yamada10','aaaa10@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:24','2016-06-13 20:20:24'),(11,'taro11','yamada11','aaaa11@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:25','2016-06-13 20:20:25'),(12,'taro12','yamada12','aaaa12@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:25','2016-06-13 20:20:25'),(13,'taro13','yamada13','aaaa13@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:25','2016-06-13 20:20:25'),(14,'taro14','yamada14','aaaa14@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:25','2016-06-13 20:20:25'),(15,'taro15','yamada15','aaaa15@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:26','2016-06-13 20:20:26'),(16,'taro16','yamada16','aaaa16@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:26','2016-06-13 20:20:26'),(17,'taro17','yamada17','aaaa17@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:26','2016-06-13 20:20:26'),(18,'taro18','yamada18','aaaa18@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:26','2016-06-13 20:20:26'),(19,'taro19','yamada19','aaaa19@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:26','2016-06-13 20:20:26'),(20,'taro20','yamada20','aaaa20@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:26','2016-06-13 20:20:26'),(21,'taro21','yamada21','aaaa21@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:26','2016-06-13 20:20:26'),(22,'taro22','yamada22','aaaa22@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:27','2016-06-13 20:20:27'),(23,'taro23','yamada23','aaaa23@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:27','2016-06-13 20:20:27'),(24,'taro24','yamada24','aaaa24@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:27','2016-06-13 20:20:27'),(25,'taro25','yamada25','aaaa25@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:27','2016-06-13 20:20:27'),(26,'taro26','yamada26','aaaa26@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:27','2016-06-13 20:20:27'),(27,'taro27','yamada27','aaaa27@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:28','2016-06-13 20:20:28'),(28,'taro28','yamada28','aaaa28@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:28','2016-06-13 20:20:28'),(29,'taro29','yamada29','aaaa29@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:28','2016-06-13 20:20:28'),(30,'taro30','yamada30','aaaa30@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:28','2016-06-13 20:20:28'),(31,'taro31','yamada31','aaaa31@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:28','2016-06-13 20:20:28'),(32,'taro32','yamada32','aaaa32@test.jp','02aaa55939a894316cfc3427234bf201','0','0','2016-06-13 20:20:28','2016-06-13 20:20:28');
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
