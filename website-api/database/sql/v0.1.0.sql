-- Migration for English Step v0.1.0
-- Creates the SynonymExplanations table for storing AI-generated synonym explanations

CREATE TABLE IF NOT EXISTS synonym_explanations (
    id SERIAL PRIMARY KEY,
    vocabularies TEXT[] NOT NULL,
    explanation TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (vocabularies)
);

-- To ensure uniqueness regardless of order, consider storing vocabularies in sorted order in the application layer before insertion.
-- Optionally, you can add an index for faster lookup:
CREATE INDEX IF NOT EXISTS idx_synonym_explanations_vocabularies ON synonym_explanations USING GIN (vocabularies);
