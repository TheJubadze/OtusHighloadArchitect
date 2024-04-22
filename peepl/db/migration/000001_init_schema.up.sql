CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "birthdate" date NOT NULL,
  "sex" varchar NOT NULL,
  "interests" varchar,
  "city_id" integer NOT NULL,
  "role_id" integer NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "cities" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "user_roles" (
  "id" serial PRIMARY KEY,
  "role" varchar NOT NULL
);

ALTER TABLE "users" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "user_roles" ("id");
