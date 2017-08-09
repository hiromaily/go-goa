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
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT'User ID',
  `user_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT'User Name',
  `email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT'E-Mail Address',
  `password` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT'Password',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Users Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_users` WRITE;
/*!40000 ALTER TABLE `t_users` DISABLE KEYS */;
INSERT INTO `t_users` VALUES
  (1,'hiroki','hiroki@goa.com','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (2,'guest','guest@guest.com','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (3,'harry3','aaaa3@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (4,'harry4','aaaa4@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (5,'harry5','aaaa5@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (6,'harry6','aaaa6@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (7,'harry7','aaaa7@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (8,'harry8','aaaa8@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (9,'harry9','aaaa9@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (10,'harry10','aaaa10@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (11,'harry11','aaaa11@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now()),
  (12,'harry12','aaaa12@test.jp','02aaa55939a894316cfc3427234bf201','0',now(),now());
/*!40000 ALTER TABLE `t_users` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_companies`
--

DROP TABLE IF EXISTS `t_companies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_companies` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'Company ID',
  `name`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'Company Name',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Company Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_companies` WRITE;
/*!40000 ALTER TABLE `t_companies` DISABLE KEYS */;
INSERT INTO `t_companies` VALUES
  (1,'Freelancer','0',now(),now()),
  (2,'Hugo.events','0',now(),now()),
  (3,'2Gears','0',now(),now()),
  (4,'J-Stream Inc.','0',now(),now()),
  (5,'PROTERAS Co., Ltd.','0',now(),now()),
  (6,'JIP INFO BRIDGE Inc.','0',now(),now());
/*!40000 ALTER TABLE `t_companies` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_company_detail`
--

CREATE TABLE `t_company_detail` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'Company detail ID',
  `company_id` int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Company ID',
  `hq_flg`     char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'headquarters flg',
  `country_id` smallint COLLATE utf8_unicode_ci NOT NULL COMMENT'Country ID',
  `address`    varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT'Address',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Company Detail Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_company_detail` WRITE;
/*!40000 ALTER TABLE `t_company_detail` DISABLE KEYS */;
INSERT INTO `t_company_detail` VALUES
  (null,1,'1',110,'','0',now(),now()),
  (null,1,'0',155,'','0',now(),now()),
  (null,1,'0',80,'','0',now(),now()),
  (null,2,'1',155,'Keienbergweg 97 1101 GG Amsterdam','0',now(),now()),
  (null,3,'1',128,'17, Boulevard Prince Henri L-1724','0',now(),now()),
  (null,3,'0',155,'Molslaan 111 NL-2611 RK Delft','0',now(),now()),
  (null,4,'1',110,'東京都港区芝二丁目5-6 芝256スクエアビル6階','0',now(),now()),
  (null,5,'1',110,'東京都港区赤坂4-13-13 赤坂ビル4F','0',now(),now()),
  (null,6,'1',110,'東京都江東区東陽2-4-24','0',now(),now());
/*!40000 ALTER TABLE `t_company_detail` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_techs`
--

DROP TABLE IF EXISTS `t_techs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_techs` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'Tech ID',
  `name`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'Tech Name',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Tech Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_techs` WRITE;
/*!40000 ALTER TABLE `t_techs` DISABLE KEYS */;
INSERT INTO `t_techs` VALUES
  (1,'Golang','0',now(),now()),
  (2,'Golang with goa','0',now(),now()),
  (3,'PHP5','0',now(),now()),
  (4,'PHP7','0',now(),now()),
  (5,'PHP7 with Laravel','0',now(),now()),
  (6,'PHP (HACK)','0',now(),now()),
  (7,'Node.js','0',now(),now()),
  (9,'Python2.7','0',now(),now()),
  (10,'Python','0',now(),now()),
  (11,'Django','0',now(),now()),
  (12,'Ruby','0',now(),now()),
  (13,'C','0',now(),now()),
  (14,'C++','0',now(),now()),
  (15,'Cocos-2dx','0',now(),now()),
  (16,'Flash','0',now(),now()),
  (17,'Flash(ActionScript)','0',now(),now()),
  (18,'ActionScript','0',now(),now()),
  (19,'ActionScript3.0','0',now(),now()),
  (20,'SilverLight with C#','0',now(),now()),
  (21,'OpenCV','0',now(),now()),
  (22,'openFrameworks','0',now(),now()),
  (23,'Bash','0',now(),now()),
  (24,'ASP.NET(C#)','0',now(),now()),
  (25,'Microsoft .NET','0',now(),now()),
  (26,'VB','0',now(),now()),
  (27,'Java','0',now(),now()),
  (30,'Javascript','0',now(),now()),
  (31,'ES6','0',now(),now()),
  (32,'React','0',now(),now()),
  (33,'Riot.js','0',now(),now()),
  (34,'Ext JS','0',now(),now()),
  (35,'TypeScript','0',now(),now()),
  (36,'CoffeeScript','0',now(),now()),
  (37,'HTML5','0',now(),now()),
  (38,'HTML5 (Canvas, CSS3)','0',now(),now()),
  (39,'Semantic UI','0',now(),now()),
  (50,'MySQL','0',now(),now()),
  (51,'PostgreSQL','0',now(),now()),
  (52,'MongoDB','0',now(),now()),
  (53,'Redis','0',now(),now()),
  (54,'memcached','0',now(),now()),
  (55,'SQL Server2005','0',now(),now()),
  (56,'SQL Server2008','0',now(),now()),
  (60,'Nginx','0',now(),now()),
  (61,'Apache','0',now(),now()),
  (62,'IIS','0',now(),now()),
  (70,'Docker','0',now(),now()),
  (71,'Vagrant','0',now(),now()),
  (72,'Travis-CI','0',now(),now()),
  (73,'Jenkins','0',now(),now()),
  (74,'Fabric','0',now(),now()),
  (80,'HEROKU','0',now(),now()),
  (81,'GCP','0',now(),now()),
  (82,'AWS','0',now(),now()),
  (83,'AWS RDS','0',now(),now()),
  (84,'AWS SQS','0',now(),now()),
  (85,'AWS DynamoDB','0',now(),now()),
  (86,'AWS ElastiCache','0',now(),now()),
  (87,'AWS lambda','0',now(),now()),
  (100,'DRM','0',now(),now()),
  (101,'Microservice architecture','0',now(),now()),
  (102,'DevOps','0',now(),now()),
  (110,'JIRA','0',now(),now());
/*!40000 ALTER TABLE `t_techs` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_user_like_tech`
--

DROP TABLE IF EXISTS `t_user_like_techs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user_like_techs` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'ID',
  `user_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'User ID',
  `tech_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Tech ID',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='User Like Tech Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_user_like_techs` WRITE;
/*!40000 ALTER TABLE `t_user_like_techs` DISABLE KEYS */;
INSERT INTO `t_user_like_techs` VALUES
  (null,1,1,'0',now(),now()),
  (null,1,10,'0',now(),now()),
  (null,1,32,'0',now(),now()),
  (null,1,70,'0',now(),now()),
  (null,1,101,'0',now(),now()),
  (null,1,102,'0',now(),now());
/*!40000 ALTER TABLE `t_user_like_techs` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_user_dislike_tech`
--

DROP TABLE IF EXISTS `t_user_dislike_tech`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user_dislike_tech` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'ID',
  `user_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'User ID',
  `tech_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Tech ID',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='User Dislike Tech Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_user_dislike_tech` WRITE;
/*!40000 ALTER TABLE `t_user_dislike_tech` DISABLE KEYS */;
INSERT INTO `t_user_dislike_tech` VALUES
  (null,1,25,'0',now(),now()),
  (null,1,12,'0',now(),now()),
  (null,1,26,'0',now(),now());
/*!40000 ALTER TABLE `t_user_dislike_tech` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_user_work_history`
--

DROP TABLE IF EXISTS `t_user_work_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user_work_history` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'ID',
  `user_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'User ID',
  `company_branch_id`  int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Company Branch ID',
  `title`      varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'Title',
  `started_at` date DEFAULT NULL COMMENT'Started Date',
  `ended_at`   date DEFAULT NULL COMMENT'Ended Date',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='User Work History Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_user_work_history` WRITE;
/*!40000 ALTER TABLE `t_user_work_history` DISABLE KEYS */;
INSERT INTO `t_user_work_history` VALUES
  (null,1,2,'Developer for this resume site','2017-07-16','2017-08-08','0',now(),now()),
  (null,1,6,'Full Stack Engineer','2016-12-20','2017-07-15','0',now(),now()),
  (null,1,4,'Full Stack Engineer','2016-10-19','2017-11-30','0',now(),now()),
  (null,1,1,'Golang and React(ES6) programmer','2016-07-01','2017-09-30','0',now(),now()),
  (null,1,1,'Golang and PHP (HACK) programmer','2015-10-01','2016-06-30','0',now(),now()),
  (null,1,1,'Python programmer','2015-04-01','2015-09-30','0',now(),now()),
  (null,1,1,'Front-End Engineer','2011-11-01','2013-12-31','0',now(),now()),
  (null,1,1,'Application Engineer','2010-09-01','2011-10-31','0',now(),now()),
  (null,1,1,'Application Engineer','2011-02-01','2011-04-30','0',now(),now()),
  (null,1,1,'PHP Programmer','2010-07-01','2010-08-31','0',now(),now()),
  (null,1,1,'ActionScript3.0 and PHP Programmer','2010-01-15','2010-06-30','0',now(),now()),
  (null,1,1,'ASP.NET Programmer','2009-01-01','2009-11-30','0',now(),now()),
  (null,1,7,'Application Engineer','2007-01-01','2008-12-31','0',now(),now()),
  (null,1,8,'Application Engineer','2005-04-01','2006-11-30','0',now(),now()),
  (null,1,9,'Junior Software Engineer','2002-09-01','2005-03-31','0',now(),now());
/*!40000 ALTER TABLE `t_user_work_history` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_user_work_description`
--

DROP TABLE IF EXISTS `t_user_work_description`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user_work_description` (
  `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT'ID',
  `work_id`     int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Work ID',
  `description` text    COLLATE utf8_unicode_ci NOT NULL COMMENT'Description',
  `delete_flg`  char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
  `created_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='User Work Description Table';
/*!40101 SET character_set_client = @saved_cs_client */;








--
-- Table structure for table `t_countries`
--

CREATE TABLE `m_countries` (
  `id`         SMALLINT NOT NULL AUTO_INCREMENT COMMENT'Country ID',
  `country_code` varchar(2) NOT NULL default'' COMMENT'Country Code',
  `name`       varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT'Country Name',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Country Table';
/*!40101 SET character_set_client = @saved_cs_client */;
  ;

LOCK TABLES `m_countries` WRITE;
/*!40000 ALTER TABLE `m_countries` DISABLE KEYS */;
INSERT INTO `m_countries` VALUES
  (null,'AF','Afghanistan','0',now(),now()),
  (null,'AL','Albania','0',now(),now()),
  (null,'DZ','Algeria','0',now(),now()),
  (null,'DS','American Samoa','0',now(),now()),
  (null,'AD','Andorra','0',now(),now()),
  (null,'AO','Angola','0',now(),now()),
  (null,'AI','Anguilla','0',now(),now()),
  (null,'AQ','Antarctica','0',now(),now()),
  (null,'AG','Antigua and Barbuda','0',now(),now()),
  (null,'AR','Argentina','0',now(),now()),
  (null,'AM','Armenia','0',now(),now()),
  (null,'AW','Aruba','0',now(),now()),
  (null,'AU','Australia','0',now(),now()),
  (null,'AT','Austria','0',now(),now()),
  (null,'AZ','Azerbaijan','0',now(),now()),
  (null,'BS','Bahamas','0',now(),now()),
  (null,'BH','Bahrain','0',now(),now()),
  (null,'BD','Bangladesh','0',now(),now()),
  (null,'BB','Barbados','0',now(),now()),
  (null,'BY','Belarus','0',now(),now()),
  (null,'BE','Belgium','0',now(),now()),
  (null,'BZ','Belize','0',now(),now()),
  (null,'BJ','Benin','0',now(),now()),
  (null,'BM','Bermuda','0',now(),now()),
  (null,'BT','Bhutan','0',now(),now()),
  (null,'BO','Bolivia','0',now(),now()),
  (null,'BA','Bosnia and Herzegovina','0',now(),now()),
  (null,'BW','Botswana','0',now(),now()),
  (null,'BV','Bouvet Island','0',now(),now()),
  (null,'BR','Brazil','0',now(),now()),
  (null,'IO','British Indian Ocean Territory','0',now(),now()),
  (null,'BN','Brunei Darussalam','0',now(),now()),
  (null,'BG','Bulgaria','0',now(),now()),
  (null,'BF','Burkina Faso','0',now(),now()),
  (null,'BI','Burundi','0',now(),now()),
  (null,'KH','Cambodia','0',now(),now()),
  (null,'CM','Cameroon','0',now(),now()),
  (null,'CA','Canada','0',now(),now()),
  (null,'CV','Cape Verde','0',now(),now()),
  (null,'KY','Cayman Islands','0',now(),now()),
  (null,'CF','Central African Republic','0',now(),now()),
  (null,'TD','Chad','0',now(),now()),
  (null,'CL','Chile','0',now(),now()),
  (null,'CN','China','0',now(),now()),
  (null,'CX','Christmas Island','0',now(),now()),
  (null,'CC','Cocos (Keeling) Islands','0',now(),now()),
  (null,'CO','Colombia','0',now(),now()),
  (null,'KM','Comoros','0',now(),now()),
  (null,'CG','Congo','0',now(),now()),
  (null,'CK','Cook Islands','0',now(),now()),
  (null,'CR','Costa Rica','0',now(),now()),
  (null,'HR','Croatia (Hrvatska)','0',now(),now()),
  (null,'CU','Cuba','0',now(),now()),
  (null,'CY','Cyprus','0',now(),now()),
  (null,'CZ','Czech Republic','0',now(),now()),
  (null,'DK','Denmark','0',now(),now()),
  (null,'DJ','Djibouti','0',now(),now()),
  (null,'DM','Dominica','0',now(),now()),
  (null,'DO','Dominican Republic','0',now(),now()),
  (null,'TP','East Timor','0',now(),now()),
  (null,'EC','Ecuador','0',now(),now()),
  (null,'EG','Egypt','0',now(),now()),
  (null,'SV','El Salvador','0',now(),now()),
  (null,'GQ','Equatorial Guinea','0',now(),now()),
  (null,'ER','Eritrea','0',now(),now()),
  (null,'EE','Estonia','0',now(),now()),
  (null,'ET','Ethiopia','0',now(),now()),
  (null,'FK','Falkland Islands (Malvinas)','0',now(),now()),
  (null,'FO','Faroe Islands','0',now(),now()),
  (null,'FJ','Fiji','0',now(),now()),
  (null,'FI','Finland','0',now(),now()),
  (null,'FR','France','0',now(),now()),
  (null,'FX','France, Metropolitan','0',now(),now()),
  (null,'GF','French Guiana','0',now(),now()),
  (null,'PF','French Polynesia','0',now(),now()),
  (null,'TF','French Southern Territories','0',now(),now()),
  (null,'GA','Gabon','0',now(),now()),
  (null,'GM','Gambia','0',now(),now()),
  (null,'GE','Georgia','0',now(),now()),
  (null,'DE','Germany','0',now(),now()),
  (null,'GH','Ghana','0',now(),now()),
  (null,'GI','Gibraltar','0',now(),now()),
  (null,'GK','Guernsey','0',now(),now()),
  (null,'GR','Greece','0',now(),now()),
  (null,'GL','Greenland','0',now(),now()),
  (null,'GD','Grenada','0',now(),now()),
  (null,'GP','Guadeloupe','0',now(),now()),
  (null,'GU','Guam','0',now(),now()),
  (null,'GT','Guatemala','0',now(),now()),
  (null,'GN','Guinea','0',now(),now()),
  (null,'GW','Guinea-Bissau','0',now(),now()),
  (null,'GY','Guyana','0',now(),now()),
  (null,'HT','Haiti','0',now(),now()),
  (null,'HM','Heard and Mc Donald Islands','0',now(),now()),
  (null,'HN','Honduras','0',now(),now()),
  (null,'HK','Hong Kong','0',now(),now()),
  (null,'HU','Hungary','0',now(),now()),
  (null,'IS','Iceland','0',now(),now()),
  (null,'IN','India','0',now(),now()),
  (null,'IM','Isle of Man','0',now(),now()),
  (null,'ID','Indonesia','0',now(),now()),
  (null,'IR','Iran (Islamic Republic of)','0',now(),now()),
  (null,'IQ','Iraq','0',now(),now()),
  (null,'IE','Ireland','0',now(),now()),
  (null,'IL','Israel','0',now(),now()),
  (null,'IT','Italy','0',now(),now()),
  (null,'CI','Ivory Coast','0',now(),now()),
  (null,'JE','Jersey','0',now(),now()),
  (null,'JM','Jamaica','0',now(),now()),
  (null,'JP','Japan','0',now(),now()),
  (null,'JO','Jordan','0',now(),now()),
  (null,'KZ','Kazakhstan','0',now(),now()),
  (null,'KE','Kenya','0',now(),now()),
  (null,'KI','Kiribati','0',now(),now()),
  (null,'KP','Korea, Democratic People''s Republic of','0',now(),now()),
  (null,'KR','Korea, Republic of','0',now(),now()),
  (null,'XK','Kosovo','0',now(),now()),
  (null,'KW','Kuwait','0',now(),now()),
  (null,'KG','Kyrgyzstan','0',now(),now()),
  (null,'LA','Lao People''s Democratic Republic','0',now(),now()),
  (null,'LV','Latvia','0',now(),now()),
  (null,'LB','Lebanon','0',now(),now()),
  (null,'LS','Lesotho','0',now(),now()),
  (null,'LR','Liberia','0',now(),now()),
  (null,'LY','Libyan Arab Jamahiriya','0',now(),now()),
  (null,'LI','Liechtenstein','0',now(),now()),
  (null,'LT','Lithuania','0',now(),now()),
  (null,'LU','Luxembourg','0',now(),now()),
  (null,'MO','Macau','0',now(),now()),
  (null,'MK','Macedonia','0',now(),now()),
  (null,'MG','Madagascar','0',now(),now()),
  (null,'MW','Malawi','0',now(),now()),
  (null,'MY','Malaysia','0',now(),now()),
  (null,'MV','Maldives','0',now(),now()),
  (null,'ML','Mali','0',now(),now()),
  (null,'MT','Malta','0',now(),now()),
  (null,'MH','Marshall Islands','0',now(),now()),
  (null,'MQ','Martinique','0',now(),now()),
  (null,'MR','Mauritania','0',now(),now()),
  (null,'MU','Mauritius','0',now(),now()),
  (null,'TY','Mayotte','0',now(),now()),
  (null,'MX','Mexico','0',now(),now()),
  (null,'FM','Micronesia, Federated States of','0',now(),now()),
  (null,'MD','Moldova, Republic of','0',now(),now()),
  (null,'MC','Monaco','0',now(),now()),
  (null,'MN','Mongolia','0',now(),now()),
  (null,'ME','Montenegro','0',now(),now()),
  (null,'MS','Montserrat','0',now(),now()),
  (null,'MA','Morocco','0',now(),now()),
  (null,'MZ','Mozambique','0',now(),now()),
  (null,'MM','Myanmar','0',now(),now()),
  (null,'NA','Namibia','0',now(),now()),
  (null,'NR','Nauru','0',now(),now()),
  (null,'NP','Nepal','0',now(),now()),
  (null,'NL','Netherlands','0',now(),now()),
  (null,'AN','Netherlands Antilles','0',now(),now()),
  (null,'NC','New Caledonia','0',now(),now()),
  (null,'NZ','New Zealand','0',now(),now()),
  (null,'NI','Nicaragua','0',now(),now()),
  (null,'NE','Niger','0',now(),now()),
  (null,'NG','Nigeria','0',now(),now()),
  (null,'NU','Niue','0',now(),now()),
  (null,'NF','Norfolk Island','0',now(),now()),
  (null,'MP','Northern Mariana Islands','0',now(),now()),
  (null,'NO','Norway','0',now(),now()),
  (null,'OM','Oman','0',now(),now()),
  (null,'PK','Pakistan','0',now(),now()),
  (null,'PW','Palau','0',now(),now()),
  (null,'PS','Palestine','0',now(),now()),
  (null,'PA','Panama','0',now(),now()),
  (null,'PG','Papua New Guinea','0',now(),now()),
  (null,'PY','Paraguay','0',now(),now()),
  (null,'PE','Peru','0',now(),now()),
  (null,'PH','Philippines','0',now(),now()),
  (null,'PN','Pitcairn','0',now(),now()),
  (null,'PL','Poland','0',now(),now()),
  (null,'PT','Portugal','0',now(),now()),
  (null,'PR','Puerto Rico','0',now(),now()),
  (null,'QA','Qatar','0',now(),now()),
  (null,'RE','Reunion','0',now(),now()),
  (null,'RO','Romania','0',now(),now()),
  (null,'RU','Russian Federation','0',now(),now()),
  (null,'RW','Rwanda','0',now(),now()),
  (null,'KN','Saint Kitts and Nevis','0',now(),now()),
  (null,'LC','Saint Lucia','0',now(),now()),
  (null,'VC','Saint Vincent and the Grenadines','0',now(),now()),
  (null,'WS','Samoa','0',now(),now()),
  (null,'SM','San Marino','0',now(),now()),
  (null,'ST','Sao Tome and Principe','0',now(),now()),
  (null,'SA','Saudi Arabia','0',now(),now()),
  (null,'SN','Senegal','0',now(),now()),
  (null,'RS','Serbia','0',now(),now()),
  (null,'SC','Seychelles','0',now(),now()),
  (null,'SL','Sierra Leone','0',now(),now()),
  (null,'SG','Singapore','0',now(),now()),
  (null,'SK','Slovakia','0',now(),now()),
  (null,'SI','Slovenia','0',now(),now()),
  (null,'SB','Solomon Islands','0',now(),now()),
  (null,'SO','Somalia','0',now(),now()),
  (null,'ZA','South Africa','0',now(),now()),
  (null,'GS','South Georgia South Sandwich Islands','0',now(),now()),
  (null,'ES','Spain','0',now(),now()),
  (null,'LK','Sri Lanka','0',now(),now()),
  (null,'SH','St. Helena','0',now(),now()),
  (null,'PM','St. Pierre and Miquelon','0',now(),now()),
  (null,'SD','Sudan','0',now(),now()),
  (null,'SR','Suriname','0',now(),now()),
  (null,'SJ','Svalbard and Jan Mayen Islands','0',now(),now()),
  (null,'SZ','Swaziland','0',now(),now()),
  (null,'SE','Sweden','0',now(),now()),
  (null,'CH','Switzerland','0',now(),now()),
  (null,'SY','Syrian Arab Republic','0',now(),now()),
  (null,'TW','Taiwan','0',now(),now()),
  (null,'TJ','Tajikistan','0',now(),now()),
  (null,'TZ','Tanzania','0',now(),now()),
  (null,'TH','Thailand','0',now(),now()),
  (null,'TG','Togo','0',now(),now()),
  (null,'TK','Tokelau','0',now(),now()),
  (null,'TO','Tonga','0',now(),now()),
  (null,'TT','Trinidad and Tobago','0',now(),now()),
  (null,'TN','Tunisia','0',now(),now()),
  (null,'TR','Turkey','0',now(),now()),
  (null,'TM','Turkmenistan','0',now(),now()),
  (null,'TC','Turks and Caicos Islands','0',now(),now()),
  (null,'TV','Tuvalu','0',now(),now()),
  (null,'UG','Uganda','0',now(),now()),
  (null,'UA','Ukraine','0',now(),now()),
  (null,'AE','United Arab Emirates','0',now(),now()),
  (null,'GB','United Kingdom','0',now(),now()),
  (null,'US','United States','0',now(),now()),
  (null,'UM','United States minor outlying islands','0',now(),now()),
  (null,'UY','Uruguay','0',now(),now()),
  (null,'UZ','Uzbekistan','0',now(),now()),
  (null,'VU','Vanuatu','0',now(),now()),
  (null,'VA','Vatican City State','0',now(),now()),
  (null,'VE','Venezuela','0',now(),now()),
  (null,'VN','Vietnam','0',now(),now()),
  (null,'VG','Virgin Islands (British)','0',now(),now()),
  (null,'VI','Virgin Islands (U.S.)','0',now(),now()),
  (null,'WF','Wallis and Futuna Islands','0',now(),now()),
  (null,'EH','Western Sahara','0',now(),now()),
  (null,'YE','Yemen','0',now(),now()),
  (null,'ZR','Zaire','0',now(),now()),
  (null,'ZM','Zambia','0',now(),now()),
  (null,'ZW','Zimbabwe','0',now(),now());
  /*!40000 ALTER TABLE `m_countries` ENABLE KEYS */;
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
