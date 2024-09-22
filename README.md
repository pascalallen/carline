# Carline

![Logo](web/static/logo.svg)

Carline app

## Core Project Tree

```
├── bin/       # Executable CLI commands
├── docs/      # Additional documentation
├── cmd/       # Go commands
├── internal/  # Supporting packages
├── scripts/   # Application-specific scripts
└── web/       # Web app components
```

## Features

- Configurable CI/CD pipeline
- Helper scripts
- Google Wire DI container
- JWT/HMAC authentication services
- RabbitMQ message broker
- Query bus
- Asynchronous command bus
- Asynchronous event dispatcher
- Middleware
- Frontend linting with ESLint and Prettier
- Database migrations w/ tooling for managing migrations
- Domain models
- API endpoints for authentication and registration
- API endpoints for server-sent events
- Repositories

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
bin/up <prod>
``` 

You will find the site running at [http://localhost:9990/](http://localhost:9990/)

### Migrate database

```bash
bin/migrate -database "postgres://pascalallen:password@postgres:5432/carline?sslmode=disable" -path . up
```

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
bin/down <prod>
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
