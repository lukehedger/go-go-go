# Go, Go, Go!

> Learning Go

## Install

Build the `go-go-go` command into an executable binary and install binary to the
Go [workspace](https://golang.org/doc/code.html#Workspaces)'s bin directory as
`go-go-go`:
```bash
cd $GOPATH/src/github.com/lukehedger/go-go-go
go install
```

## Run

If you have added [$GOPATH/bin to your PATH](https://golang.org/doc/code.html#GOPATH),
just type the binary name:
```bash
go-go-go
# => Go, Go, Go!
```

Otherwise, run `$GOPATH/bin/go-go-go`

## Code Organisation

- Go programmers typically keep all their Go code in a single workspace.
- A workspace contains many version control repositories (managed by Git, for
  example).
- Each repository contains one or more packages.
- Each package consists of one or more Go source files in a single directory.
- The path to a package's directory determines its import path.

Note that this differs from other programming environments in which every
project has a separate workspace and workspaces are closely tied to version
control repositories.
