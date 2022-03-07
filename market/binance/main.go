package main

//import (
//  "bot/pkg/app/launch"
//  "bot/pkg/lg"
//)
//
//var (
//  log = lg.New().Pkg("main").Apl()
//)
//
//func main() {
//  log.Info("Приложение запускается")
//  defer log.Info("Приложение остановлено")
//
//  ins := launch.New(launch.Cfg{
//    Timeout: cfg.WaitAppLaunch,
//    StartedFn: func() {
//      log.Info("Приложение запущено")
//    },
//    Log: log,
//  })
//
//  ins.Task(ctrlApi.Start)
//
//  ins.Wait()
//}
