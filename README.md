# GOLAF

***GO** **L**ocal **A**lignment **F**uzzy finder*

Fuzzy word/text finder using [Smith-Waterman algorithm](https://en.wikipedia.org/wiki/Smith%E2%80%93Waterman_algorithm) for finding matches.

![TEST](https://github.com/gwirn/golaf/actions/workflows/go.yml/badge.svg)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](/LICENSE.md)
<a title="Code Size" target="_blank" href="https://github.com/gwirn/golaf"><img src="https://img.shields.io/github/languages/code-size/gwirn/golaf"></a>


https://github.com/gwirn/golaf/assets/71886945/6def3831-e0d2-4ae2-9c5e-49b57af5c4d3

## Setup

Make sure go is [installed](https://go.dev/doc/install)

From within the base directory of this repository run following commands

```sh
git clone https://github.com/gwirn/golaf.git
cd golaf/src
go build -o golaf
```

After running the commands above, on unix systems you can either do `mv golaf /usr/bin` or run `echo 'alias golaf="/PATH/TO/GOLAF"' >> ~/.bashrc` (or `~/.zshrc` depending on your shell) in order to make **GOLAF** easier accessible.

## Usage

Basic search can be done with `golaf [PATTERN] [FILE | STDIN]` to search in files or in stdin.

The search can also be performed reading from StdIn with e.g. `cat [FILE | STDIN] | golaf [PATTERN]`

In order to fuzzy find files or directories by their name but not their content `-type` has to be set to `n`.

To search trough all directories starting in a given directory use the `-recursive` argument (this can be used for file content and for file/directory search).

### Possible optional argument

```go
SYNOPSIS golaf [-bin] [-color] [-gapp] [-match] [-mmp] [-quality] [-recH] [-recursive] [-type] [pattern] [file ...]
  -bin
        include binary files in the search
  -color string
        color option for highlighting the found results- options: [ red green yellow blue purple cyan white ] (default "green")
  -gapp int
        gap penalty [NEGATIVE] (default -2)
  -match int
        score for a match [POSITIVE] (default 3)
  -mmp int
        mismatch penalty [NEGATIVE] (default -3)
  -quality int
        percentage of the pattern that have to match to be seen as match (default 75)
  -recH
        include hidden files in the search
  -recursive string
        root directory for recursively searching through all files (default ".")
  -type string
        Search type
        c - search in file content
        n - search for files and directories (default "c")
```
