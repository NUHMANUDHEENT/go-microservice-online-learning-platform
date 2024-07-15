# Online Learning Platform microservice(basic)

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
- [gRPC Protobuf Definitions](#grpc-protobuf-definitions)
- [Testing](#testing)
  - [Using grpcurl](#using-grpcurl)
- [Project Structure](#project-structure)

## Services

### User Service

The User Service manages user information, including registration, login, and profile management.

### Course Service

The Course Service handles course information, including course creation, updates, and fetching course details.

### Enrollment Service

The Enrollment Service manages enrollments of users in courses. It ensures users and courses exist before enrolling.


### Prerequisites

- Go 1.22
- Docker
- Protocol Buffers Compiler (`protoc`)
- grpcurl (for testing)


## gRPC Protobuf Definitions

Each service has its own protobuf definitions located in the `proto` directory of each service. To generate Go code from protobuf files, run:

```sh
protoc --go_out=plugins=grpc:. proto/*.proto
