-- Rename current table to tmp
ALTER TABLE
    tags RENAME TO tmp_tags;

-- Create new tags table
CREATE TABLE "tags" (
    "tag_id" INTEGER,
    "input" TEXT UNIQUE,
    PRIMARY KEY("tag_id" AUTOINCREMENT)
);

-- Create new table for tag alias
CREATE TABLE "tags_alias" (
    "alias" TEXT,
    "tag_id" INTEGER,
    PRIMARY KEY("alias")
);

-- Insert temp records back to new tag table
INSERT INTO
    "tags" ("input")
SELECT
    ("input")
FROM
    tmp_tags;

-- Drop temp table
DROP TABLE "tmp_tags";

-- Create view for tag & tag alias
CREATE VIEW "view_tags_alias" AS
SELECT
    tags_alias.alias,
    tags.input as tag,
    tags.tag_id as tag_id
FROM
    "tags_alias"
    LEFT JOIN tags on tags_alias.tag_id = tags.tag_id;

-- Create view for easier lookup
CREATE VIEW map_keyword_tags as
SELECT
    tags_alias.alias as keyword,
    tags.input as tag
FROM
    tags_alias
    JOIN tags ON tags_alias.tag_id = tags.tag_id
UNION
SELECT
    tags.input as keyword,
    tags.input as tag
FROM
    tags;

-- SQLite version
PRAGMA user_version = 3;