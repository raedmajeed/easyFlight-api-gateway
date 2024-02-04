-- MySQL dump 10.13  Distrib 8.3.0, for macos13.6 (arm64)
--
-- Host: localhost    Database: flight_booking_service
-- ------------------------------------------------------
-- Server version	8.3.0

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
-- Table structure for table `admin_tables`
--

DROP TABLE IF EXISTS `admin_tables`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_tables` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` longtext,
  `password` longtext,
  `phone` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_admin_tables_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_tables`
--

LOCK TABLES `admin_tables` WRITE;
/*!40000 ALTER TABLE `admin_tables` DISABLE KEYS */;
/*!40000 ALTER TABLE `admin_tables` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `airline_baggages`
--

DROP TABLE IF EXISTS `airline_baggages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `airline_baggages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `airline_id` bigint DEFAULT NULL,
  `fare_class` longtext,
  `cabin_allowed_weight` bigint DEFAULT NULL,
  `cabin_allowed_length` bigint DEFAULT NULL,
  `cabin_allowed_breadth` bigint DEFAULT NULL,
  `cabin_allowed_height` bigint DEFAULT NULL,
  `hand_allowed_weight` bigint DEFAULT NULL,
  `hand_allowed_length` bigint DEFAULT NULL,
  `hand_allowed_breadth` bigint DEFAULT NULL,
  `hand_allowed_height` bigint DEFAULT NULL,
  `fee_extra_per_kg_cabin` bigint DEFAULT NULL,
  `fee_extra_per_kg_hand` bigint DEFAULT NULL,
  `restrictions` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_airline_baggages_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `airline_baggages`
--

LOCK TABLES `airline_baggages` WRITE;
/*!40000 ALTER TABLE `airline_baggages` DISABLE KEYS */;
/*!40000 ALTER TABLE `airline_baggages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `airline_cancellations`
--

DROP TABLE IF EXISTS `airline_cancellations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `airline_cancellations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `airline_id` bigint DEFAULT NULL,
  `fare_class` longtext,
  `cancellation_deadline_before` bigint DEFAULT NULL,
  `cancellation_percentage` bigint DEFAULT NULL,
  `refundable` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_airline_cancellations_deleted_at` (`deleted_at`)
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
-- Table structure for table `airline_seats`
--

DROP TABLE IF EXISTS `airline_seats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `airline_seats` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `airline_id` bigint DEFAULT NULL,
  `economy_seat_number` bigint DEFAULT NULL,
  `business_seat_number` bigint DEFAULT NULL,
  `economy_seats_per_row` bigint DEFAULT NULL,
  `business_seats_per_row` bigint DEFAULT NULL,
  `economy_seat_layout` longblob,
  `business_seat_layout` longblob,
  PRIMARY KEY (`id`),
  KEY `idx_airline_seats_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `airline_seats`
--

LOCK TABLES `airline_seats` WRITE;
/*!40000 ALTER TABLE `airline_seats` DISABLE KEYS */;
/*!40000 ALTER TABLE `airline_seats` ENABLE KEYS */;
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
  `email` varchar(191) NOT NULL,
  `password` longtext,
  `airline_name` longtext,
  `company_address` longtext,
  `phone_number` longtext,
  `airline_code` varchar(191) NOT NULL,
  `airline_logo_link` longtext,
  `support_document_link` longtext,
  `is_account_locked` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `airline_code` (`airline_code`),
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
-- Table structure for table `airports`
--

DROP TABLE IF EXISTS `airports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `airports` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `airport_code` varchar(191) DEFAULT NULL,
  `airport_name` varchar(191) DEFAULT NULL,
  `city` longtext,
  `country` longtext,
  `region` longtext,
  `latitude` double DEFAULT NULL,
  `longitude` double DEFAULT NULL,
  `iatafcs_code` varchar(191) DEFAULT NULL,
  `icao_code` varchar(191) DEFAULT NULL,
  `website` longtext,
  `contact_email` longtext,
  `contact_phone` longtext,
  `black_listed` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `airport_code` (`airport_code`),
  UNIQUE KEY `airport_name` (`airport_name`),
  UNIQUE KEY `latitude` (`latitude`),
  UNIQUE KEY `longitude` (`longitude`),
  UNIQUE KEY `iatafcs_code` (`iatafcs_code`),
  UNIQUE KEY `icao_code` (`icao_code`),
  KEY `idx_airports_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `airports`
--

LOCK TABLES `airports` WRITE;
/*!40000 ALTER TABLE `airports` DISABLE KEYS */;
/*!40000 ALTER TABLE `airports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booked_seats`
--

DROP TABLE IF EXISTS `booked_seats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `booked_seats` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `flight_chart_no` bigint DEFAULT NULL,
  `airline_id` bigint DEFAULT NULL,
  `economy_seat_total` bigint DEFAULT NULL,
  `business_seat_total` bigint DEFAULT NULL,
  `economy_seat_booked` bigint DEFAULT NULL,
  `business_seat_booked` bigint DEFAULT NULL,
  `economy_seat_layout` longblob,
  `business_seat_layout` longblob,
  PRIMARY KEY (`id`),
  KEY `idx_booked_seats_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booked_seats`
--

LOCK TABLES `booked_seats` WRITE;
/*!40000 ALTER TABLE `booked_seats` DISABLE KEYS */;
/*!40000 ALTER TABLE `booked_seats` ENABLE KEYS */;
UNLOCK TABLES;

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
  `direct_flight_chart_ids` longtext,
  `return_flight_chart_ids` longtext,
  `direct_flight_chart_id` longtext,
  `return_flight_chart_id` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `pnr` (`pnr`),
  KEY `idx_bookings_deleted_at` (`deleted_at`),
  KEY `fk_bookings_user` (`user_id`),
  CONSTRAINT `fk_bookings_user` FOREIGN KEY (`user_id`) REFERENCES `user_data` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bookings`
--

LOCK TABLES `bookings` WRITE;
/*!40000 ALTER TABLE `bookings` DISABLE KEYS */;
INSERT INTO `bookings` VALUES (1,'2023-11-30 17:30:44.249','2023-12-01 19:50:13.561',NULL,'PR0265097175383705443703','raedam786@gmail.com',0,'pay_N76zexAbDsA1eF','0213c922-152a-40a9-98b9-08f4deb279b1','CONFIRMED','DEL','HYD',NULL,'','false',2,'﻿',NULL,NULL,NULL),(2,'2023-12-01 21:46:00.520','2023-12-01 21:48:08.163',NULL,'PR5237231858811335226825','raedam786@outlook.com',0,'pay_N790EUJgEIWMnB','52752cb1-a49a-4bdd-9b44-51a3c43461c4','CONFIRMED','DEL','HYD',NULL,'','false',2,NULL,NULL,NULL,NULL),(3,'2024-02-01 23:14:59.213','2024-01-26 23:14:59.213',NULL,'PRd82268601403354235799f','raedam786@gmail.com',0,'','d8af3e34-da6a-4026-8e95-40ee84f26a2e','PENDING','DEL','BOM',NULL,'','false',2,NULL,NULL,NULL,NULL),(4,'2024-01-26 23:32:26.652','2024-01-26 23:32:26.652',NULL,'PR4934783630126817080548','raedam786@gmail.com',0,'','4928e961-a65a-4ae8-aae7-2df9e94047b1','PENDING','DEL','BOM',NULL,'','false',2,NULL,NULL,NULL,NULL),(5,'2024-01-26 23:45:58.509','2024-01-26 23:47:44.705',NULL,'PR488418685480233395542b','raedam786@gmail.com',0,'pay_NTKxHeYq2H9NlA','48cb9083-6414-4699-baa7-cc02d3e9907a','PENDING','DEL','BOM',NULL,'','false',2,NULL,NULL,NULL,NULL),(6,'2024-02-02 23:48:40.156','2024-02-02 23:53:39.250',NULL,'PR1139684520048826718449','raedam786@gmail.com',0,'pay_NW6nNDhsoxSvFi','118927b7-3f9e-4ba7-b2d5-1e150b2afe8e','CONFIRMED','DEL','BOM',NULL,'20327','false',2,NULL,NULL,NULL,NULL),(7,'2024-02-04 10:10:51.277','2024-02-04 10:16:56.305',NULL,'PR088193642895278826168c','raedam786@gmail.com',0,'pay_NWfwteUdG8ddAI','082cdf1d-6a32-40a0-84d0-829f0fa01dfc','CONFIRMED','DEL','BOM',NULL,'20327','false',2,NULL,NULL,NULL,NULL),(8,'2024-02-04 10:24:12.521','2024-02-04 10:31:24.954',NULL,'PRd131686613655011522960','raedam786@gmail.com',0,'pay_NWgCBzL2GI6D6G','d100e65f-672a-4efe-9109-25e05ae13d31','CONFIRMED','DEL','BOM',NULL,'20327','false',2,'{\"flight_chart_ids\":null}','{\"flight_chart_ids\":null}',NULL,NULL),(17,'2024-02-04 10:37:46.538','2024-02-04 10:37:46.538',NULL,'','raedam786@gmail.com',0,'','1d5c3ce5-249e-40e3-9156-b62b50576da3','PENDING','DEL','BOM',NULL,'','false',2,NULL,NULL,NULL,NULL),(22,'2024-02-04 11:07:45.347','2024-02-04 11:07:45.347',NULL,'PR904418487453965366535a','raedam786@gmail.com',0,'','90ba9061-7976-47c8-908c-3a3c48842cd1','PENDING','DEL','BOM',NULL,'','false',2,NULL,NULL,NULL,NULL),(23,'2024-02-04 11:24:33.385','2024-02-04 11:25:28.790',NULL,'PR6c39345250609081692409','raedam786@gmail.com',0,'pay_NWh7JUU9hEhyEo','6c89ec41-6b29-433e-8038-8976ad9623e7','CONFIRMED','DEL','BOM',NULL,'20327','false',2,'{\"flight_chart_ids\":[6,10]}','{\"flight_chart_ids\":[11]}',NULL,NULL),(24,'2024-02-04 12:11:43.172','2024-02-04 12:18:18.099',NULL,'PR0d82043420123328660771','raedam786@gmail.com',0,'pay_NWi15lHI0WaW0m','0d211856-2a8c-4a76-8097-6e967b00418f','CONFIRMED','DEL','BOM',NULL,'20327','false',2,'{\"flight_chart_ids\":[6,10]}','{\"flight_chart_ids\":[11]}',NULL,NULL),(25,'2024-02-04 12:24:12.935','2024-02-04 12:30:13.661',NULL,'PR8962631282720489538070','raedam786@gmail.com',0,'pay_NWiDgftFMXkCmt','89c00421-3f60-418e-84ae-db04da398fdc','CONFIRMED','DEL','BOM',NULL,'20327','false',2,'{\"flight_chart_ids\":[6,10]}','{\"flight_chart_ids\":[11]}',NULL,NULL),(26,'2024-02-04 12:35:11.990','2024-02-04 12:41:38.742',NULL,'PR533152855785236694681d','raedam786@gmail.com',0,'pay_NWiPl6t8WwmpAx','533defdb-ccaf-4e60-9d0c-eccf42327c7d','CONFIRMED','DEL','BOM',NULL,'20327','false',2,'{\"flight_chart_ids\":[6,10]}','{\"flight_chart_ids\":[11]}',NULL,NULL);
/*!40000 ALTER TABLE `bookings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `flight_charts`
--

DROP TABLE IF EXISTS `flight_charts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `flight_charts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `flight_number` longtext NOT NULL,
  `flight_id` bigint unsigned NOT NULL,
  `status` bigint DEFAULT '0',
  `schedule_id` bigint unsigned NOT NULL,
  `economy_fare` double DEFAULT NULL,
  `business_fare` double DEFAULT NULL,
  `departure_airport` longtext,
  `arrival_airport` longtext,
  `departed` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_flight_charts_deleted_at` (`deleted_at`),
  KEY `fk_flight_charts_flight` (`flight_id`),
  KEY `fk_flight_charts_schedule` (`schedule_id`),
  CONSTRAINT `fk_flight_charts_flight` FOREIGN KEY (`flight_id`) REFERENCES `flight_fleets` (`id`),
  CONSTRAINT `fk_flight_charts_schedule` FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `flight_charts`
--

LOCK TABLES `flight_charts` WRITE;
/*!40000 ALTER TABLE `flight_charts` DISABLE KEYS */;
/*!40000 ALTER TABLE `flight_charts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `flight_fleets`
--

DROP TABLE IF EXISTS `flight_fleets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `flight_fleets` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `flight_number` varchar(191) NOT NULL,
  `airline_id` bigint unsigned NOT NULL,
  `seat_id` bigint unsigned NOT NULL,
  `flight_type_id` bigint unsigned NOT NULL,
  `baggage_policy_id` bigint unsigned NOT NULL,
  `cancellation_policy_id` bigint unsigned NOT NULL,
  `maintenance` tinyint(1) DEFAULT '0',
  `is_in_service` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `flight_number` (`flight_number`),
  KEY `idx_flight_fleets_deleted_at` (`deleted_at`),
  KEY `fk_flight_fleets_airline` (`airline_id`),
  KEY `fk_flight_fleets_seat` (`seat_id`),
  KEY `fk_flight_fleets_flight_type` (`flight_type_id`),
  KEY `fk_flight_fleets_baggage` (`baggage_policy_id`),
  KEY `fk_flight_fleets_cancellation` (`cancellation_policy_id`),
  CONSTRAINT `fk_flight_fleets_airline` FOREIGN KEY (`airline_id`) REFERENCES `airlines` (`id`),
  CONSTRAINT `fk_flight_fleets_baggage` FOREIGN KEY (`baggage_policy_id`) REFERENCES `airline_baggages` (`id`),
  CONSTRAINT `fk_flight_fleets_cancellation` FOREIGN KEY (`cancellation_policy_id`) REFERENCES `airline_cancellations` (`id`),
  CONSTRAINT `fk_flight_fleets_flight_type` FOREIGN KEY (`flight_type_id`) REFERENCES `flight_type_models` (`id`),
  CONSTRAINT `fk_flight_fleets_seat` FOREIGN KEY (`seat_id`) REFERENCES `airline_seats` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `flight_fleets`
--

LOCK TABLES `flight_fleets` WRITE;
/*!40000 ALTER TABLE `flight_fleets` DISABLE KEYS */;
/*!40000 ALTER TABLE `flight_fleets` ENABLE KEYS */;
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
  `flight_model` varchar(191) NOT NULL,
  `description` longtext NOT NULL,
  `manufacturer_name` longtext NOT NULL,
  `manufacturer_country` longtext NOT NULL,
  `max_distance` int NOT NULL,
  `cruise_speed` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `flight_model` (`flight_model`),
  KEY `idx_flight_type_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `flight_type_models`
--

LOCK TABLES `flight_type_models` WRITE;
/*!40000 ALTER TABLE `flight_type_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `flight_type_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedules`
--

DROP TABLE IF EXISTS `schedules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schedules` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `departure_time` longtext,
  `arrival_time` longtext,
  `departure_date` longtext,
  `arrival_date` longtext,
  `departure_airport` longtext,
  `arrival_airport` longtext,
  `departure_date_time` datetime(3) DEFAULT NULL,
  `arrival_date_time` datetime(3) DEFAULT NULL,
  `scheduled` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_schedules_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedules`
--

LOCK TABLES `schedules` WRITE;
/*!40000 ALTER TABLE `schedules` DISABLE KEYS */;
/*!40000 ALTER TABLE `schedules` ENABLE KEYS */;
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
INSERT INTO `traveller_booking` VALUES (1,1),(2,1),(3,2),(4,2),(5,3),(6,3),(7,4),(8,4),(9,5),(10,5),(11,6),(12,6),(13,7),(14,7),(15,8),(16,8),(17,17),(18,17),(19,22),(20,22),(21,23),(22,23),(23,24),(24,24),(25,25),(26,25),(27,26),(28,26);
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
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `travellers`
--

LOCK TABLES `travellers` WRITE;
/*!40000 ALTER TABLE `travellers` DISABLE KEYS */;
INSERT INTO `travellers` VALUES (1,'2023-11-30 17:30:44.258','2023-11-30 17:30:44.258',NULL,'Raed Abdul Majeed','23','male','',0,0,2,0),(2,'2023-11-30 17:30:44.258','2023-11-30 17:30:44.258',NULL,'Shaikh Zidhin','25','male','',0,0,2,0),(3,'2023-12-01 21:46:00.549','2023-12-01 21:46:00.549',NULL,'Raed Abdul Majeed','23','male','',0,0,2,0),(4,'2023-12-01 21:46:00.549','2023-12-01 21:46:00.549',NULL,'Shaikh Zidhin','25','male','',0,0,2,0),(5,'2024-01-26 23:14:59.219','2024-01-26 23:14:59.219',NULL,'Aslam Ak','23','male','',0,0,2,0),(6,'2024-01-26 23:14:59.219','2024-01-26 23:14:59.219',NULL,'Sajesj','25','male','',0,0,2,0),(7,'2024-01-26 23:32:26.653','2024-01-26 23:32:26.653',NULL,'Aslam Ak','23','male','',0,0,2,0),(8,'2024-01-26 23:32:26.653','2024-01-26 23:32:26.653',NULL,'Sajesj','25','male','',0,0,2,0),(9,'2024-01-26 23:45:58.513','2024-01-26 23:45:58.513',NULL,'Aslam Ak','23','male','',0,0,2,0),(10,'2024-01-26 23:45:58.513','2024-01-26 23:45:58.513',NULL,'Sajesj','25','male','',0,0,2,0),(11,'2024-02-02 23:48:40.161','2024-02-02 23:48:40.161',NULL,'Aslam Ak','23','male','',0,0,2,0),(12,'2024-02-02 23:48:40.161','2024-02-02 23:48:40.161',NULL,'Sajesj','25','male','',0,0,2,0),(13,'2024-02-04 10:10:51.280','2024-02-04 10:10:51.280',NULL,'Aslam Ak','23','male','',0,0,2,0),(14,'2024-02-04 10:10:51.280','2024-02-04 10:10:51.280',NULL,'Sajesj','25','male','',0,0,2,0),(15,'2024-02-04 10:24:12.524','2024-02-04 10:24:12.524',NULL,'Aslam Ak','23','male','',0,0,2,0),(16,'2024-02-04 10:24:12.524','2024-02-04 10:24:12.524',NULL,'Sajesj','25','male','',0,0,2,0),(17,'2024-02-04 10:37:46.550','2024-02-04 10:37:46.550',NULL,'Aslam Ak','23','male','',0,0,2,0),(18,'2024-02-04 10:37:46.550','2024-02-04 10:37:46.550',NULL,'Sajesj','25','male','',0,0,2,0),(19,'2024-02-04 11:07:45.350','2024-02-04 11:07:45.350',NULL,'Aslam Ak','23','male','',0,0,2,0),(20,'2024-02-04 11:07:45.350','2024-02-04 11:07:45.350',NULL,'Sajesj','25','male','',0,0,2,0),(21,'2024-02-04 11:24:33.388','2024-02-04 11:24:33.388',NULL,'Aslam Ak','23','male','',0,0,2,0),(22,'2024-02-04 11:24:33.388','2024-02-04 11:24:33.388',NULL,'Sajesj','25','male','',0,0,2,0),(23,'2024-02-04 12:11:43.187','2024-02-04 12:11:43.187',NULL,'Aslam Ak','23','male','',0,0,2,0),(24,'2024-02-04 12:11:43.187','2024-02-04 12:11:43.187',NULL,'Sajesj','25','male','',0,0,2,0),(25,'2024-02-04 12:24:12.945','2024-02-04 12:24:12.945',NULL,'Aslam Ak','23','male','',0,0,2,0),(26,'2024-02-04 12:24:12.945','2024-02-04 12:24:12.945',NULL,'Sajesj','25','male','',0,0,2,0),(27,'2024-02-04 12:35:11.994','2024-02-04 12:35:11.994',NULL,'Aslam Ak','23','male','',0,0,2,0),(28,'2024-02-04 12:35:11.994','2024-02-04 12:35:11.994',NULL,'Sajesj','25','male','',0,0,2,0);
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

-- Dump completed on 2024-02-04 17:06:50
