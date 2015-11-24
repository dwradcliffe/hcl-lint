HCL-LINT
---

## Usage:

```sh
$ go get github.com/dwradcliffe/hcl-lint
$ go run /go/src/github.com/dwradcliffe/hcl-lint/lint.go folder_path config_file ...
```

Lint's an arbitrary number of input arguments. If a folder path is specified,
it will check that folder for *.tf files and run the linter on them as well.
