package lg

import (
  "github.com/sirupsen/logrus"
  "time"
)

var (
  defaultTimeFormat = time.RFC3339
  defaultDev        = false
  defaultMode       = ModeJson
  defaultLev        = logrus.InfoLevel
  defaultCli        = false
)

// флаг разработки
func SetDev(on bool) {
  defaultDev = on
}

func GetDev() bool {
  return defaultDev
}

// Уровень логгирования
func SetDefLev(l logrus.Level) {
  defaultLev = l
}

func GetDefLev() logrus.Level {
  return defaultLev
}

// Режим логов
func SetDefMod(m string) {
  if !isMod(m) {
    panic("The value is not a valid mode")
  }
  defaultMode = m
}

func GetDefMod() string {
  return defaultMode
}

// Режим вывода логов в консоль с подсветкой
func SetDefCLI(on bool) {
  defaultCli = on
}

func GetDefCLI() bool {
  return defaultCli
}
