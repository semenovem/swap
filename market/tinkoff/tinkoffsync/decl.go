package tinkoffsync

import (
  "context"
  "errors"
  "github.com/semenovem/bot/market/tinkoff/tinkoffapi"
  "github.com/semenovem/bsf/cmdapi"
  "github.com/sirupsen/logrus"
  "sync"
)

var (
  ErrNoCfg         = errors.New("tinkoff-sync: no configuration")
)

const (
  Sys    = "tfsync"
  System = "tinkoff.sync"
)

type status uint

const (
  stUnkn status = iota
  stReady
  stDisconnect
  stErr
)

var statusKeyVal = map[status]string{
  stUnkn:       "unknown",
  stReady:      "ready",
  stDisconnect: "disconnect",
  stErr:        "error",
}

var statusValKey = map[string]status{
  statusKeyVal[stUnkn]:       stUnkn,
  statusKeyVal[stReady]:      stReady,
  statusKeyVal[stDisconnect]: stDisconnect,
  statusKeyVal[stErr]:        stErr,
}

type TfSync struct {
  ctx     context.Context
  log     *logrus.Entry
  cfg     *Config
  enabled bool
  cmd     *cmdapi.Cmd
  mx      sync.Mutex
  st    status
  tfApi *tinkoffapi.TfApi
}
