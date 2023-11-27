# GUI ComicInfo Parser

## About

 [ComicInfo.xml](https://anansi-project.github.io/docs/comicinfo/documentation)  is a metadata for manga/comic. It is used in some self-hosted app, e.g. `komga`.

This Project is aim to provide a simple GUI for create `ComicInfo.xml` and `.cbz` archive at easy way.

## Feature Available

### Quick Export (Komga Only)

Create a directory that can copy directly to `komga` comic directory.

Already contains a `.cbz` archive and `ComicInfo.xml` at correct location.

## Develop

This project using `go` language, with `wails` framework and `react` typescript as frontend.

### Live Development Mode

To run in live development mode, run `wails dev` in the project directory.

This will run a Vite development server that will provide very fast hot reload of your frontend changes.

If you want to develop in a browser and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect to this in your browser, and you can call your Go code from devtools.

### Building

To build a re-distributable, production mode package, use `wails build`.
