# GOLAF

***GO** **L**ocal **A**lignment **F**uzzy finder*

Fuzzy finder using [Smith-Waterman algorithm](https://en.wikipedia.org/wiki/Smith%E2%80%93Waterman_algorithm) for finding matches. Can be used for words and DNA/Protein sequences.

![TEST](https://github.com/gwirn/golaf/actions/workflows/go.yml/badge.svg)

## Setup

Make sure go is [installed](https://go.dev/doc/install)

From within the base directory run following commands

```
cd src
go build
mv src golaf
```

On unix systems you can either do `mv golaf /usr/bin` or run `echo 'alias golaf="/PATH/TO/GOLAF"' >> ~/.bashrc` (or `~/.zshrc` depending on your shell).

To run it use it with `golaf [PATTERN] [FILE]`.

## Usage

Basic search can be done with `golaf [PATTERN] [FILE]`

The search can also be performed reading from StdIn with e.g. `cat testfile.txt | golaf [PATTERN]`

### Possible optional argument

```
-color string
    true to get colored the output - options: [ red green yellow blue purple cyan white ] (default "green")
-gapp int
    gap penalty [NEGATIVE] (default -2)
-match int
    score for a match [POSITIVE] (default 3)
-mmp int
    missmatch penalty [NEGATIVE] (default -3)
-quality int
    percentage of the pattern that have to macht to be seen as match (default 60)
```
