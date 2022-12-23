-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-12-23T18:18:20.746Z

CREATE TABLE "group" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "active" bool NOT NULL DEFAULT true
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "group" uuid NOT NULL,
  "cpf" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "pass" varchar NOT NULL,
  "active" bool NOT NULL DEFAULT true
);

CREATE TABLE "system" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "active" bool NOT NULL DEFAULT true
);

CREATE TABLE "access" (
  "id" uuid PRIMARY KEY,
  "system" uuid NOT NULL,
  "table" varchar UNIQUE NOT NULL,
  "type" varchar[]
);

ALTER TABLE "user" ADD FOREIGN KEY ("group") REFERENCES "group" ("id");

ALTER TABLE "access" ADD FOREIGN KEY ("system") REFERENCES "system" ("id");

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

