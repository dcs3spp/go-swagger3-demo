# Overview

This repository demonstrates usage of go-swagger3.

## Prerequisites

Install go-swagger3.

```bash
go install github.com/parvez3019/go-swagger3@latest
```

## Getting Started

A docker-compose stack is provided that contains the following services:

- _api_: Go rest API
- **swagger**: swagger-ui instance serving openapi 3.0 specification of the API.

To start the docker-compose services, issue the following commands:

```bash
make build-api
make build-swagger
make up
```

Open a browser and view the openapi swagger docs at _http://localhost:3000_.

The docker-compose services can be stopped by issuing the following command:

```bash
make down
```

## Known Issues

- If `go.mod` uses `replace` directive then current release does not parse.
I reported this [here](https://github.com/parvez3019/go-swagger3/issues/21)

They have merged a [fix](https://github.com/parvez3020/go-swagger3/pull/23) into
their master branch, however it has not been released yet. So for the time being
if wanted to use this would have to build from source.

- Also raised an [issue](https://github.com/parvez3019/go-swagger3/issues/22)
for how to document a different example for a type for a different endpoint.

