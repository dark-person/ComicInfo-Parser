-- Delete view & alias tag table
DROP VIEW "view_tags_alias";

DROP TABLE "tags_alias";

DROP TABLE "map_keyword_tags";

-- Create a temp table of tags 
CREATE TABLE "tmp_tags" (
    "input" TEXT NOT NULL,
    PRIMARY KEY("input")
);

-- Copy data from the old table to the new table.
INSERT INTO
    "tmp_tags" ("input")
SELECT
    ("input")
FROM
    "tags";

-- Drop the old table.
DROP TABLE "tags";

-- Rename the new table to the original table's name.
ALTER TABLE
    "tmp_tags" RENAME TO "tags";

-- SQLite version
PRAGMA user_version = 2;