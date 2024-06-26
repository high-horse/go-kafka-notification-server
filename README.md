# Kafka-Notify Project

## Overview

This project sets up a real-time notification system using Kafka and Go. It demonstrates how to produce and consume messages in a Kafka-based architecture.

## Prerequisites

- Docker and Docker Compose
- Go programming language installed
- `curl` (for Unix/Linux) or PowerShell (for Windows)

## Setup Instructions

### 1. Download the Docker Compose File

#### For Linux/Mac:

Open your terminal and run the following command to download the `docker-compose.yml` file:

```bash
curl -sSL https://raw.githubusercontent.com/bitnami/containers/main/bitnami/kafka/docker-compose.yml -o docker-compose.yml
```

#### For Windows:

Open PowerShell and run the following command to download the docker-compose.yml file:
```bash
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/bitnami/containers/main/bitnami/kafka/docker-compose.yml" -OutFile "docker-compose.yml"
```

### 2. Modify the Docker Compose File

Open the docker-compose.yml file in a text editor and replace the line:
```bash
KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://:9092
```

with 
```bash
KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
```

### 3. Start Kafka and Zookeeper Services

Navigate to the directory containing the docker-compose.yml file and run the following command to start the services:
```bash
docker-compose up -d
```

### 4. Verify Docker Containers

Check the running Docker containers to ensure Kafka and Zookeeper are up and running:
```bash
docker ps
```
You should see containers for Kafka and Zookeeper running.

### 5. Running the Go Servers

Navigate to the project directory and run the Go servers for the producer and consumer.
Start the Producer:
```bash
go run cmd/producer/producer.go
```

Start the Consumer:
```bash
go run cmd/consumer/consumer.go
```

### 6. Sending Messages

You can send messages to the Kafka broker using curl or PowerShell.
#### Using curl (Linux/Mac):

To send a message indicating that Bruno started following:
```bash
curl -X POST http://localhost:8080/send -d "fromID=2&toID=1&message=Bruno started following you."
```

To send a message indicating that Lena liked a post:
```bash
curl -X POST http://localhost:8080/send -d "fromID=4&toID=1&message=Lena liked your post: 'My weekend getaway!'"
```

#### Using Invoke-RestMethod (Windows PowerShell):

To send a message indicating that Bruno started following:
```bash
Invoke-RestMethod -Uri http://localhost:8080/send -Method Post -Body @{fromID='2'; toID='1'; message='Bruno started following you.'} -ContentType "application/x-www-form-urlencoded"
```

To send a message indicating that Lena liked a post:
```bash
Invoke-RestMethod -Uri http://localhost:8080/send -Method Post -Body @{fromID='4'; toID='1'; message="Lena liked your post: 'My weekend getaway!'"} -ContentType "application/x-www-form-urlencoded"
```

### 7. Retrieving Notifications

To retrieve notifications for a specific user (e.g., User 1), use the following curl command:
```bash
curl http://localhost:8081/notifications/1
```


## Additional Resources

- [Docker Hub - Bitnami Kafka](https://hub.docker.com/r/bitnami/kafka/)
- [Official Apache Kafka](https://kafka.apache.org/quickstart/)
- [Github Sarama Kafka Package](https://github.com/IBM/sarama)
- [Go Gin Web Framework](https://github.com/gin-gonic/gin)

## Troubleshooting

- Ensure Docker and Docker Compose are properly installed and running.
- Verify that your Go environment is set up correctly.
- Check the logs of your Kafka and Go services for any errors.

If you encounter any issues, please consult the documentation for Docker, Kafka, or the Go language, or reach out for community support.

