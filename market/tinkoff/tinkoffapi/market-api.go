package tinkoffapi

import (
  "errors"
  "fmt"
  "github.com/semenovem/bot/core"
  "github.com/semenovem/bsf/cmdapi"
  "strings"
)

func (t *TfApi) ApiMarketStocksHttp(reply *cmdapi.Reply) error {
  _, err := t.ApiMarketStocks()
  if err != nil {
    reply.Err(err)
  } else {
    reply.Success()
  }
  return err
}

func (t *TfApi) ApiMarketStocks() (*core.MarketStocksReply, error) {
  var regResp MarketStocksReply

  err := t.entryRequest(t.declApi.market.stocks, nil, &regResp)
  if err != nil {
    return nil, err
  }

  if strings.ToUpper(regResp.Status) != "OK" {
    err := fmt.Errorf("register failed, trackingId: '%s'", regResp.TrackingId)
    t.log.Error(err)
    return nil, err
  }

  if regResp.Payload.Total < 0 {
    return nil, errors.New("поле Payload.total не может быть меньше 0")
  }

  if regResp.Payload.Total != len(regResp.Payload.Instruments) {
    return nil, errors.New("значение в Payload.total не равно кол-ву элементов в массиве Payload.Instruments")
  }

  fmt.Printf("\n>>>>>> Total = %+v\n\n", regResp.Payload.Total)
  fmt.Printf("\n>>>>>> Instruments = %+v\n\n", regResp.Payload.Instruments[0])
  fmt.Println("Register succeed")

  v, err := convMarketStocksReply(&regResp)

  return v, nil
}

func convMarketStocksReply(v *MarketStocksReply) (*core.MarketStocksReply, error) {
  v2 := &core.MarketStocksReply{
    Total:       v.Payload.Total,
    Instruments: make([]*core.Instrument, len(v.Payload.Instruments)),
  }



  // Перенести данные в новую структуру
  //v2.Instruments

  return v2, nil
}
