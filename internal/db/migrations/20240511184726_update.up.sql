-- modify "creations" table
ALTER TABLE "creations" ADD COLUMN "status" character varying NOT NULL DEFAULT 'created';
-- set comment to column: "status" on table: "creations"
COMMENT ON COLUMN "creations" ."status" IS '创作状态';
