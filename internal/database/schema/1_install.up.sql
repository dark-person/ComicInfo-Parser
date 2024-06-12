-- List for available prefix
CREATE TABLE "list_prefix" (
	"prefix_name" TEXT NOT NULL,
	PRIMARY KEY("prefix_name")
);

-- List for available suffix
CREATE TABLE "list_suffix" (
	"suffix_name" TEXT NOT NULL,
	PRIMARY KEY("suffix_name")
);

-- Mapping for alias, allow one tags_id has multiple alias.
CREATE TABLE "tag_alias" (
	"alias_id" INTEGER,
	"tag_id" INTEGER,
	"alias_name" TEXT NOT NULL,
	PRIMARY KEY("alias_id" AUTOINCREMENT)
);

-- Master table for tags.
CREATE TABLE "tags" (
	"tag_id" INTEGER,
	"prefix" TEXT,
	"name" TEXT NOT NULL,
	"suffix" TEXT,
	PRIMARY KEY("tag_id" AUTOINCREMENT)
);

-- View for tag full name
CREATE VIEW view_tags AS
SELECT
	tag_id,
	prefix_tmp || name || suffix_tmp full_name
FROM
	(
		SELECT
			tag_id,
			CASE
				WHEN prefix IS NULL THEN ""
				ELSE CASE
					WHEN prefix = "" THEN ""
					ELSE prefix || ":"
				END
			END prefix_tmp,
			name,
			CASE
				WHEN suffix IS NULL THEN ""
				ELSE CASE
					WHEN suffix = "" THEN ""
					ELSE "(" || suffix || ")"
				END
			END suffix_tmp
		FROM
			tags
	);

-- View for alias with original tag
CREATE VIEW view_alias AS
SELECT
	alias_id,
	alias_name,
	tags.prefix,
	tags.name,
	tags.suffix
FROM
	tag_alias
	LEFT JOIN tags ON tags.tag_id = tag_alias.tag_id;

-- SQLite version
PRAGMA user_version = 1