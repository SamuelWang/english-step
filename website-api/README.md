# website-api

This is the backend API service for the English Step project.

## Overview
This service provides API endpoints for the English Step application. It is written in Go and is designed to be stateless and configurable via environment variables.

## Getting Started

### Prerequisites
- Go 1.20 or newer
- (Optional) [godotenv](https://github.com/joho/godotenv) for local development

### Installation
1. Clone the repository.
2. Navigate to the `website-api` directory:
   ```sh
   cd website-api
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

### Environment Variables
Configuration is managed via environment variables. See `.env.example` for required variables. To set up your environment:

1. Copy `.env.example` to `.env`:
   ```sh
   cp .env.example .env
   ```
2. Edit `.env` and fill in your values.

### Running the API
To run the API locally:

```sh
go run main.go
```
