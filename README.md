# srctotext

[![goreleaser](https://github.com/PerArneng/srctotext/actions/workflows/release.yaml/badge.svg)](https://github.com/PerArneng/srctotext/actions/workflows/release.yaml)

`srctotext` is a CLI tool written in Go that aggregates
source files into a single text file. It's useful for
generating datasets for generative AI applications.




## Features

- Recursive directory search
- Filter files by patterns
- Handles non-binary text files
- Outputs to a specified file

## Installation

```shell
$ brew install perarneng/tap/srctotext
```

## Usage

Execute the command with required flags:

```shell
$ srctotext --path src/ --include "*.go,*.yaml,Makefile,*.md" -o source.txt
```

## Output Example

The output file will include contents like:

```
# FILE: src/test.go

<contents of test.go>

# FILE: src/test.yaml

<contents of test.yaml>
```

## Requirements

- Specified path must be a directory
- At least one matching file must be found
- Binary files are ignored

## Options

| Option     | Short | Description                          |
|------------|-------|--------------------------------------|
| `--path`   | `-p`  | Root folder to search.               |
| `--include`| `-i`  | Patterns to include (comma-separated)|
| `--output` | `-o`  | Output file.                         |

## Build and Run

Using Makefile commands:

```
make run
```