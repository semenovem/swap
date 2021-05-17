package control

import (
  "fmt"
  "net/http"
  "strconv"
)

func (a *Api) SetPort(port int) {
  a.port = port
}

func (a *Api) IsStarted() bool {
  return a.server != nil
}

func (a *Api) Start() error {
  a.mx.Lock()
  defer a.mx.Unlock()

  if a.server != nil {
    Log.Warn("Повторный вызов старта сервера")
    return nil
  }

  if a.port < 1 {
    err := fmt.Errorf("invalid value 'port': %d", a.port)
    Log.Error(err)
    return err
  }

  server := &http.Server{
    Addr:    ":" + strconv.Itoa(a.port),
    Handler: http.HandlerFunc(a.router),
  }

  a.server = server

  go func() {
    defer a.Stop()

    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      Log.Error("API server error: ", err)
    }
  }()

  Log.Infof("API server started on port %d", a.port)
  return nil
}

func (a *Api) Stop() {
  a.mx.Lock()
  defer a.mx.Unlock()

  if a.server != nil {
    err := a.server.Close()
    a.server = nil
    if err != nil {
      Log.Error("API server stop error")
    }
    Log.Info("API server stopped")
  }
}
