# logtime

This tool allows analysis where is time wasted in the logs.

It will create a report you can analyze further in your favorite
text editor/spreadsheet program.

## Installation

```
go get -u github.com/milanaleksic/logtime/cmd/logtime
```

## Usage

```
$ logtime --help
Usage of logtime:
  -input-file string
        which file to process (default - stdin)
  -log-time string
        pattern that should match beginning of all log lines (default "2006-01-02 15:04:05")
```

> Note: read about the [time format layout used in Go language](https://golang.org/pkg/time/#Parse), 
> that's the format expected by `log-time` argument

## Example

Example with progress provided via `pv` utility

```bash
pv -i 0.1 clean.txt | \
    logtime \
    > output.tsv
```
