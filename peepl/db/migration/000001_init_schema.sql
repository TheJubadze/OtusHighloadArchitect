-- +goose Up

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "login" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "firstname" varchar,
  "lastname" varchar,
  "birthdate" date,
  "sex" varchar,
  "interests" varchar,
  "city_id" integer NOT NULL,
  "role_id" integer NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "cities" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL UNIQUE
);

CREATE TABLE "user_roles" (
  "id" serial PRIMARY KEY,
  "role" varchar NOT NULL UNIQUE
);

ALTER TABLE "users" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "user_roles" ("id");

-- +goose Down

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS cities;
