package api

import (
  "context"
  "github.com/pkg/errors"
  "github.com/semenovem/bsf/cmdapi"
  "github.com/sirupsen/logrus"
  "io/ioutil"
  "net/http"
  "strings"
)

func New(ctx context.Context, l *logrus.Entry) *TfApi {
  api := cmdapi.New(ctx, l.WithField("", "api"))

  o := &TfApi{
    ctx:     ctx,
    log:     l,
    api:     api,
    declApi: declEntry,
  }

  o.setRoutes()

  return o
}

func (t *TfApi) Start() error {
  t.mx.Lock()
  defer t.mx.Unlock()

  if t.cfg == nil {
    return ErrNoCfg
  }

  if !t.cfg.Enable {
    t.log.Warn("Система 'tinkoff.api' выключена [enable: false]")
    return nil
  }

  t.log.Info("Запуск")

  _ = t.api.Start()

  tkn, err := t.readToken()
  if err != nil {
    return errors.Wrapf(err, "Ошибка старта системы [%s]", System)
  }
  t.token = tkn

  t.httpClient = &http.Client{
    Timeout: t.timeoutHttp,
  }

  return nil
}

func (t *TfApi) readToken() (string, error) {
  if t.cfg.TokenFile == "" {
    return "", ErrNoToken
  }
  byt, err := ioutil.ReadFile(t.cfg.TokenFile)
  if err != nil {
    return "", errors.Wrap(err, ErrReadFileToken.Error())
  }
  return strings.TrimSpace(string(byt)), nil
}
