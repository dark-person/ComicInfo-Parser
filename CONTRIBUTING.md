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

[Github Flow](https://guides.github.com/introduction/flow/index.html) will be used in this project, so all code changes happen through pull requests.

### Version Release

This project will use `git` tag for marking release version, in branch `main` ONLY.

The version tag should be in this format: `v1.2.3`. For pre-release version, use format `v1.2.3-alpha20231228` instead.

### Use a Consistent Coding Style

If you using vscode, then you will find the coding style is prepared in `setting.json`.

Don't try to change this file unless you have discuss with other developers. In most cases, we suggest you should open a new issue.

### Create Pull Request

Developer are encouraged to contribute their code through pull request. You should ensure your pull request is meaningful, and not contains any unnecessary / harmful code.

After have a self-review for your code, you are free to create a new pull request with below steps:

1. Fork the repo and create your branch from `develop` / `main`.
2. Make modification to your branch.
3. Create a new pull request in Github, with template provided. Depend on your code:
    1. Create normal pull request, if code fulfills the requirements in checklist.
    2. Create draft pull request, if code require further modification

It is recommended to use `git rebase develop` on your code before creating a new pull request.

## References

This document was adapted from the open-source contribution guidelines for [Facebook's Draft](https://github.com/facebook/draft-js/blob/a9316a723f9e918afde44dea68b5f9f39b7d9b00/CONTRIBUTING.md)

_Last updated: 2024-03-03_
