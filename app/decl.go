package main

import "errors"

var (
  errFileNameEmpty = errors.New("app: file name not specified")
)

const (
  loggerMain = ""
  loggerCmd = "cmd "
)
