-- reverse: set comment to column: "status" on table: "creations"
COMMENT ON COLUMN "creations" ."status" IS '';
-- reverse: modify "creations" table
ALTER TABLE "creations" DROP COLUMN "status";
