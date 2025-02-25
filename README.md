# b0

Your AI backend builder.

> [!IMPORTANT]
> Currently in development.

## Database Migration

To generate a migration file, run the following command:

```bash
$ cd ./backend && migrate create database/migrate/postgres migration_name
```

## Running Locally

### Prerequisites

- Go 1.21 or later
- Node.js 18 or later
- PNPM
- PostgreSQL 14
- Redis

### Without Docker

1. Set up the database:

```bash
# Create PostgreSQL database
createdb b0
```

2. Start the backend:

```bash
# Navigate to backend directory
cd backend

# Copy and configure environment variables
cp .env.example .env

# Install dependencies and build
go mod download
go build -o b0 ./cmd

# Run migrations
./b0 migrate up

# Start the server
./b0 server
```

3. Start the frontend:

```bash
# Navigate to frontend directory
cd frontend

# Copy and configure environment variables
cp .env.example .env

# Install dependencies
pnpm install

# Start development server
pnpm dev
```

The application will be available at:

- Frontend: http://localhost:3000
- Backend API: http://localhost:5555

### With Docker

1. Configure environment:

```bash
# Copy environment files
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
```

2. Start all services:

```bash
# Start the development stack
docker compose -f docker-compose.dev.yaml up -d
```

The application will be available at:

- Frontend: http://localhost:80
- Backend API: http://localhost:5555
- Traefik Dashboard: http://localhost:8080

### Available Services

- Frontend: React application
- Backend: Go API server
- PostgreSQL: Database (port 5432)
- Redis: Cache (port 6379)
- Traefik: Reverse proxy and load balancer

## Development

For hot-reload development:

- Frontend changes will automatically reload
- Backend requires rebuilding the Go binary for changes to take effect
