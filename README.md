# gh-dev-branch

[![Build](https://github.com/rm3l/gh-dev-branch/actions/workflows/build.yml/badge.svg)](https://github.com/rm3l/gh-dev-branch/actions/workflows/build.yml)
[![Lint](https://github.com/rm3l/gh-dev-branch/actions/workflows/lint.yml/badge.svg)](https://github.com/rm3l/gh-dev-branch/actions/workflows/lint.yml)
[![CodeQL Analysis](https://github.com/rm3l/gh-dev-branch/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/rm3l/gh-dev-branch/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rm3l/gh-dev-branch)](https://goreportcard.com/report/github.com/rm3l/gh-dev-branch)

> GitHub CLI extension to generate meaningful branch names off of issues.

This is a simple extension that helps generate meaningful names for branches as soon as you want to start working on a given issue.

Think of it as the CLI counterpart for the beta feature of development branches that can be created in the GitHub UI,
as documented here: https://docs.github.com/en/issues/tracking-your-work-with-issues/creating-a-branch-for-an-issue 

## Installation

- Install the `gh` CLI. See [https://github.com/cli/cli#installation](https://github.com/cli/cli#installation) for further details.
- Install the extension:

```bash
gh extension install rm3l/gh-dev-branch
```

## Usage

```
❯ gh dev-branch -h

Usage: gh dev-branch <issue> [options]
Options:
  -max-length int
        max length of generated branch name (default -1)
  -repo string
        repository to use for finding the issue
```

### Examples

```shell
# Using the current directory. Base repo is determined from the current Git repository
❯ gh dev-branch 294
294-could-not-find-unspecified-when-importing-transitive-dependencies

# or you can use the result in other commands, like to create and switch to that new branch
❯ git switch -c `gh dev-branch 294`     
Switched to a new branch '294-could-not-find-unspecified-when-importing-transitive-dependencies'
```

Or

```shell
# Using an issue from any other GitHub repo
❯ gh dev-branch 9792 -repo openwrt/openwrt
9792-feature-request-deterministic-configuration-export
```

## Working with the source code

1. Clone the repository:

```
git clone https://github.com/rm3l/gh-dev-branch
cd gh-dev-branch
```

2. Download and install [Go](https://go.dev/doc/install) to build the project.
   Or if you are already using the [asdf](https://asdf-vm.com/) version manager, you can just run `asdf install` to install all the necessary tools (declared in the [.tool-versions](.tool-versions) file).

3. Install the local version of this extension; `gh` symlinks to your local source code directory.

```bash
# Install the local version
gh extension install .
```

4. At this point, you can start using it:

```bash
gh dev-branch <my-issue-number> [-repo <my_org/my_repo>]
```

5. To see changes in the code as you develop, simply build and use the extension

```bash
go build && gh dev-branch <my-issue-number> [-repo <my_org/my_repo>]
```

## Contribution Guidelines

Contributions and issue reporting are more than welcome. So to help out, do feel free to fork this repo and open up a pull request.
I'll review and merge your changes as quickly as possible.

You can use [GitHub issues](https://github.com/rm3l/gh-dev-branch/issues) to report bugs.
However, please make sure your description is clear enough and has sufficient instructions to be able to reproduce the issue.

## Developed by

* Armel Soro
    * [keybase.io/rm3l](https://keybase.io/rm3l)
    * [rm3l.org](https://rm3l.org) - &lt;armel+gh-dev-branch@rm3l.org&gt; - [@rm3l](https://twitter.com/rm3l)
    * [paypal.me/rm3l](https://paypal.me/rm3l)
    * [coinbase.com/rm3l](https://www.coinbase.com/rm3l)

## License

    The MIT License (MIT)

    Copyright (c) 2022-2023 Armel Soro

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
