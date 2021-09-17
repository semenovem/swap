package tinkoffapi

import (
  "github.com/semenovem/bot/common"
  "time"
)

func (t *TfApi) Cfg(c *Config) {
  t.log.Infof("%+v", *c)

  failure := false
  if c.TimeoutHttp == 0 {
    t.log.Warn("Не указано значение TimeoutHttp")
    failure = true
  }

  if failure {
    t.log.Error("Конфигурация содержит ошибки и не может быть использована")
    return
  }

  t.cfg = c
  t.cmd.SetPort(c.Base.CmdPort)

  t.timeoutHttp = time.Duration(c.TimeoutHttp) * time.Millisecond
}

type Config struct {
  Base      common.BaseConfig `yaml:"base"`
  Enable    bool              `yaml:"enable"`
  UrlEntry  string            `yaml:"url-entry"`
  TokenFile string            `yaml:"token-file"` // токен

  Sandbox bool // Включена песочница

  TimeoutHttp uint `yaml:"timeout-http"` // Время ожидания http запроса time.Millisecond
}
