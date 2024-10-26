# Develop Guidelines

This project using `go` language, with `wails` framework and `react` typescript as frontend.

## Live Development Mode

To run in live development mode, run `wails dev` in the project directory.

This will run a Vite development server that will provide very fast hot reload of your frontend changes.

If you want to develop in a browser and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect to this in your browser, and you can call your Go code from devtools.

## Building

To build a re-distributable, production mode package, use `wails build`.

## Makefile

A `Makefile` is provided for easy development. Please enter `make help` for available commands.

If you are using Window OS, and would like to use make, consider using Chocolatey (link) command:

```
choco install make
```

## Git related

### Branch Management

The git branch management should be followed below principle:

1. Ensure branch has `rebase` to latest branch.
2. Ensure your commit history should be simple as possible, i.e. you should not have multiple commits for typo fixes.
3. File changes in each commit should be as few as possible. This is to prevent `rebase` difficulty.
    - i.e. if you change `module1` & `module2`, you should not use one commit ONLY to conclude both changes, instead use >2 commits depend on complexity of these changes.

_Last updated: 2024-08-03_
