-- Table for inputted record value in GUI
CREATE TABLE "list_inputted" (
    "category" TEXT NOT NULL,
    "input" TEXT NOT NULL,
    PRIMARY KEY("category", "input")
);

-- SQLite version
PRAGMA user_version = 2