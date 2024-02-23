# BackGoffice

![logo](./img/logo.png)

This repository serves as starting point to projects that require user management on an SQL database through a REST API.

## Routes

- 🩺 Healthcheck
- 🔧 Metrics
- 🙋‍♂️ User creation, activation and password reset

## Features

- Stateful token based authentication
- Graceful shutdown (including Goroutines)
- Configuration by flags
- Structured logging
- Validation
- Error responses
- Pagination, filtering and sorting
- JSON marshalling/unmarshalling
- Various middleware:
  - Rate limiting
  - Panic recovery
  - Authenticated user in request context
  - CORS
- Versioning with git commit hash
- Migrations
