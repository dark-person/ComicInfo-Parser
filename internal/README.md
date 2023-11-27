# Internal Package

This folder contains several Internal Package.

## archive

The package to convert image directory to `.zip` and `.cbz` file.

## comicinfo

The ComicInfo Structure in `Go`, Converted from `ComicInfo.xsd`.

Current Schema version is `2.1`.

## parser

Utils Package. Extract information from filename/directory name to comicInfo.

Currently Support `Author`, `Title`.

## scanner

The package for scanner image directory, also for identify special tags.

This package will produce a `ComicInfo` Struct, which contains pages detail, and also information that extract from `parser` package.
