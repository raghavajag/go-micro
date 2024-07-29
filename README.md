# Go Microservices Architecture

This project implements a microservices architecture using Go, gRPC, and RabbitMQ.

## Features

### Microservices Architecture
- Designed and implemented a microservices architecture using Go, enabling scalability and loose coupling between services.
- Developed individual services such as auth-service, log-service, and main-service, each responsible for specific functionalities.

### gRPC Communication
- Utilized gRPC for efficient inter-service communication, leveraging its high performance and low latency.
- Implemented gRPC server and client code to enable seamless communication between microservices.

### RabbitMQ Message Broker
- Integrated RabbitMQ as a message broker to facilitate asynchronous communication between services.
- Implemented message publishing and consuming functionality to enable reliable and decoupled communication.

### Authentication and Authorization
- Developed an auth-service to handle user authentication and authorization.
- Implemented user registration, login, and token-based authentication mechanisms.

### Logging and Monitoring
- Created a log-service to centralize logging functionality across all microservices.
- Implemented logging to a database (MongoDB) to persist logs for later analysis and troubleshooting.

### Database Integration
- Connected services to databases such as PostgreSQL and MongoDB for data persistence.
- Implemented database migrations and ORM (Object-Relational Mapping) using Go libraries.

### Containerization with Docker
- Containerized individual microservices using Docker for easy deployment and scalability.
- Created Dockerfiles for each service to define their runtime environments and dependencies.
