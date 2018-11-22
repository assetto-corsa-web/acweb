ALTER TABLE "configurations"
ADD COLUMN "auth_plugin_address" integer NOT NULL;

ALTER TABLE "cars"
ADD COLUMN "ballast" integer NOT NULL,
ADD COLUMN "restrictor" integer NOT NULL;
