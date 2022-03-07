package main

import (
  "github.com/semenovem/bsf/cmdapi"
)

func init() {
  api.RouteText("/version", buildVersion)
  api.RouteHealth("/health")
  api.RouteReply("/exit", exit)
  api.RouteReply("/reset", apiReset)
  api.RouteReply("/ready", readyHttp)
}

func apiReset(reply *cmdapi.Reply) error {
  if !cfg.Base.TestMode {
    reply.Err("available only in test mode. Set ENV_TEST_MODE=true")
    return nil
  }

  reply.Success()

  return nil
}

func readyHttp(reply *cmdapi.Reply) error {
  reply.SetOk(app.Ready())
  return nil
}

func exit(reply *cmdapi.Reply) error {
  reply.Success()
  app.Exit()
  return nil
}
