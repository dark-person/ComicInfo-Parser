-- Category table 
CREATE TABLE "category" (
    "category_id" INTEGER,
    "category_name" VARCHAR(255) NOT NULL,
    PRIMARY KEY("category_id" AUTOINCREMENT)
);

INSERT INTO
    "category"(category_name)
VALUES
    ("Genre"),
    ("Publisher"),
    ("Translator");

-- ===========================================================
-- Copy List inputted
CREATE TABLE "_tmp_list_inputted" (
    "category" TEXT NOT NULL,
    "input" TEXT NOT NULL,
    PRIMARY KEY("category", "input")
);

INSERT INTO
    "_tmp_list_inputted"
SELECT
    *
FROM
    list_inputted;

-- Update copied list inputted by new genre value
UPDATE
    "_tmp_list_inputted"
SET
    category = 1
WHERE
    category = 'Genre';

UPDATE
    "_tmp_list_inputted"
SET
    category = 2
WHERE
    category = 'Publisher';

UPDATE
    "_tmp_list_inputted"
SET
    category = 3
WHERE
    category = 'Translator';

-- Drop existing table
DROP TABLE IF EXISTS "list_inputted";

-- Create new table structure with same name
CREATE TABLE "list_inputted" (
    "category" INTEGER NOT NULL,
    "input" TEXT NOT NULL,
    FOREIGN KEY("category") REFERENCES "category"("category_id") ON UPDATE CASCADE,
    PRIMARY KEY("category", "input")
);

INSERT INTO
    "list_inputted" ("category", "input")
SELECT
    CAST("category" AS INTEGER),
    "input"
FROM
    _tmp_list_inputted;

-- Drop temporary table
DROP TABLE "_tmp_list_inputted";

-- SQLite version
PRAGMA user_version = 2;