package main

func isErr(err error) {
  if err != nil {
    log.Warn(err)
  }
}
