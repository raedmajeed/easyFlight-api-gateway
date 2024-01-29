-- MySQL dump 10.13  Distrib 8.1.0, for macos13.3 (arm64)
--
-- Host: localhost    Database: flight_booking_db
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
-- Table structure for table `airline_cancellations`
--

DROP TABLE IF EXISTS `airline_cancellations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `airline_cancellations` (
  `airline_id` bigint unsigned DEFAULT NULL,
  `fare_class` longtext,
  `cancellation_deadline_before` bigint unsigned DEFAULT NULL,
  `cancellation_percentage` bigint DEFAULT NULL,
  `refundable` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `airline_cancellations`
--

LOCK TABLES `airline_cancellations` WRITE;
/*!40000 ALTER TABLE `airline_cancellations` DISABLE KEYS */;
/*!40000 ALTER TABLE `airline_cancellations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `airlines`
--

DROP TABLE IF EXISTS `airlines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `airlines` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `airline_name` longtext,
  `company_address` longtext,
  `phone_number` longtext,
  `email` longtext,
  `airline_code` longtext,
  `airline_logo_link` longtext,
  `support_document_link` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_airlines_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `airlines`
--

LOCK TABLES `airlines` WRITE;
/*!40000 ALTER TABLE `airlines` DISABLE KEYS */;
/*!40000 ALTER TABLE `airlines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `flight_type_models`
--

DROP TABLE IF EXISTS `flight_type_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `flight_type_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `type` longtext NOT NULL,
  `flight_model` longtext NOT NULL,
  `description` longtext NOT NULL,
  `manufacturer_name` longtext NOT NULL,
  `manufacturer_country` longtext NOT NULL,
  `max_distance` longtext NOT NULL,
  `cruise_speed` longtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_flight_type_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `flight_type_models`
--

LOCK TABLES `flight_type_models` WRITE;
/*!40000 ALTER TABLE `flight_type_models` DISABLE KEYS */;
INSERT INTO `flight_type_models` VALUES (1,'2023-10-26 18:02:02.370','2023-10-26 18:02:02.370',NULL,'','','','','','',''),(2,'2023-10-26 18:02:08.099','2023-10-26 18:02:08.099',NULL,'','','','','','',''),(3,'2023-10-26 18:03:27.915','2023-10-26 18:03:27.915',NULL,'','','','','','',''),(4,'2023-10-26 18:04:28.607','2023-10-26 18:04:28.607',NULL,'','','','','','',''),(5,'2023-10-26 18:14:37.394','2023-10-26 18:14:37.394',NULL,'','','','','','',''),(6,'2023-10-26 18:21:33.311','2023-10-26 18:21:33.311',NULL,'COMMERCIAL','','','','','',''),(7,'2023-10-26 18:23:34.968','2023-10-26 18:23:34.968',NULL,'COMMERCIAL','','','','','',''),(8,'2023-10-26 18:24:08.379','2023-10-26 18:24:08.379',NULL,'COMMERCIAL','','','','','',''),(9,'2023-10-26 18:25:16.318','2023-10-26 18:25:16.318',NULL,'COMMERCIAL','','','','','',''),(10,'2023-10-26 18:26:02.248','2023-10-26 18:26:02.248',NULL,'COMMERCIAL','','','','','',''),(11,'2023-10-26 18:27:04.142','2023-10-26 18:27:04.142',NULL,'COMMERCIAL','','','','','',''),(12,'2023-10-26 18:27:46.189','2023-10-26 18:27:46.189',NULL,'COMMERCIAL','','','','','',''),(13,'2023-10-26 18:28:15.408','2023-10-26 18:28:15.408',NULL,'COMMERCIAL','','','','','',''),(14,'2023-10-26 18:28:44.249','2023-10-26 18:28:44.249',NULL,'COMMERCIAL','','','','','',''),(15,'2023-10-26 18:31:22.971','2023-10-26 18:31:22.971',NULL,'COMMERCIAL','','','','','',''),(16,'2023-10-26 18:31:40.414','2023-10-26 18:31:40.414',NULL,'COMMERCIAL','','','','','',''),(17,'2023-10-26 18:33:32.315','2023-10-26 18:33:32.315',NULL,'COMMERCIAL','','','','','',''),(18,'2023-10-26 18:34:22.792','2023-10-26 18:34:22.792',NULL,'COMMERCIAL','','','','','',''),(19,'2023-10-26 18:34:55.336','2023-10-26 18:34:55.336',NULL,'COMMERCIAL','','','','','',''),(20,'2023-10-26 18:38:56.807','2023-10-26 18:38:56.807',NULL,'COMMERCIAL','','','','','',''),(21,'2023-10-26 18:39:13.477','2023-10-26 18:39:13.477',NULL,'COMMERCIAL','8888','Your description here','','Manufacturer country','Max distance','Cruise speed'),(22,'2023-10-26 18:52:16.829','2023-10-26 18:52:16.829',NULL,'COMMERCIAL','8888','Your description here','','Manufacturer country','Max distance','Cruise speed'),(23,'2023-10-26 18:54:41.618','2023-10-26 18:54:41.618',NULL,'COMMERCIAL','8888','Your description here','','Manufacturer country','Max distance','Cruise speed');
/*!40000 ALTER TABLE `flight_type_models` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-27 19:15:04
