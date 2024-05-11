-- create "creations" table
CREATE TABLE "creations" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "creation_type" character varying NOT NULL DEFAULT 'txt2img', "parameter" character varying NOT NULL DEFAULT '', "url" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "creations"
COMMENT ON COLUMN "creations" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "creations"
COMMENT ON COLUMN "creations" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "creations"
COMMENT ON COLUMN "creations" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "creations"
COMMENT ON COLUMN "creations" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "creations"
COMMENT ON COLUMN "creations" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "creations"
COMMENT ON COLUMN "creations" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "creation_type" on table: "creations"
COMMENT ON COLUMN "creations" ."creation_type" IS '创作类型';
-- set comment to column: "parameter" on table: "creations"
COMMENT ON COLUMN "creations" ."parameter" IS '参数';
-- set comment to column: "url" on table: "creations"
COMMENT ON COLUMN "creations" ."url" IS '作品链接';
-- set comment to column: "user_id" on table: "creations"
COMMENT ON COLUMN "creations" ."user_id" IS '用户 id';
