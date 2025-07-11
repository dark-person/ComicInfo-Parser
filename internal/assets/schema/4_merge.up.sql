CREATE TABLE "word_store" (
    "word_id" INTEGER,
    "word" TEXT,
    "category_id" INTEGER,
    PRIMARY KEY("word_id" AUTOINCREMENT),
    UNIQUE("word", "category_id")
);

CREATE TABLE "triggers" (
    "trigger_id" INTEGER,
    "keyword" TEXT,
    "word_id" INTEGER,
    PRIMARY KEY("trigger_id" AUTOINCREMENT)
);

-- New category
INSERT INTO
    category (category_id, category_name)
VALUES
    (4, 'Tag');

-- Migration from list_inputted & tags
INSERT INTO
    word_store (word, category_id)
SELECT
    input AS word,
    category AS category_id
FROM
    list_inputted;

INSERT INTO
    word_store (word, category_id)
SELECT
    input AS word,
    4 AS category_id
FROM
    tags;

-- Migration from tag_alias
INSERT INTO
    triggers (keyword, word_id)
SELECT
    tags_alias.alias,
    word_store.word_id
FROM
    tags_alias
    LEFT JOIN tags ON tags.tag_id = tags_alias.tag_id
    LEFT JOIN word_store ON tags.input = word_store.word
    AND word_store.category_id = 4;

-- New views on trigger and word
CREATE VIEW view_triggers_words AS
SELECT
    trigger_id,
    keyword,
    word,
    category_name,
    word_store.category_id
FROM
    triggers
    LEFT JOIN word_store ON word_store.word_id = triggers.word_id
    LEFT JOIN category ON category.category_id = word_store.category_id;

-- DROP old tables
DROP TABLE list_inputted;

DROP TABLE tags;

DROP TABLE tags_alias;

DROP VIEW view_tags_alias;

DROP VIEW map_keyword_tags;

-- SQLite version
PRAGMA user_version = 4;