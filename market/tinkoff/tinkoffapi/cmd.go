package tinkoffapi

import (
  "github.com/semenovem/bsf/cmdapi"
)

func (t *TfApi) setRoutes() {
  t.cmd.RouteHealth("/health")
  t.cmd.RouteReply("/reset", t.apiReset)
  t.cmd.RouteReply("/ready", t.ready)
  t.cmd.RouteReply("/status", t.status)

  t.cmd.RouteReply("/register", t.ApiSandboxRegister)
  t.cmd.RouteReply("/portfolio", t.ApiPortfolio)
  t.cmd.RouteReply("/portfolio/currencies", t.ApiPortfolioCurrencies)
  t.cmd.RouteReply("/market/stocks", t.ApiMarketStocksHttp)
}

func (t *TfApi) apiReset(reply *cmdapi.Reply) error {
  if !t.cfg.Base.TestMode {
    reply.Err("available only in test mode")
    return nil
  }
  reply.Success()
  return nil
}

func (t *TfApi) ready(reply *cmdapi.Reply) error {
  reply.SetOk(t.st == stReady)
  return nil
}

func (t *TfApi) status(reply *cmdapi.Reply) error {
  reply.Success()
  reply.Payload(statusKeyVal[t.st])
  return nil
}
