# Internal Package

This folder contains several Internal Package.

## application

Package for wails app.

## archive

The package to convert image directory to `.zip` and `.cbz` file.

## comicinfo

The [ComicInfo.xml](https://anansi-project.github.io/docs/comicinfo/documentation) Structure in `Go`, Converted from `ComicInfo.xsd`.

Current Schema version is `2.1`.

## database

The package to manipulate the database. Current `user_version` is `1`.

## files

The package contains multiple utility function for files system.

## history

The package to control record of database, specially user inputted values.

## parser

Utils Package. Extract information from filename/directory name to comicInfo.

Currently Support `Author`, `Title`, also for identify special tags.

## scanner

The package for scanner image directory.

This package will produce a `ComicInfo` Struct, which contains pages detail, and also information that extract from `parser` package.
