# easyFlight-api-gateway
Easy Fly Api Gateway Service

### MYSQL DOCKER COMMAND
docker run -d --network mynetwork -p 3307:3306 \
-e MYSQL_ROOT_PASSWORD=12345 \
-e MYSQL_DATABASE="flight_booking_booking" \
--name mysql mysql

### KAFKA DOCKER COMMAND
docker run -d --network easyflight-network -p 9092:9092 \
--name kafka-service moeenz/docker-kafka-kraft 

### REDIS DOCKER COMMAND
docker run -d --network easyflight-network -p 6379:6379 \
--name redis-service redis   

### AIRLINE SERVICE DOCKER COMMAND
docker run -d --network easyflight-network -p 6061:6061 \
--name airline-service airline-service    

### API SERVICE DOCKER COMMAND
docker run -d --network easyflight-network -p 8080:8080 \
--name api-service api-service

### NOTIFICATION SERVICE DOCKER COMMAND
docker run -d --network easyflight-network \
--name notification-service notification-service

### BOOKING SERVICE DOCKER COMMAND
docker run -d --network mynetwork -p 9091:9091 \
--name booking-service booking-service:6

### FRONTEND SERVICE DOCKER COMMAND
docker run -d --network easyflight-network -p 3030:3030 \
--name frontend-service frontend-service


# CREATING MYSQL DUMP

### CREATE MYSQL DUMP FIRST
mysqldump -u <username> -p <database_name> > backup.sql

### CREATE VOLUME
docker volume create flight-booking-data

### ADD THIS FIRST TO CREATE VOLUME DIRECTORY
-v mysql-data:/var/lib/mysql [add this to mysql run query]

### MOVE .SQL DUMP INSIDE CONTAINER
docker cp <path_to_dump_file_on_host> <container_id>:<path_inside_container>

### GET TO INTERACTIVE SCRREN INSIDE CONTAINER
docker exec -it mysql-service bin/bash

### CREATE A DATABASE IN MYSQL
create databases

### RUN THIS TO ADD THE .SQL DUMP TO MYSQL DATABASE
mysql -u <username> -p <database_name> < <dump_file_name>

