# Online Learning Platform

This project is an online learning platform that leverages microservices architecture using gRPC for communication between services. The platform includes three main services:

1. User Service
2. Course Service
3. Enrollment Service

Each service is designed to be isolated and independently deployable, communicating with each other via gRPC.

## Table of Contents

- [Services](#services)
  - [User Service](#user-service)
  - [Course Service](#course-service)
  - [Enrollment Service](#enrollment-service)
- [Setup](#setup)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running Services](#running-services)
- [gRPC Protobuf Definitions](#grpc-protobuf-definitions)
- [Testing](#testing)
  - [Using grpcurl](#using-grpcurl)
- [Project Structure](#project-structure)
- [Using Go Kit](#using-go-kit)

## Services

### User Service

The User Service manages user information, including registration, login, and profile management.

### Course Service

The Course Service handles course information, including course creation, updates, and fetching course details.

### Enrollment Service

The Enrollment Service manages enrollments of users in courses. It ensures users and courses exist before enrolling.

## Setup

### Prerequisites

- Go 1.18+
- Docker
- Protocol Buffers Compiler (`protoc`)
- grpcurl (for testing)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/online-learning-platform.git
    cd online-learning-platform
    ```

2. Install dependencies for each service:
    ```sh
    cd user-service
    go mod tidy
    cd ../course-service
    go mod tidy
    cd ../enrollment-service
    go mod tidy
    ```

### Running Services

1. Start the User Service:
    ```sh
    cd user-service
    go run cmd/main.go
    ```

2. Start the Course Service:
    ```sh
    cd course-service
    go run cmd/main.go
    ```

3. Start the Enrollment Service:
    ```sh
    cd enrollment-service
    go run cmd/main.go
    ```

## gRPC Protobuf Definitions

Each service has its own protobuf definitions located in the `proto` directory of each service. To generate Go code from protobuf files, run:

```sh
protoc --go_out=plugins=grpc:. proto/*.proto
