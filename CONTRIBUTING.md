# Contributing Guide

Thanks for contributing to this project! :+1: This project and everyone participating in it is governed by the [Octopus Deploy Code of Conduct](https://github.com/OctopusDeploy/.github/blob/main/CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior using the instructions in the code of conduct.

This guide provides an overview of the contribution workflow from opening an issue, creating a PR, reviewing, and merging the PR.

## Getting Started

This project is built, tested, and released by workflows defined in GitHub Actions (see [Actions](../actions/) for more information). Release management is controlled through [Release-Please](https://github.com/googleapis/release-please).

### Issues

We :heart: feedback! Submitting an issue (i.e. feature, bug) is the best way to document things your experience with this project. For example, if there's a feature missing or there's behavior that doesn't match your expectations then we strongly encourage you to submit an issue. That way, contributors can track them and have interested folks (like you) by notified if/when they're resolved.

#### Create a New Issue

Use the Issues feature in GitHub to document bugs and/or features related to this project. Please ensure to apply any/all associated metadata (such as labels) in order to classify them appropriately. Also, please provide as much contextual information as you can, especially when documenting bugs. Templates are provided in this project to guide the authoring process.

#### Resolve an Issue

Issues will be triaged and modified (if necessary) by the [CODEOWNERS](CODEOWNERS) for this project. It is important to associate pull requests with issues by referencing their issue ID in the commit message. That way, issues will be able to document changes and/or fixes. This will assist visitors when reading through issue lists.

### Commit Your Change(s) through Pull Requests

This project employs [branch protection](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/defining-the-mergeability-of-pull-requests/managing-a-branch-protection-rule); the `main` branch is protected. Therefore, your changes MUST be committed to a branch and submitted as a pull request. Also, this project requires the use of [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for all commit messages. Using Conventional Commits enables this project to autogenerate its [CHANGELOG.md](../CHANGELOG.md) and release notes.

### Your Pull Request is Merged! Now What?

Congratulations! :tada: And thank you very much for your contribution to this project!

Once your pull request is merged, our build and test workflow will execute once again to validate changes. Afterward, your changes will be committed to the `main` branch.

### Releasing a New Version

To release a new version, create a tag at the commit that you want to release and push it to the repository. You will require appropriate permissions to push tags to the repository. 

The tag should be in the format `v[major].[minor].[patch]` (e.g. `v1.0.0`). Once the tag is created, the release will be automatically created by the GitHub workflow.

#### What version number do I use?

The version number should be determined by the type of changes that are being released.
- **Major**: Breaking changes. Ideally these should be avoided, if you feel they are necessary please discuss them with the maintainers as part of your pull request.
- **Minor**: New features that are backwards compatible.
- **Patch**: Bug fixes.

## Development Guides

### Enums

There are fields in some types that are represented by enums, for example `FilterType`. If you add a new enum value you need to use `enumer` to update the string representation that gets used in API calls.

- To install enumer run: `go install github.com/dmarkham/enumer@latest`
- Then add it to your path
- To update the string representations, run enumer from the folder where the enum file is located, for example `enumer -type=FilterType -json -output filter_type_string.go`
