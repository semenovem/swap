package api

import (
  "github.com/semenovem/bsf/cmdapi"
)

func (t *TfApi) setRoutes() {
  t.api.RouteHealth("/health")
  t.api.RouteReply("/reset", t.apiReset)
  t.api.RouteReply("/ready", t.ready)
  t.api.RouteReply("/status", t.status)

  t.api.RouteReply("/register", t.ApiSandboxRegister)
  t.api.RouteReply("/portfolio", t.ApiPortfolio)
  t.api.RouteReply("/portfolio/currencies", t.ApiPortfolioCurrencies)
  t.api.RouteReply("/market/stocks", t.ApiMarketStocks)
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
