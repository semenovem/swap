package api

import (
  "fmt"
  "github.com/semenovem/bsf/cmdapi"
  "strings"
)

func (t *TfApi) ApiMarketStocks(reply *cmdapi.Reply) error {
  var regResp MarketStocksReply

  err := t.entryRequest(t.declApi.market.stocks, nil, &regResp)
  if err != nil {
    return err
  }

  if strings.ToUpper(regResp.Status) != "OK" {
    err := fmt.Errorf("register failed, trackingId: '%s'", regResp.TrackingId)
    t.log.Error(err)
    reply.Err(err)
    return err
  }

  fmt.Printf("\n>>>>>> Total = %+v\n\n", regResp.Payload.Total)
  fmt.Printf("\n>>>>>> Instruments = %+v\n\n", regResp.Payload.Instruments[0])
  fmt.Println("Register succeed")
  reply.Success()

  return nil
}
