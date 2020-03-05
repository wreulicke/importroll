# importroll

importroll is a control tool for import.

## Install

TBD

## Usage

```bash
$ cat 
"*/controller*":
  deny:
    - "*/repository*"
"*/repository*":
  deny:
    - "*/service*"
$ importroll -rule=./importroll.yaml ./...
```