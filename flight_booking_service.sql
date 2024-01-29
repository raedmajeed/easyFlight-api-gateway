-- MySQL dump 10.13  Distrib 8.1.0, for macos13.3 (arm64)
--
-- Host: localhost    Database: flight_booking_service
-- ------------------------------------------------------
-- Server version	8.1.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `bookings`
--

DROP TABLE IF EXISTS `bookings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bookings` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `pnr` varchar(191) DEFAULT NULL,
  `email` longtext,
  `economy` tinyint(1) DEFAULT NULL,
  `payment_id` longtext,
  `booking_reference` longtext NOT NULL,
  `booking_status` varchar(191) DEFAULT 'PENDING',
  `departure_airport` longtext,
  `arrival_airport` longtext,
  `flight_chart_ids` longblob,
  `total_fare` longtext,
  `cancelled_status` varchar(191) DEFAULT 'false',
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `pnr` (`pnr`),
  KEY `idx_bookings_deleted_at` (`deleted_at`),
  KEY `fk_bookings_user` (`user_id`),
  CONSTRAINT `fk_bookings_user` FOREIGN KEY (`user_id`) REFERENCES `user_data` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bookings`
--

LOCK TABLES `bookings` WRITE;
/*!40000 ALTER TABLE `bookings` DISABLE KEYS */;
INSERT INTO `bookings` VALUES (1,'2023-11-30 17:30:44.249','2023-12-01 19:50:13.561',NULL,'PR0265097175383705443703','raedam786@gmail.com',0,'pay_N76zexAbDsA1eF','0213c922-152a-40a9-98b9-08f4deb279b1','CONFIRMED','DEL','HYD',NULL,'','false',2),(2,'2023-12-01 21:46:00.520','2023-12-01 21:48:08.163',NULL,'PR5237231858811335226825','raedam786@gmail.com',0,'pay_N790EUJgEIWMnB','52752cb1-a49a-4bdd-9b44-51a3c43461c4','CONFIRMED','DEL','HYD',NULL,'','false',2),(3,'2024-01-26 23:14:59.213','2024-01-26 23:14:59.213',NULL,'PRd82268601403354235799f','raedam786@gmail.com',0,'','d8af3e34-da6a-4026-8e95-40ee84f26a2e','PENDING','DEL','BOM',NULL,'','false',2),(4,'2024-01-26 23:32:26.652','2024-01-26 23:32:26.652',NULL,'PR4934783630126817080548','raedam786@gmail.com',0,'','4928e961-a65a-4ae8-aae7-2df9e94047b1','PENDING','DEL','BOM',NULL,'','false',2),(5,'2024-01-26 23:45:58.509','2024-01-26 23:47:44.705',NULL,'PR488418685480233395542b','raedam786@gmail.com',0,'pay_NTKxHeYq2H9NlA','48cb9083-6414-4699-baa7-cc02d3e9907a','CONFIRMED','DEL','BOM',NULL,'','false',2);
/*!40000 ALTER TABLE `bookings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `traveller_booking`
--

DROP TABLE IF EXISTS `traveller_booking`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `traveller_booking` (
  `traveller_id` bigint unsigned NOT NULL,
  `booking_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`traveller_id`,`booking_id`),
  KEY `fk_traveller_booking_booking` (`booking_id`),
  CONSTRAINT `fk_traveller_booking_booking` FOREIGN KEY (`booking_id`) REFERENCES `bookings` (`id`),
  CONSTRAINT `fk_traveller_booking_traveller` FOREIGN KEY (`traveller_id`) REFERENCES `travellers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `traveller_booking`
--

LOCK TABLES `traveller_booking` WRITE;
/*!40000 ALTER TABLE `traveller_booking` DISABLE KEYS */;
INSERT INTO `traveller_booking` VALUES (1,1),(2,1),(3,2),(4,2),(5,3),(6,3),(7,4),(8,4),(9,5),(10,5);
/*!40000 ALTER TABLE `traveller_booking` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `travellers`
--

DROP TABLE IF EXISTS `travellers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `travellers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `age` longtext,
  `gender` longtext,
  `seat_no` longtext,
  `veg_meal_option` tinyint(1) DEFAULT NULL,
  `checked_in` tinyint(1) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `economy` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_travellers_deleted_at` (`deleted_at`),
  KEY `fk_travellers_user` (`user_id`),
  CONSTRAINT `fk_travellers_user` FOREIGN KEY (`user_id`) REFERENCES `user_data` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `travellers`
--

LOCK TABLES `travellers` WRITE;
/*!40000 ALTER TABLE `travellers` DISABLE KEYS */;
INSERT INTO `travellers` VALUES (1,'2023-11-30 17:30:44.258','2023-11-30 17:30:44.258',NULL,'Raed Abdul Majeed','23','male','',0,0,2,0),(2,'2023-11-30 17:30:44.258','2023-11-30 17:30:44.258',NULL,'Shaikh Zidhin','25','male','',0,0,2,0),(3,'2023-12-01 21:46:00.549','2023-12-01 21:46:00.549',NULL,'Raed Abdul Majeed','23','male','',0,0,2,0),(4,'2023-12-01 21:46:00.549','2023-12-01 21:46:00.549',NULL,'Shaikh Zidhin','25','male','',0,0,2,0),(5,'2024-01-26 23:14:59.219','2024-01-26 23:14:59.219',NULL,'Aslam Ak','23','male','',0,0,2,0),(6,'2024-01-26 23:14:59.219','2024-01-26 23:14:59.219',NULL,'Sajesj','25','male','',0,0,2,0),(7,'2024-01-26 23:32:26.653','2024-01-26 23:32:26.653',NULL,'Aslam Ak','23','male','',0,0,2,0),(8,'2024-01-26 23:32:26.653','2024-01-26 23:32:26.653',NULL,'Sajesj','25','male','',0,0,2,0),(9,'2024-01-26 23:45:58.513','2024-01-26 23:45:58.513',NULL,'Aslam Ak','23','male','',0,0,2,0),(10,'2024-01-26 23:45:58.513','2024-01-26 23:45:58.513',NULL,'Sajesj','25','male','',0,0,2,0);
/*!40000 ALTER TABLE `travellers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_data`
--

DROP TABLE IF EXISTS `user_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_data` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` longtext,
  `phone` longtext,
  `password` longtext,
  `name` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_user_data_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_data`
--

LOCK TABLES `user_data` WRITE;
/*!40000 ALTER TABLE `user_data` DISABLE KEYS */;
INSERT INTO `user_data` VALUES (1,'2023-11-21 23:20:22.963','2023-11-21 23:20:22.963',NULL,'raedam786@outlook.com','8547514157','$2a$10$/2VBIRcRoRwimB5wOS0yB.zh6Kmd90RMBhgKpsjTaSEyvFsQS/J2K','Raed Abdul Majeed'),(2,'2023-11-30 15:02:44.016','2023-11-30 15:02:44.016',NULL,'raedam786@gmail.com','7902498141','$2a$10$DEvghUM4zzn3Iud012DgCOP2tX21ZW72u2ZETNf39Pp5bxqoO4HQi','Raed Abdul Majeed');
/*!40000 ALTER TABLE `user_data` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-27 19:24:28
