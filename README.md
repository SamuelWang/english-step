# English Step

A monorepo project containing a Next.js frontend and a Go backend API.

## Folder Structure
- `website/` – Next.js frontend application (TypeScript)
- `website-api/` – Go backend API service

## Documentations

- [Specifications](specs)
- [Development Guideline](DEVELOPMENT.md)
- [Backend README](website-api/README.md)
- [Backend Development Guideline](website-api/DEVELOPMENT.md)

## Getting Started

## Sub-projects
- Frontend code: `website/`
- Backend code: `website-api/`

### Environment Variables
1. Copy the file of the `.env.example`.
2. Set values of variables.

### Database
1. Run `cd website-api` to switch the working folder to `website-api` folder.
2. Run `./database/scripts/atlas/migrate-apply.sh` to apply the DB schema to the target database.

### Frontend
1. Run `cd website` to switch the working folder to `website` folder.
2. Install dependencies: `pnpm install`.
3. Run the development server: `pnpm dev`.

### Backend
1. Run `cd website-api` to switch the working folder to `website-api` folder.
2. Run the API server: `go run .`.

See [DEVELOPMENT.md](DEVELOPMENT.md) for more details.
