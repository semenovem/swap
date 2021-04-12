package lg

import (
  "github.com/sirupsen/logrus"
)

type op struct {
  ent        *logrus.Entry
  dev        bool
  timeFormat string
  mode       string
  lev        logrus.Level
  pkgName    string
  cli        bool
}

const (
  ModeJson      = "json"
  ModeText      = "text"
  ModeShortText = "short"

  Trace = logrus.TraceLevel
  Debug = logrus.DebugLevel
  Info  = logrus.InfoLevel
  Warn  = logrus.WarnLevel
  Error = logrus.ErrorLevel
  Panic = logrus.PanicLevel
)

func With(ent *logrus.Entry) *op {
  return &op{
    ent:        ent,
    timeFormat: defaultTimeFormat,
    mode:       defaultMode,
    lev:        defaultLev,
    dev:        defaultDev,
    cli:        defaultCli,
  }
}

func New() *op {
  return With(logrus.NewEntry(logrus.New()))
}

// Добавляет в лог имя пакета
func (o *op) Pkg(n string) *op {
  o.pkgName = n
  return o
}

// Установить уровень логгирования
func (o *op) Lev(s string) *op {
  if l, err := logrus.ParseLevel(s); err == nil {
    o.lev = l
  }
  return o
}

func (o *op) Apl() *logrus.Entry {
  o.ent.Logger.SetLevel(o.lev)
  o.ent.Logger.SetFormatter(o.formatter())

  if o.pkgName != "" {
    o.ent.Data[logAdditionalKey] = o.pkgName
  }

  return o.ent
}

func (o *op) Dev(on bool) *op {
  o.dev = on
  return o
}

func (o *op) Mod(m string) *op {
  o.mode = m
  return o
}
