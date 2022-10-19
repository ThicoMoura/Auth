CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "cpf" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "pass" varchar NOT NULL,
  "active" bool NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS "groups" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "user_group" (
  "user" bigint NOT NULL,
  "group" bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS "permission" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "create" bool DEFAULT false,
  "read" bool DEFAULT false,
  "update" bool DEFAULT false,
  "delete" bool DEFAULT false
);

CREATE TABLE IF NOT EXISTS "group_permission" (
  "group" bigint NOT NULL,
  "permission" bigint NOT NULL
);

CREATE INDEX ON "user_group" ("user");

CREATE INDEX ON "user_group" ("group");

CREATE INDEX ON "group_permission" ("group");

CREATE INDEX ON "group_permission" ("permission");

ALTER TABLE "user_group" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");

ALTER TABLE "user_group" ADD FOREIGN KEY ("group") REFERENCES "groups" ("id");

ALTER TABLE "group_permission" ADD FOREIGN KEY ("group") REFERENCES "groups" ("id");

ALTER TABLE "group_permission" ADD FOREIGN KEY ("permission") REFERENCES "permission" ("id");
