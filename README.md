# srctotext
A tool written in golang to take source code and turn it in to a single text file to use with generative AI. It searches
throgh a folder recursivly and try to find any file that matches the include patterns. For each file
it writes the contents to the file specified by the --output/-o flag.

## Usage

```
$ srctotext --path src/ --include *.go,*.yaml,Makefile,*.md -o source.txt
```

This will result in a file that looks something like this:

```
# FILE: /my-full/path/to/project/src/test.go

<contents of test.go here>

# FILE: /my-full/path/to/project/src/test.yaml

<content of test.yaml>
```

So that each files contents will be in this single file with a header that starts with
`# FILE:` and then the absolute file name including all parents ditectories. Then a
blank line and then full file contents. After the full file contents ther will be an
empty line before the next file header begins.

Requirements:
* If the path is not a folder then log an error and exit
* If no files are found then log an error and exit
* IMPORTANT: Ignore all files that seem to be binary. Has to be checked.

## Options

| Option             | Short | Description                               |
|--------------------|-------|-------------------------------------------|
| `--path`           | `-p`  | Specify the root folder to search.        |
| `--include`        | `-i`  | File patterns to include (comma-separated).|
| `--output`         | `-o`  | Output file where content will be written.|

