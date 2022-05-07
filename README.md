# gh-dev-branch

[![Build](https://github.com/rm3l/gh-dev-branch/actions/workflows/build.yml/badge.svg)](https://github.com/rm3l/gh-dev-branch/actions/workflows/build.yml)
[![Lint](https://github.com/rm3l/gh-dev-branch/actions/workflows/lint.yml/badge.svg)](https://github.com/rm3l/gh-dev-branch/actions/workflows/lint.yml)
[![CodeQL Analysis](https://github.com/rm3l/gh-dev-branch/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/rm3l/gh-dev-branch/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rm3l/gh-dev-branch)](https://goreportcard.com/report/github.com/rm3l/gh-dev-branch)

> GitHub CLI extension to generate branch names off of issues.

This is a simple extension that helps generate names for branches as soon as you want to start working on a given issue.

Think of it as the CLI counterpart for the beta feature of development branches that can be created in the GitHub UI,
as documented here: https://docs.github.com/en/issues/tracking-your-work-with-issues/creating-a-branch-for-an-issue 

## Installation

- Install the `gh` CLI. See [https://github.com/cli/cli#installation](https://github.com/cli/cli#installation) for further details.
- Install the extension:

```bash
gh extension install rm3l/gh-dev-branch
```

## Usage

```bash
❯ gh gh-dev-branch -h

Usage: gh dev-branch <issue> [options]
Options: 
  -repo string
        repository to use for finding the issue
```

## Working with the source code

- If not done already, download and install [Go](https://go.dev/doc/install) to build the project. This project has been built with Go 1.18.
- Clone the repository and install the local version.

```bash
cd dev-branch

# Install the local version
gh extension install .

# At this point, you can start using it
gh dev-branch <my-issue-number>

# To see changes in the code as you develop,
# simply build and use the extension:
go build && gh dev-branch <my-issue-number>
```

## Contribution Guidelines

Contributions and issue reporting are more than welcome. So to help out, do feel free to fork this repo and open up a pull request.
I'll review and merge your changes as quickly as possible.

You can use [GitHub issues](https://github.com/rm3l/dev-branch/issues) to report bugs.
However, please make sure your description is clear enough and has sufficient instructions to be able to reproduce the issue.

## Developed by

* Armel Soro
    * [keybase.io/rm3l](https://keybase.io/rm3l)
    * [rm3l.org](https://rm3l.org) - &lt;armel+dev-branch@rm3l.org&gt; - [@rm3l](https://twitter.com/rm3l)
    * [paypal.me/rm3l](https://paypal.me/rm3l)
    * [coinbase.com/rm3l](https://www.coinbase.com/rm3l)

## License

    The MIT License (MIT)

    Copyright (c) 2022 Armel Soro

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
