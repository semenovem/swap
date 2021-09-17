package api

import (
  "fmt"
  "github.com/semenovem/bsf/cmdapi"
  "strings"
)

func (t *TfApi) ApiPortfolio(reply *cmdapi.Reply) error {
  var regResp PortfolioReply

  err := t.entryRequest(t.declApi.portfolio.portfolio, nil, &regResp)
  if err != nil {
    return err
  }

  if strings.ToUpper(regResp.Status) != "OK" {
    err := fmt.Errorf("register failed, trackingId: '%s'", regResp.TrackingId)
    t.log.Error(err)
    reply.Err(err)
    return err
  }

  fmt.Printf("\n>>>>>> %+v\n\n", regResp)
  fmt.Println("Register succeed")
  reply.Success()

  return nil
}

func (t *TfApi) ApiPortfolioCurrencies(reply *cmdapi.Reply) error {
  var regResp PortfolioCurrenciesReply

  err := t.entryRequest(t.declApi.portfolio.currencies, nil, &regResp)
  if err != nil {
    return err
  }

  if strings.ToUpper(regResp.Status) != "OK" {
    err := fmt.Errorf("register failed, trackingId: '%s'", regResp.TrackingId)
    t.log.Error(err)
    reply.Err(err)
    return err
  }

  fmt.Printf("\n>>>>>> %+v\n\n", regResp)
  fmt.Println("Register succeed")
  reply.Success()

  return nil
}
