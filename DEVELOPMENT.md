# Development Documentation

This document describes the information of the project to help developers understand how to develop the project.

## Folder Structure

This section describes the folder structure of the project to help developers understand the organization and purpose of each directory.

### Root Directory
- `website/` - The main frontend application, built with Next.js and TypeScript.
- `website-api/` - The backend API service, implemented in Go.

### website/
This directory contains the frontend web application.

- `app/` - Main application source code (Next.js app directory).
- `public/` - Static assets served directly (images, SVGs, etc.).
- `eslint.config.mjs` - ESLint configuration for code linting.
- `next-env.d.ts` - TypeScript types for Next.js.
- `next.config.ts` - Next.js configuration file.
- `package.json` - Project dependencies and scripts.
- `pnpm-lock.yaml` - Lockfile for pnpm package manager.
- `postcss.config.mjs` - PostCSS configuration.
- `README.md` - Project overview and instructions.
- `tsconfig.json` - TypeScript configuration.

### website-api/
This directory contains the backend API service.

- `go.mod` - Go module definition.
- `go.sum` - Go module checksums.

### Notes
- All frontend development should occur in the `website/` directory.
- All backend/API development should occur in the `website-api/` directory.
- Shared assets or documentation can be placed in the root directory if needed.

For more details on each part, refer to the respective README files or documentation within each directory.

## Database

### Generating Migration Scripts

When the database models are update, it should generate the migration scripts for upgrading.

1. Run `cd website-api` to switch the working folder to `website-api` folder.
2. Run `./database/scripts/atlas/migrate-diff.sh` to generate the new migration script file.
