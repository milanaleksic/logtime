# logtime

This tool allows analysis where is time wasted in the logs.

It will create a report you can analyze further in your favorite
text editor/spreadsheet program.

## Installation

```
go get -u github.com/milanaleksic/logtime
```

## Usage

```
```

## Example

Example with progress provided via `pv` utility

```bash
pv -i 0.1 biglogfile.txt | gclog-cleaner \
  --exclusions '\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.*\[GC' \
  --exclusions "Creating an interceptor chain" \
  --exclusions 'request:84' \
  --exclusions 'DEBUG Sdk' \ 
  > clean.txt
```
