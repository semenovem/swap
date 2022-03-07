package tinkoffsync

import (
  "context"
  "github.com/semenovem/bot/market/tinkoff/tinkoffapi"
  "github.com/semenovem/bsf/cmdapi"
  "github.com/sirupsen/logrus"
)

func New(ctx context.Context, l *logrus.Entry, tfApi *tinkoffapi.TfApi) *TfSync {
  cmd := cmdapi.New(ctx, l.WithField("", "cmd"))

  o := &TfSync{
    ctx:   ctx,
    log:   l,
    cmd:   cmd,
    tfApi: tfApi,
  }

  return o
}

func (t *TfSync) Start() error {
  t.log.Info("Start")

  _ = t.cmd.Start()

  return nil
}
