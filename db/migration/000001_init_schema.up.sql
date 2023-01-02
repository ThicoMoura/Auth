-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-01-02T19:00:16.024Z

CREATE TABLE "group" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "active" bool NOT NULL DEFAULT true
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "group" uuid NOT NULL,
  "email" varchar NOT NULL,
  "name" varchar NOT NULL,
  "pass" varchar NOT NULL,
  "active" bool NOT NULL DEFAULT true
);

CREATE TABLE "system" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "active" bool NOT NULL DEFAULT true
);

CREATE TABLE "access" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "system" uuid NOT NULL,
  "table" varchar UNIQUE NOT NULL,
  "type" varchar[]
);

CREATE TABLE "session" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user" uuid NOT NULL,
  "token" varchar NOT NULL,
  "ip" varchar NOT NULL,
  "agent" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expires_at" timestamptz NOT NULL
);

ALTER TABLE "user" ADD FOREIGN KEY ("group") REFERENCES "group" ("id");

ALTER TABLE "access" ADD FOREIGN KEY ("system") REFERENCES "system" ("id");

ALTER TABLE "session" ADD FOREIGN KEY ("user") REFERENCES "user" ("id");

CREATE TABLE "group_access" (
  "group_id" uuid,
  "access_id" uuid,
  PRIMARY KEY ("group_id", "access_id")
);

ALTER TABLE "group_access" ADD FOREIGN KEY ("group_id") REFERENCES "group" ("id");

ALTER TABLE "group_access" ADD FOREIGN KEY ("access_id") REFERENCES "access" ("id");


CREATE TABLE "user_access" (
  "user_id" uuid,
  "access_id" uuid,
  PRIMARY KEY ("user_id", "access_id")
);

ALTER TABLE "user_access" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_access" ADD FOREIGN KEY ("access_id") REFERENCES "access" ("id");

