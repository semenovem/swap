package tinkoffapi

import (
  "fmt"
  "github.com/semenovem/bsf/cmdapi"
  "strings"
)

func (t *TfApi) ApiSandboxRegister(reply *cmdapi.Reply) error {
  fmt.Println(">>>>>>>>>> create account", t.token)

  var regResp PortfolioReply

  err := t.entryRequest(t.declApi.sandbox.register, nil, &regResp)
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
