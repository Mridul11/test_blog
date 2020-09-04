CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "todoes" (
"id" TEXT PRIMARY KEY,
"title" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, "description" TEXT NOT NULL DEFAULT '');
CREATE TABLE IF NOT EXISTS "widgets" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"body" TEXT,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
