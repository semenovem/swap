package control

import (
  "net/http"
  "sync"
)

type Api struct {
  port   int
  server *http.Server
  routes map[string]*Route
  mx     sync.Mutex
}

func New() Api {
  a := Api{
    routes: make(map[string]*Route),
  }

  handlerHelp := func(_ HandlerProps) (*Response, error) {
    resp := &Response{
      Ok:    true,
      Reply: a.routerHelpReply(),
    }

    return resp, nil
  }

  a.Add("/help", handlerHelp)

  return a
}

func (a *Api) Ready() bool {
  return a.server != nil
}
