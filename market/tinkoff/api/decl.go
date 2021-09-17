package api

import (
  "context"
  "errors"
  "github.com/semenovem/bsf/cmdapi"
  "github.com/sirupsen/logrus"
  "net/http"
  "sync"
  "time"
)

var (
  ErrNoCfg         = errors.New("tinkoff-api: no configuration")
  ErrNoToken       = errors.New("tinkoff-api: no file token")
  ErrReadFileToken = errors.New("tinkoff-api: error reading token file")
)

const (
  Sys    = "tfapi"
  System = "tinkoff.api"
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

type TfApi struct {
  ctx         context.Context
  log         *logrus.Entry
  cfg         *Config
  enabled     bool
  api         *cmdapi.Cmd
  mx          sync.Mutex
  st          status
  token       string
  timeoutHttp time.Duration // Время ожидания http запроса
  declApi     *descApi
  httpClient  *http.Client
}
