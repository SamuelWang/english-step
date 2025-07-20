-- Create "synonym_explanations" table
CREATE TABLE "synonym_explanations" (
  "id" bigserial NOT NULL,
  "vocabularies" json NULL,
  "explanation" text NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
