package main

import (
  "flag"
  "github.com/caarlos0/env/v6"
  "github.com/pkg/errors"
  "github.com/semenovem/bot/core"
  "github.com/semenovem/bot/market/tinkoff/tinkoffapi"
  "github.com/semenovem/bot/market/tinkoff/tinkoffsync"
  "gopkg.in/yaml.v3"
  "io/ioutil"
  "os"
  "time"
)

// Приоритет: файл > переменные окружения > флаги
type appConfig struct {
  FileConfig string          `env:"ENV_FILE_CONFIG"`
  Base       core.BaseConfig `yaml:"base"`

  // Время ожидания для запуска time.Millisecond
  TimeoutAppLaunch int64 `yaml:"timeout-app-launch"`

  TfApiConfig  string `yaml:"tinkoff-api-config"` // файл конфигурации
  TfNotSandbox bool   `yaml:"not-sandbox"`

  TfSyncConfig string `yaml:"tinkoff-sync-config"` // файл конфигурации
}

func init() {
  if v, ok := os.LookupEnv("ENV_FILE_CONFIG"); ok {
    cfg.FileConfig = v
  }
  flag.StringVar(&cfg.FileConfig, "fileConfig", cfg.FileConfig, "")
  flag.Parse()

  if cfg.FileConfig != "" {
    _ = parseYaml(cfg.FileConfig, &cfg)
  }

  if err := env.Parse(cfg); err != nil {
    log.Error(err)
  }

  // APP
  man.Timeout = time.Millisecond * time.Duration(cfg.TimeoutAppLaunch)
  if cfg.Base.DevMode {
    man.IsCli = cfg.Base.DevMode
    _ = logHolder.SetLevel(loggerMain, "WARN")
  }

  // LOG DEFAULT
  b := cfg.Base
  err := logHolder.SetDef(b.LogLevel, b.LogMode, b.LogDest)
  isErr(err)

  // API
  isErr(logHolder.SetLevel(loggerCmd, cfg.Base.CmdLogLevel))

  // TINKOFF-API
  if cfg.TfApiConfig != "" {
    c := &tinkoffapi.Config{}
    err := parseYaml(cfg.TfApiConfig, c)
    if err == nil {
      tfApi.Cfg(c)
      b = c.Base
      err := logHolder.Set(tinkoffapi.Sys, b.LogLevel, b.LogMode, b.LogDest)
      isErr(err)
    } else {
      log.Warn("Ошибка получения конфигурации tinkoff api: ", err)
    }
  }

  // TINKOFF-SYNC
  if cfg.TfSyncConfig != "" {
    c := &tinkoffsync.Config{}
    err := parseYaml(cfg.TfSyncConfig, c)
    if err == nil {
      tfSync.Cfg(c)
      b = c.Base
      err := logHolder.Set(tinkoffsync.Sys, b.LogLevel, b.LogMode, b.LogDest)
      isErr(err)
    } else {
      log.Warn("Ошибка получения конфигурации tinkoff sync: ", err)
    }
  }
}

func parseYaml(f string, c interface{}) error {
  if f == "" {
    return errFileNameEmpty
  }
  byt, err := ioutil.ReadFile(f)
  if err != nil {
    return errors.Wrap(err, "Ошибка чтения файла")
  } else {
    err = yaml.Unmarshal(byt, c)
    if err != nil {
      return errors.Wrap(err, "Ошибка парсинга yaml")
    }
  }
  return nil
}
