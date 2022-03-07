package tinkoffsync

import (
  "github.com/semenovem/bot/core"
)

func (t *TfSync) Cfg(c *Config) {
  t.log.Infof("%+v", *c)

  failure := false

  if failure {
    t.log.Error("Конфигурация содержит ошибки и не может быть использована")
    return
  }

  t.cfg = c
  t.cmd.SetPort(c.Base.CmdPort)
}

type Config struct {
  Base   core.BaseConfig `yaml:"base"`
  Enable bool            `yaml:"enable"`

  Sandbox bool // Включена песочница
}
