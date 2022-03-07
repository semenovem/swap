package main

//import (
//  "bot/pkg/apicmd"
//  "bot/pkg/app/build"
//)
//
//var ctrlApi = apicmd.New()
//
//func init() {
//  ctrlApi.Add("/reset", reset)
//  ctrlApi.Add("/ready", readyHttpHandler)
//
//  ctrlApi.AddSimple("/version", build.Ver)
//  ctrlApi.AddSimple("/app", build.ToString)
//  ctrlApi.AddSimple("/health", func() string {
//    return ""
//  })
//}
//
//func reset(_ apicmd.HandlerProps) (*apicmd.Response, error) {
//  log.Info("Сброс приложения")
//
//  resp := &apicmd.Response{
//    Ok: true,
//  }
//
//  return resp, nil
//}
//
//func readyHttpHandler(_ apicmd.HandlerProps) (*apicmd.Response, error) {
//  resp := &apicmd.Response{
//    Ok: true,
//  }
//
//  return resp, nil
//}
