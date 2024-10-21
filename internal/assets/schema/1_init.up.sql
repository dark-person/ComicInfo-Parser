-- Table for inputted record value in GUI
CREATE TABLE "list_inputted" (
    "category" TEXT NOT NULL,
    "input" TEXT NOT NULL,
    PRIMARY KEY("category", "input")
);

-- Master table for tags.
CREATE TABLE "tags" (
    "input" TEXT NOT NULL,
    PRIMARY KEY("input")
);

-- SQLite version
PRAGMA user_version = 1