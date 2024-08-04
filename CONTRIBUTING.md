# Contributing Guide

Please read this guide for how to make contributing to this project, which includes:

-   User contributions
    -   Reporting bugs
    -   Feature Requests
-   Developer contributions
    -   Version Release
    -   Use a Consistent Coding Style
    -   Create Pull Request

## Contributing Guide for User

### Report Bugs

We ONLY use GitHub issues to track public bugs. Report a bug by opening a new issue.

**Great Bug Reports** tend to have:

-   A quick summary and/or background
-   Steps to reproduce
    -   Be specific & precise
    -   Give sample if you can
-   What you expected would happen
-   What actually happens
-   Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

### Feature Requests

Please open a new issue, with label `feature request`. It suggested that include these when creating a new feature request:

-   Few line to describe feature
-   The reason for why this feature helps others
-   A brief example of this feature, which include input & expected outcome
-   (Optional) An image to visualize if you want to request some GUI changes

## Contributing Guide for Developer

The code is hosted at github, which performs:

-   track issues
-   feature requests
-   accept pull requests

### Git workflows

We have only one main branch `master`.

This project will use `git` tag for marking release version, format SHOULD be:

-   For stable version: `v1.2.3`
-   For pre-release version: `v1.2.3-alpha20231228`

### Hotfix

Hotfix will be created on BOTH target version & developing version.

For example, `v1.4.0` has CVE issue, current developing `v1.5.0-alpha20250228` in `master`, Then:

1. Create branch `v1.4` on tag `v1.4.0`
2. Fix issue in single commit `bce4df1`
3. Tag `v1.4.1` on commit `bce4df1`
4. Cherry-pick `bce4df1` into `master` branch

### Use a Consistent Coding Style

If you using vscode, then you will find the coding style is prepared in `setting.json`.

Don't try to change this file unless you have discuss with other developers. In most cases, we suggest you should open a new issue.

### Create Pull Request

Developer are encouraged to contribute their code through pull request. You should ensure your pull request is meaningful, and not contains any unnecessary / harmful code.

Before start coding, developers SHOULD refer to `DEVELOPMENT.md` for development guide.

After have a self-review for your code, you are free to create a new pull request with below steps:

1. Fork the repo and create your branch from `master`.
2. Make modification to your branch.
3. Create a new pull request in Github, with template provided. Depend on your code:
    1. Create normal pull request, if code fulfills the requirements in checklist.
    2. Create draft pull request, if code require further modification

## References

This document was adapted from the open-source contribution guidelines for [Facebook's Draft](https://github.com/facebook/draft-js/blob/a9316a723f9e918afde44dea68b5f9f39b7d9b00/CONTRIBUTING.md)

_Last updated: 2024-08-03_
