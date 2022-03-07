package main

//import (
//  "bot/pkg/apicmd"
//  "bot/pkg/lg"
//  "flag"
//  "github.com/caarlos0/env/v6"
//)
//
//type appConfig struct {
//  DevMode  bool `env:"ENV_DEV_MODE"`
//  TestMode bool `env:"ENV_TEST_MODE"`
//
//  LogLevel string `env:"ENV_LOG_LEVEL"`
//  LogMode  string `env:"ENV_LOG_MODE"`
//  LogCli   bool   `env:"ENV_LOG_CLI"`
//
//  // Управляющий порт
//  ControlApiPort int `env:"ENV_CONTROL_API_PORT"`
//
//  // Время ожидания для подсистем прилоложения time.Millisecond
//  WaitAppLaunch int64 `env:"ENV_WAIT_APP_LAUNCH" envDefault:"10000"`
//
//  // Пара ключей для доступа к api биржи
//  // если задан файл - читаем из него. В противном случае берем из переменной окружения
//  ApiKeyPub     string `env:"ENV_API_KEY_PUB"`
//  ApiKeyPrv     string `env:"ENV_API_KEY_PRV"`
//  ApiKeyPubFile string `env:"ENV_API_KEY_PUB_FILE"`
//  ApiKeyPrvFile string `env:"ENV_API_KEY_PRV_FILE"`
//
//  // hsm
//  HsmPathHostLib string `env:"ENV_HSM_PATH_HOST_LIB"`
//  HsmSlot        string `env:"ENV_HSM_SLOT"`
//  HsmPin         string `env:"ENV_HSM_PIN"`
//  HsmPubLab      string `env:"ENV_HSM_PUB_LAB"`
//  HsmPrvLab      string `env:"ENV_HSM_PRV_LAB"`
//
//  // список api
//  Endpoints      string `env:"ENV_ENDPOINTS"`
//
//
//
//}
//
//var cfg = &appConfig{}
//
//func init() {
//  if err := env.Parse(cfg); err != nil {
//    log.Error(err)
//  }
//
//  flag.BoolVar(&cfg.DevMode, "devMode", cfg.DevMode, "")
//  flag.BoolVar(&cfg.TestMode, "testMode", cfg.TestMode, "")
//  flag.StringVar(&cfg.LogMode, "logMode", cfg.LogMode, "")
//  flag.BoolVar(&cfg.LogCli, "logCli", cfg.LogCli, "")
//
//  flag.IntVar(&cfg.ControlApiPort, "controlApiPort", cfg.ControlApiPort, "")
//  flag.StringVar(&cfg.LogLevel, "logLevel", cfg.LogLevel, "")
//
//  flag.Int64Var(&cfg.WaitAppLaunch, "waitTimeAppLaunch", cfg.WaitAppLaunch, "")
//
//  flag.StringVar(&cfg.ApiKeyPub, "apiKeyPub", cfg.ApiKeyPub, "")
//  flag.StringVar(&cfg.ApiKeyPrv, "apiKeyPrv", cfg.ApiKeyPrv, "")
//  flag.StringVar(&cfg.ApiKeyPubFile, "apiKeyPubFile", cfg.ApiKeyPubFile, "")
//  flag.StringVar(&cfg.ApiKeyPrvFile, "apiKeyPrvFile", cfg.ApiKeyPrvFile, "")
//
//  flag.StringVar(&cfg.HsmPathHostLib, "hsmPathHostLib", cfg.HsmPathHostLib, "")
//  flag.StringVar(&cfg.HsmSlot, "hsmSlot", cfg.HsmSlot, "")
//  flag.StringVar(&cfg.HsmPin, "hsmPin", cfg.HsmPin, "")
//  flag.StringVar(&cfg.HsmPubLab, "hsmPubLab", cfg.HsmPubLab, "")
//  flag.StringVar(&cfg.HsmPrvLab, "hsmPrvLab", cfg.HsmPrvLab, "")
//  flag.Parse()
//
//  lg.SetDefLev(lg.ParseLev(cfg.LogLevel, lg.Trace))
//  lg.SetDefMod(lg.ParseMod(cfg.LogMode, lg.ModeJson))
//  lg.SetDefCLI(cfg.LogCli)
//
//  lg.With(log).Apl()
//  lg.With(apicmd.Log).Pkg("api ").Apl()
//
//  log.Debugf("Application configuration: %+v", *cfg)
//
//  ctrlApi.SetPort(cfg.ControlApiPort)
//}
