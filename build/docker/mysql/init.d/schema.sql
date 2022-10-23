-- Host: 127.0.0.1    Database: go-goa
-- ------------------------------------------------------
-- Server version	5.7.30

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
-- DATABASE go-goa
--
DROP DATABASE IF EXISTS `go-goa`;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `go-goa` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `go-goa`;


--
-- Table structure for table `t_users`
--

DROP TABLE IF EXISTS `t_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT'user id',
  `user_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT'user name',
  `email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT'mail address',
  `password` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT'password',
  `is_deleted` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flag',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`),
  INDEX login (`email`, `password`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Users Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_user` WRITE;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
INSERT INTO `t_user` VALUES
  (1,'Hiroki Yasui','hiroki@goa.com','pITkElZAslj3TdCjrwTmRveORU4LcV7sn/EGiCqm0b0=','0',now(),now()),
  (2,'guest','guest@guest.com','d978eb967fbe04345371478a97f3c903','0',now(),now());
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_company`
--

DROP TABLE IF EXISTS `t_company`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_company` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'company id',
  `name`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'company name',
  `country_id` smallint COLLATE utf8_unicode_ci NOT NULL COMMENT'country id',
  `address`    varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT'address',
  `is_deleted` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flag',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='company table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_company` WRITE;
/*!40000 ALTER TABLE `t_company` DISABLE KEYS */;
INSERT INTO `t_company` VALUES
  (1,'Freelancer',110,'Meguro-ku, Tokyo','0',now(),now()),
  (2,'BookerZzz',155,'Stationsplein 92, 2011 LM Haarlem','0',now(),now()),
  (3,'2Gears',155,'Molslaan 111 NL-2611 RK Delft','0',now(),now()),
  (4,'Hugo.events',155,'Keienbergweg 97 1101 GG Amsterdam','0',now(),now()),
  (5,'J-Stream Inc.',110,'Shibuya-ku, Tokyo','0',now(),now()),
  (6,'PROTERAS Co., Ltd.',110,'Minato-ku, Tokyo','0',now(),now());
/*!40000 ALTER TABLE `t_company` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_tech`
--

DROP TABLE IF EXISTS `t_tech`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_tech` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'tech id',
  `name`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'tech name',
  `is_deleted` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flag',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='tech table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_tech` WRITE;
/*!40000 ALTER TABLE `t_tech` DISABLE KEYS */;
INSERT INTO `t_tech` VALUES
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
  (57,'Oracle','0',now(),now()),
  (60,'Nginx','0',now(),now()),
  (61,'Apache','0',now(),now()),
  (62,'IIS','0',now(),now()),
  (70,'Docker','0',now(),now()),
  (71,'Vagrant','0',now(),now()),
  (72,'Travis-CI','0',now(),now()),
  (73,'Jenkins','0',now(),now()),
  (74,'Fabric','0',now(),now()),
  (80,'HEROKU','0',now(),now()),
  (82,'AWS','0',now(),now()),
  (83,'AWS RDS','0',now(),now()),
  (84,'AWS SQS','0',now(),now()),
  (85,'AWS DynamoDB','0',now(),now()),
  (86,'AWS ElastiCache','0',now(),now()),
  (87,'AWS lambda','0',now(),now()),
  (88,'AWS ECS','0',now(),now()),
  (90,'GCP','0',now(),now()),
  (91,'Kubernetes','0',now(),now()),
  (100,'DRM','0',now(),now()),
  (101,'Microservice architecture','0',now(),now()),
  (102,'DevOps','0',now(),now()),
  (110,'JIRA','0',now(),now()),
  (120,'Selenium','0',now(),now()),
  (121,'PhantomJS','0',now(),now()),
  (122,'Jasmine','0',now(),now()),
  (131,'gRPC','0',now(),now()),
  (132,'Microservice Architecture','0',now(),now());
/*!40000 ALTER TABLE `t_tech` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_user_like_tech`
--

DROP TABLE IF EXISTS `t_user_like_tech`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user_like_tech` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'id',
  `user_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'user id',
  `tech_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'tech id',
  `is_deleted` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flag',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`),
  INDEX user_id (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='user like tech table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_user_like_tech` WRITE;
/*!40000 ALTER TABLE `t_user_like_tech` DISABLE KEYS */;
INSERT INTO `t_user_like_tech` VALUES
  (null,1,1,'0',now(),now()),
  (null,1,10,'0',now(),now()),
  (null,1,32,'0',now(),now()),
  (null,1,70,'0',now(),now()),
  (null,1,101,'0',now(),now()),
  (null,1,102,'0',now(),now());
/*!40000 ALTER TABLE `t_user_like_tech` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_user_dislike_tech`
--

DROP TABLE IF EXISTS `t_user_dislike_tech`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user_dislike_tech` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'id',
  `user_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'user id',
  `tech_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'tech id',
  `is_deleted` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flag',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`),
  INDEX user_id (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='user dislike tech table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_user_dislike_tech` WRITE;
/*!40000 ALTER TABLE `t_user_dislike_tech` DISABLE KEYS */;
INSERT INTO `t_user_dislike_tech` VALUES
  (null,1,25,'0',now(),now()),
  (null,1,12,'0',now(),now()),
  (null,1,27,'0',now(),now());
/*!40000 ALTER TABLE `t_user_dislike_tech` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_user_work_history`
--

DROP TABLE IF EXISTS `t_user_work_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user_work_history` (
  `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT'id',
  `user_id`     int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'user id',
  `company_branch_id`  int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'company branch id',
  `title`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'title',
  `description` json NOT NULL COMMENT'description',
  `tech_ids`    json NOT NULL COMMENT'tech ids',
  `started_at`  date DEFAULT NULL COMMENT'started date',
  `ended_at`    date DEFAULT NULL COMMENT'ended date',
  `is_deleted`  char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flag',
  `created_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`),
  INDEX user_id (`user_id`),
  INDEX started_at (`started_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='user work history table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_user_work_history` WRITE;
/*!40000 ALTER TABLE `t_user_work_history` DISABLE KEYS */;
INSERT INTO `t_user_work_history` VALUES
  (null,1,10,'Golang Developer',
   '["Developping hotel booking web service on <a>Microservice Architecture</a> communicating by <a>gRPC</a>."]',
   '[1,30,83,84,88,131,132]',
   '2017-12-10','2018-06-30','0',now(),now()),
  (null,1,2,'Developer for this resume site',
   '["Developed resume site for job seeking."]',
   '[2,33,39,50,70,72,90,91,80]',
   '2017-07-16','2017-08-08','0',now(),now()),
  (null,1,6,'Full Stack Engineer',
   '["Developed <a>fin-tech</a> related resources management web application both front-end and back-end.","Setting <a>Docker</a> configuration."]',
   '[7,31,34,70,60,51,110]',
   '2016-12-20','2017-07-15','0',now(),now()),
  (null,1,4,'Full Stack Engineer',
   '["Developed event management web application both front-end and back-end.",["API Server was developed by <a>Golang</a> with GOA framework"],"Setting <a>Docker</a> configuration."]',
   '[1,5,7,31,70,60,50,110]',
   '2016-10-19','2016-11-30','0',now(),now()),
  (null,1,1,'Golang and React(ES6) programmer',
   '["Developed web-server (framework) and reverse proxy and <a>JWT</a> based API and login with <a>OAuth2</a> as authentication plus various libraries using <a>Golang</a>.", "Developed worker program to scrape news information to set into <a>MongoDB</a>.", "eveloped front-end using <a>React</a> plus <a>ES6</a> and setup script for environment.", "Setting <a>Docker</a> configuration."]',
   '[1,32,31,70,50,52,72,80]',
   '2016-07-01','2016-09-30','0',now(),now()),
  (null,1,1,'Golang and PHP (HACK) programmer',
   '["Developed advertisement distribution system on back-end for messaging application called <a>LINE</a> using <a>Golang</a>.", ["<a href=\\"https://line.me/en/\\">LINE</a> is popular messaging application like WhatsApp.","Number of monthly active LINE users in 2016 was more than <a>217 million</a> monthly active users."],"Built CI (<a>Jenkins</a>) Environment"]',
   '[1,6,12,83,84,85,86,73,74,71,60]',
   '2015-10-01','2016-06-30','0',now(),now()),
  (null,1,1,'Python programmer',
   '["Developed education system so-called <a>MOOC</a> using edx open source project.", "Built videos distribution environment on AWS S3 and 3rd party CDN and developed various batch program using <a>AWS boto(Python)</a>."]',
   '[9,11,30,60,87,83,86,73,120,121,122]',
   '2015-04-01','2015-09-30','0',now(),now()),
  (null,1,1,'Front-End Engineer',
   '["Developed iPhone and Android game apps.", "Developed auto-generator for View template in MVC from mock designed HTML and runtime library in Javascript.</a>.", "Strategized developing process for Smartphone application."]',
   '[38,30,7,36,14,15,16,12]',
   '2011-11-01','2013-12-31','0',now(),now()),
  (null,1,1,'Application Engineer',
   '["Developed 3D a terrestrial globe using <a>ActionScript3.0</a> and PaperVision3D.", "Developed various site using Javascript.</a>.", "Planed and developed Augmented Reality application so-called <a>AR</a> that can change clothes for those who stand in front of web camera connected to PC using <a>C++</a>, <a>OpenCV</a>, <a>openFrameworks</a>."]',
   '[19,14,21,22]',
   '2010-09-01','2011-10-31','0',now(),now()),
  (null,1,1,'Application Engineer',
   '["Developed management system for students attendance and leaving on Linux Set Top Box. In order to make real time interactive communications on the web, built HTTP Server using <a>Python</a> and <a>Commet</a> was used for it. And developed <a>Flash</a> applicaton as UI and shell scripts for sending signal when touching devices detected."]',
   '[10,16,23]',
   '2011-02-01','2011-04-30','0',now(),now()),
  (null,1,1,'PHP Programmer',
   '["Developed fashion related game on cell phone on <a>LAMP</a> environment", "Improved performance on <a>MySQL</a>."]',
   '[3,50,54]',
   '2010-07-01','2010-08-31','0',now(),now()),
  (null,1,1,'ActionScript3.0 and PHP Programmer',
   '["Developed real-time chat online game among multiple users."]',
   '[19,3,50,54]',
   '2010-01-15','2010-06-30','0',now(),now()),
  (null,1,1,'ASP.NET Programmer',
   '["Developed EC site for music data protected by <a>Microsoft DRM</a>."]',
   '[24,30,56,62,100]',
   '2009-01-01','2009-11-30','0',now(),now()),
  (null,1,7,'Application Engineer',
   '["Developed various video players using <a>Flash</a>, <a>Silverlight</a>, <a>Javascript</a> and <a>BittorrentDNA</a> for P2P on the web and Nintendo Wii Platform.", "Managed and developed podcast portal site including optimizing database performance.", "Developed <a>Microsoft DRM</a> management system packaging video and issuing license using <a>ASP.NET(C#)</a>.", "Developed billing API of credit card and electronic money on movie download sites."]',
   '[24,3,17,20,30,55,51,62,61,100]',
   '2007-01-01','2008-12-31','0',now(),now()),
  (null,1,8,'Application Engineer',
   '["Managed entire projects including planning and proposal, negotiation, system requirement definition, estimates of developement, progress management.", "Developed video contributing software using C++, Flash.", "Developed EC site for video using PHP, MySQL.", "Developed billing API of credit card and electronic money on movie download sites."]',
   '[3,17,14,50]',
   '2005-04-01','2006-11-30','0',now(),now()),
  (null,1,9,'Junior Software Engineer',
   '["Developed car auction system using VB, C, Java."]',
   '[26,27,13,57]',
   '2002-09-01','2005-03-31','0',now(),now());
/*!40000 ALTER TABLE `t_user_work_history` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_country`
--

CREATE TABLE `m_country` (
  `id`         SMALLINT NOT NULL AUTO_INCREMENT COMMENT'country id',
  `country_code` varchar(2) NOT NULL default'' COMMENT'country code',
  `name`       varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT'country name',
  `is_deleted` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flag',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='country table';
/*!40101 SET character_set_client = @saved_cs_client */;
  ;

LOCK TABLES `m_country` WRITE;
/*!40000 ALTER TABLE `m_country` DISABLE KEYS */;
INSERT INTO `m_country` VALUES
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
  /*!40000 ALTER TABLE `m_country` ENABLE KEYS */;
UNLOCK TABLES;


/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
