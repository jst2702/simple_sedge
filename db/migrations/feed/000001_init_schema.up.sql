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
  "email" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "documents" ("ticker");

CREATE INDEX ON "documents" ("ticker", "published_at");

CREATE INDEX ON "documents" ("headline", "published_at");

CREATE INDEX ON "models" ("name");

CREATE INDEX ON "sentiment_scores" ("document_guid");

CREATE INDEX ON "sentiment_scores" ("sentiment", "confidence");

ALTER TABLE "sentiment_scores" ADD FOREIGN KEY ("model_id") REFERENCES "models" ("id");

ALTER TABLE "sentiment_scores" ADD FOREIGN KEY ("document_guid") REFERENCES "documents" ("guid");
