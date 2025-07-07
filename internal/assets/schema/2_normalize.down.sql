-- Temp table with no foreign keys
CREATE TABLE "temp_list_inputted" (
    "category" TEXT NOT NULL,
    "input" TEXT NOT NULL,
    PRIMARY KEY("category", "input")
);

-- Copy existing table contents with category value mapped
INSERT INTO
    "temp_list_inputted" ("category", "input")
SELECT
    category_name AS category,
    input
FROM
    list_inputted
    LEFT JOIN category
WHERE
    category.category_id = list_inputted.category;

-- Drop table not supported
DROP TABLE IF EXISTS "list_inputted";

DROP TABLE IF EXISTS "category";

-- Change temp table name
ALTER TABLE
    "temp_list_inputted" RENAME TO "list_inputted";

-- SQLite version
PRAGMA user_version = 1;