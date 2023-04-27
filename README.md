# Overview

This repository demonstrates usage of go-swagger3.

## Getting Started

A docker-compose stack is provided that contains the following services:

- _api_: Go rest API
- **swagger**: swagger-ui instance serving openapi 3.0 specification of the API.

To start the docker-compose services, issue the following commands:

```bash
make build-api
make swagger
make up
```

Open a browser and view the openapi swagger docs at _http://localhost:3000_.

The docker-compose services can be stopped by issuing the following command:

```bash
make down
```

## Known Issues

The following are issues currently encountered:

- How to represent a 204 response?
