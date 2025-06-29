# English Step: System Design Documentation

**Document Version:** 0.1.0  
**Date:** June 25, 2025  

## Introduction

This document describes how to design the features of version 0.1.0 based on the [Feature Spec](feature-specification.md).

## Website API

### Overview
The Website API serves as the backend for the Synonym module, handling requests from the frontend, interacting with the database, and integrating with the AI service for synonym explanations.

### Endpoints

- `POST /api/synonyms/explain`
  - **Description:** Accepts a list of vocabulary words and returns an AI-generated explanation of their synonyms.
  - **Request Body:**
    ```json
    {
      "vocabularies": ["word1", "word2", ...]
    }
    ```
  - **Validation:** Requires at least two vocabulary words.
  - **Response:**
    - If an explanation for the given vocabularies exists in the database, return the stored result.
    - Otherwise, generate a new explanation using the AI service, store it, and return the result.
  - **Response Body:**
    ```json
    {
      "vocabularies": ["word1", "word2", ...],
      "explanation": "...AI-generated explanation..."
    }
    ```

## Database Design

- **Table: SynonymExplanations**
  - `id`: Primary key
  - `vocabularies`: Array or JSON of vocabulary words (unique constraint on sorted values)
  - `explanation`: Text
  - `language`: Language code of the explanation (e.g., 'en', 'zh')
  - `created_at`: Timestamp

## AI Integration

- Integrate with Gemini AI service to generate explanations for given vocabulary sets.
- Cache/store results in the database to avoid redundant API calls for the same vocabulary sets.

## Frontend Integration

- Add a Synonym module accessible from the homepage and header.
- Create a dedicated page for users to input multiple vocabulary words.
- Validate input to ensure at least two words are provided before submitting.
- Display the AI-generated explanation or a relevant error message.

## Flow Summary

1. User navigates to the Synonym page and inputs vocabulary words.
2. Frontend validates input and sends a POST request to `/api/synonyms/explain`.
3. Backend checks the database for an existing explanation.
4. If found, returns the stored explanation. If not, calls the AI service, stores the result, and returns it.
5. Frontend displays the explanation to the user.

## Security Considerations

### CORS (Cross-Origin Resource Sharing)

- The API should implement CORS to allow requests only from trusted origins (e.g., the official website domain).
- The allowed domain(s) for CORS should be defined in the environment file (e.g., `.env`), allowing for easy configuration and deployment flexibility.
- Configure CORS middleware in the Gin framework to restrict allowed origins, methods (e.g., POST), and headers as needed.
- Example: In a Go/Gin backend, use a CORS middleware (such as `github.com/gin-contrib/cors`) and set the allowed origin from an environment variable such as `CORS_ALLOWED_ORIGIN`.

### CSRF (Cross-Site Request Forgery)

- Implement CSRF protection for all state-changing endpoints (such as POST requests).
- Use CSRF tokens that are validated on each request, or rely on secure, same-site cookies if using a modern SPA setup.
- Example: In a Go/Gin backend, use a CSRF middleware (such as `github.com/utrack/gin-csrf`) and configure the secret/key via environment variables.
