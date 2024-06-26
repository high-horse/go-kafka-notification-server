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

Open PowerShell and run the following command to download the docker-compose.yml file:```bash
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/bitnami/containers/main/bitnami/kafka/docker-compose.yml" -OutFile "docker-compose.yml"
```