-- New category
INSERT INTO
    category (category_id, category_name)
VALUES
    (6, 'Characters'),
    (7, 'Teams'),
    (8, 'Locations');

-- SQLite version
PRAGMA user_version = 6;