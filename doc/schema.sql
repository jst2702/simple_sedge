-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-03-28T23:35:56.494Z

CREATE TABLE "documents" (
  "guid" varchar PRIMARY KEY,
  "url" varchar NOT NULL,
  "site" varchar NOT NULL,
  "site_full" varchar NOT NULL,
  "site_section" varchar NOT NULL,
  "headline" varchar NOT NULL,
  "title" varchar NOT NULL,
  "body" varchar NOT NULL,
  "ticker" varchar,
  "tickers" varchar[],
  "published_at" timestamptz NOT NULL,
  "language" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "models" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "role" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sentiment_scores" (
  "id" bigserial PRIMARY KEY,
  "model_id" bigserial NOT NULL,
  "document_guid" varchar NOT NULL,
  "sentiment" bigint NOT NULL,
  "confidence" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "is_email_verified" bool NOT NULL DEFAULT false,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

CREATE TABLE "sessions" (
  "id" bigserial PRIMARY KEY,
  "username" varchar,
  "refresh_token" varchar,
  "user_agent" varchar,
  "client_up" varchar,
  "is_blocked" boolean,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "documents" ("ticker");

CREATE INDEX ON "documents" ("ticker", "published_at");

CREATE INDEX ON "documents" ("headline", "published_at");

CREATE INDEX ON "models" ("name");

CREATE INDEX ON "sentiment_scores" ("document_guid");

CREATE INDEX ON "sentiment_scores" ("sentiment", "confidence");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "sessions" ("username");

CREATE INDEX ON "sessions" ("created_at");

ALTER TABLE "sentiment_scores" ADD FOREIGN KEY ("model_id") REFERENCES "models" ("id");

ALTER TABLE "sentiment_scores" ADD FOREIGN KEY ("document_guid") REFERENCES "documents" ("guid");

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
