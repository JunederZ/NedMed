# NedMed: Self-Hosted Media Uploader

NedMed is a simple self-hosted application for uploading and storing media files on your local machine or self-hosted server, with PostgreSQL integration for metadata storage. It uses a Go backend for handling file uploads and a Svelte/TypeScript frontend for a modern user interface.

## Features

* File uploading with preview.
* Storage of files on the local/self-hosted server.
* PostgreSQL database integration for storing file metadata.


## Getting Started

### Requirements

* **Go:** Install Go (https://go.dev/doc/install).
* **Node.js and npm/pnpm/yarn:** Install Node.js and your preferred package manager (https://nodejs.org/).
* **PostgreSQL:** Install and run PostgreSQL.  Ensure you have the necessary database credentials.


### Backend (Go)

1. **Navigate to the backend directory (`go`):**  `cd go`

2. **Install dependencies:** `go mod tidy`

3. **Set up database connection:** Create a `.env` file in the `go` directory with your PostgreSQL connection details (adjust as needed):

```
DATABASE_URL=postgresql://user:password@host:port/database
  ```

4. **Run the server:** (replace `main` with your main package) `go run ./cmd/main`  or `go build ./cmd/main && ./main`


### Frontend (SvelteKit)

1. **Navigate to the frontend directory (`src`):** `cd src`
2. **Install dependencies:** `npm install` (or `pnpm install`, `yarn install`)
3. **Run the development server:** `npm run dev -- --open` (or `pnpm dev -- --open`, `yarn dev -- --open`)


## Deployment (Self-Hosting)

1. **Build the Go backend:** `go build -o nedmed ./cmd/main`
2. **Build the SvelteKit frontend:** `npm run build` (or `pnpm build`, `yarn build`)
3. **Copy the artifacts:** Copy the Go executable (`nedmed`) and the `.svelte-kit/output/server` directory contents to your server.  Also, copy the  `.svelte-kit/output/client` directory contents to a location accessible by your web server.
4. **Run the backend:** Start your Go server (e.g., using `systemd`, `pm2`).
5. **Serve the frontend:** Configure a web server (Nginx, Apache, Caddy) to serve the frontend files from `.svelte-kit/output/client` and reverse proxy requests to your Go backend as necessary. Set up your database on the server.


## Future Enhancements

* User Authentication
* File Management (delete, rename)
