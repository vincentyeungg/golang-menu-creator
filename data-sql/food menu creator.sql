CREATE TABLE "MenuItem" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" numeric NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_by" varchar NOT NULL,
  "status" varchar(1) DEFAULT 'A'
);

CREATE TABLE "Ingredient" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "brand_name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_by" varchar NOT NULL,
  "status" varchar(1) DEFAULT 'A'
);

CREATE TABLE "MenuItem_Ingredient" (
  "id" serial PRIMARY KEY,
  "food_id" int NOT NULL,
  "ingredient_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_by" varchar NOT NULL,
  "status" varchar(1) DEFAULT 'A'
);

CREATE TABLE "Menu" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_by" varchar NOT NULL,
  "status" varchar(1) DEFAULT 'A'
);

CREATE TABLE "Menu_MenuItem" (
  "id" serial PRIMARY KEY,
  "menu_id" int NOT NULL,
  "food_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_by" varchar NOT NULL,
  "status" varchar(1) DEFAULT 'A'
);

CREATE INDEX ON "MenuItem" ("id");

CREATE INDEX ON "Ingredient" ("id");

CREATE INDEX ON "Menu" ("id");

ALTER TABLE "MenuItem_Ingredient" ADD FOREIGN KEY ("food_id") REFERENCES "MenuItem" ("id");

ALTER TABLE "MenuItem_Ingredient" ADD FOREIGN KEY ("ingredient_id") REFERENCES "Ingredient" ("id");

ALTER TABLE "Menu_MenuItem" ADD FOREIGN KEY ("menu_id") REFERENCES "Menu" ("id");

ALTER TABLE "Menu_MenuItem" ADD FOREIGN KEY ("food_id") REFERENCES "MenuItem" ("id");
