-- Create "prefectures" table
CREATE TABLE "public"."prefectures" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "elias" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_prefectures_deleted_at" to table: "prefectures"
CREATE INDEX "idx_prefectures_deleted_at" ON "public"."prefectures" ("deleted_at");
-- Create "cities" table
CREATE TABLE "public"."cities" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "name_kana" text NULL,
  "prefecture_id" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_cities_prefecture" FOREIGN KEY ("prefecture_id") REFERENCES "public"."prefectures" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_cities_deleted_at" to table: "cities"
CREATE INDEX "idx_cities_deleted_at" ON "public"."cities" ("deleted_at");
