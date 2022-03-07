package main

import (
  "github.com/semenovem/bot/market/tinkoff/tinkoffapi"
  "github.com/semenovem/bot/market/tinkoff/tinkoffsync"
  "github.com/semenovem/bsf/cmdapi"
  "github.com/semenovem/bsf/logger"
  "github.com/semenovem/bsf/mgr"
  "time"
)

var (
  buildVersion = "UNKNOWN"
  cfg       = &appConfig{}
  man       = mgr.New()
  logHolder = logger.New()
  log       = logHolder.GetLog(loggerMain)
  api       = cmdapi.New(man.Ctx, logHolder.GetLog(loggerCmd))
  tfApi     = tinkoffapi.New(man.Ctx, logHolder.GetLog(tinkoffapi.Sys))
  tfSync    = tinkoffsync.New(man.Ctx, logHolder.GetLog(tinkoffsync.Sys), tfApi)
)

func init() {
  man.Log = log
  man.ShutdownTimeoutMs = time.Millisecond * 100

  man.RegisterStarted(func() {
    log.Info("Application started")
  })
}

func main() {
  log.Info("Application starts")
  defer log.Info("Application stopped")

  // http manager api
  if cfg.Base.CmdPort > 0 {
    api.SetPort(cfg.Base.CmdPort)
    man.Task(api.Start)
  } else {
    log.Warn("Management api not started, because no port is specified")
  }

  man.Task(tfApi.Start)
  man.Task(tfSync.Start)


  man.Wait()
}
