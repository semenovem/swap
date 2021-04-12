package control

import (
  "net/http"
)

type Response struct {
  Ok     bool   `json:"ok"`
  ErrMsg string `json:"errMsg"`
  Reply  string `json:"reply"`
}

type HandlerProps struct {
  Writer  http.ResponseWriter
  Request *http.Request
}

type RouteHandler func(props HandlerProps) (*Response, error)

type Route struct {
  Path    string
  Handler RouteHandler
}
