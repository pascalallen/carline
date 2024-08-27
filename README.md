# Carline

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pascalallen/carline)
[![Go Report Card](https://goreportcard.com/badge/github.com/pascalallen/carline)](https://goreportcard.com/report/github.com/pascalallen/carline)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/pascalallen/carline/go.yml)
![GitHub](https://img.shields.io/github/license/pascalallen/carline)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/pascalallen/carline)

![Logo](web/static/logo.svg)

Carline app

## Core Project Tree

```
├── bin/       # Executable CLI commands
├── cmd/       # Go commands
├── internal/  # Supporting packages
└── web/       # Web app components
```

## Features

- Configurable CI/CD pipeline
- Helper scripts
- Frontend linting with ESLint and Prettier

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Development Environment Setup

### Clone Repository

```bash
cd <projects-parent-directory> && git clone https://github.com/pascalallen/carline.git
```

### Copy & Modify `.env` File

```bash
cp .env.example .env
```

### Bring Up Environment

```bash
bin/up
``` 

You will find the site running at [http://localhost:9990/](http://localhost:9990/)

### Install JavaScript Dependencies

```bash
bin/yarn ci
```

### Compile TypeScript with Webpack

```bash
bin/yarn build
```

### Watch For Frontend Changes

```bash
bin/yarn watch
```

### Take Down Environment

```bash
bin/down
```

## Testing

Run tests and create coverage profile:

```bash
bin/exec go test ./... -covermode=count -coverprofile=coverage.out
```

Generate HTML file to view test coverage profile:

```bash
bin/exec go tool cover -html=coverage.out -o coverage.html
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](LICENSE)
