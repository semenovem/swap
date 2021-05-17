package control

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strings"
)

func (a *Api) routerHelpReply() string {
  keys := make([]string, 0)
  for it, _ := range a.routes {
    keys = append(keys, it)
  }
  return fmt.Sprintf("Available URL: [%s]", strings.Join(keys, ", "))
}

func (a *Api) Add(path string, handler RouteHandler) {
  _, has := a.routes[path]
  if has {
    Log.Panic("URL path already exists")
  }

  if handler == nil {
    Log.Panic("Handler cannot be nil")
  }

  a.routes[path] = &Route{
    Path:    path,
    Handler: handler,
  }
}

func (a *Api) AddOrReplace(path string, handler RouteHandler) {
  a.DelRoute(&path)
  a.Add(path, handler)
}

func (a *Api) DelRoute(path *string) {
  delete(a.routes, *path)
}

func (a *Api) HasRoute(path *string) bool {
  _, has := a.routes[*path]
  return has
}

func (a *Api) AddSimple(route string, fn func() string) {
  a.Add(route, func(_ HandlerProps) (*Response, error) {
    return &Response{
      Ok:    true,
      Reply: fn(),
    }, nil
  })
}

func (a *Api) router(w http.ResponseWriter, r *http.Request) {
  Log.Info("API IN", r.URL.Path)

  path := strings.ToLower(r.URL.Path)
  pathLen := len(path)

  var matchLen int
  var route *Route

  for p, r := range a.routes {
    if strings.HasPrefix(path, p) && len(p) > matchLen {
      if pathLen > len(p) && string(path[len(p)]) != "/" {
        continue
      }
      route = r
      matchLen = len(p)
    }
  }

  if route != nil {
    a.routeServe(w, r)(route.Handler)

    return
  }

  a.NotFound(HandlerProps{w, r}, a.routerHelpReply())
}

func (a *Api) routeServe(w http.ResponseWriter, r *http.Request) func(h RouteHandler) {
  return func(handler RouteHandler) {
    resp, err := handler(HandlerProps{w, r})

    if err != nil {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusInternalServerError)
      body, _ := json.Marshal(&Response{
        ErrMsg: fmt.Sprint(err),
      })
      body = append(body, []byte("\n")...)
      _, err := w.Write(body)
      if err != nil {
        Log.Error(err)
      }

      return
    }

    if resp != nil {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      body, _ := json.Marshal(resp)
      _, err := w.Write(body)
      if err != nil {
        Log.Error(err)
      }

      _, err = w.Write([]byte("\n"))
      if err != nil {
        Log.Error(err)
      }
    }
  }
}
