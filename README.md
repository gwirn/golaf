# GOLAF

***GO** **L**ocal **A**lignment **F**uzzy finder*

Fuzzy word/text finder using [Smith-Waterman algorithm](https://en.wikipedia.org/wiki/Smith%E2%80%93Waterman_algorithm) for finding matches.

![alt text](https://github.com/gwirn/golaf/blob/master/screenshot/golaf.png?raw=true)

![TEST](https://github.com/gwirn/golaf/actions/workflows/go.yml/badge.svg)
<a title="Code Size" target="_blank" href="https://github.com/gwirn/golaf"><img src="https://img.shields.io/github/languages/code-size/gwirn/golaf"></a>

## Setup

Make sure go is [installed](https://go.dev/doc/install)

From within the base directory of this repository run following commands

```
git clone https://github.com/gwirn/golaf.git
cd golaf/src
go build
mv src golaf
```

After running the commands above, on unix systems you can either do `mv golaf /usr/bin` or run `echo 'alias golaf="/PATH/TO/GOLAF"' >> ~/.bashrc` (or `~/.zshrc` depending on your shell) in order to make **GOLAF** easier accessible.

## Usage

Basic search can be done with `golaf [PATTERN] [FILE | STDIN]`

The search can also be performed reading from StdIn with e.g. `cat [FILE | STDIN] | golaf [PATTERN]`

### Possible optional argument

```
  -color string
    	color option for highlighting the found results- options: [ red green yellow blue purple cyan white ] (default "green")
  -gapp int
    	gap penalty [NEGATIVE] (default -2)
  -match int
    	score for a match [POSITIVE] (default 3)
  -mmp int
    	missmatch penalty [NEGATIVE] (default -3)
  -quality int
    	percentage of the pattern that have to macht to be seen as match (default 75)
```
