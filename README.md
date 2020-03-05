# importroll

importroll is a control tool for import.

## Install

Get binary from [GitHub Releases](https://github.com/wreulicke/importroll/releases/latest)

## Usage

```bash
$ cat importroll.yaml
"*/controller*":
  deny:
    - "*/repository*"
"*/repository*":
  deny:
    - "*/service*"
$ importroll -rule=./importroll.yaml ./...
```

## License

MIT