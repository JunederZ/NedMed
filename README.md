# NedMed: Self-Hosted Media Uploader

NedMed is a simple self-hosted application for uploading and storing media files on your local machine or self-hosted server, with PostgreSQL integration for metadata storage. It uses a Go backend for handling file uploads and a Svelte/TypeScript frontend for a modern user interface.

## Features

* File uploading with preview
* Storage of files on the local/self-hosted server
* PostgreSQL database integration for storing file metadata

## Getting Started

### Requirements

* Docker and Docker Compose
* Git (for cloning the repository)

### Quick Start with Docker

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/nedmed.git
   cd nedmed
   ```

2. Create a `.env` file in the root directory:
   ```env
   DATABASE_URL=postgresql://user:password@host:port/database
   DB_USER=user
   DB_PASSWORD=password
   DB_DATABASE=database
   ```

3. Start the application:
   ```bash
   docker-compose up --build
   ```

The application will be available at:
http://localhost:3000

### Manual Setup (Development)

If you prefer to run the services without Docker, follow these steps:

#### Requirements
* Go (https://go.dev/doc/install)
* Node.js and npm/pnpm/yarn (https://nodejs.org/)
* PostgreSQL

#### Backend (Go)
1. Navigate to backend: `cd go`
2. Install dependencies: `go mod tidy`
3. Create `.env` file with your PostgreSQL details
```
DATABASE_URL=postgresql://user:password@host:port/database
  ```
4. Run: `go run ./cmd/NedMed`

#### Frontend (SvelteKit)
1. Navigate to frontend: `cd src`
2. Install dependencies: `npm install`
3. Run development server: `npm run dev -- --open`

## Deployment (Self-Hosting)

### Using Docker (Recommended)
1. Adjust the `.env` file with your production credentials
2. Start services: `docker-compose up -d`
3. Configure your reverse proxy (Nginx, Apache, Caddy) to forward traffic to the containers

### Manual Deployment
1. Build backend: `go build -o nedmed ./cmd/NedMed`
2. Build frontend: `npm run build`
3. Copy the built files to your server
4. Configure your web server and database
5. Start the services using process managers (systemd, pm2)

## Future Enhancements

* User Authentication
* File Management (delete, rename)