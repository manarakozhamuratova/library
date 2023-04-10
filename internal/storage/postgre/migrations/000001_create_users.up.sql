create table  users (
  "id" serial primary key,
  "created_at" timestamp default now(),
  "username" VARCHAR(100) UNIQUE,
  "email" VARCHAR(100) UNIQUE,
  "password" TEXT
);