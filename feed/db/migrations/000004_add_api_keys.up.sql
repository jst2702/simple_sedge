CREATE TABLE "api_keys" (
  "api_key" varchar PRIMARY KEY,
  "active" bool DEFAULT true NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "documents" ADD COLUMN "api_key_used" VARCHAR NOT NULL;

ALTER TABLE "documents" ADD FOREIGN KEY ("api_key_used") REFERENCES "api_keys" ("api_key");