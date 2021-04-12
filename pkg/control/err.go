package control

import (
  "encoding/json"
  "fmt"
  "net/http"
)

func (a *Api) NotFound(props HandlerProps, reply string) {
  props.Writer.Header().Set("Content-Type", "application/json")
  props.Writer.WriteHeader(http.StatusNotFound)

  resp := &Response{
    ErrMsg: fmt.Sprintf("URL path '%s' not found", props.Request.URL.Path),
    Reply:  reply,
  }
  body, _ := json.Marshal(resp)

  body = append(body, []byte("\n")...)

  _, err := props.Writer.Write(body)
  if err != nil {
    Log.Error(err)
  }
}

func (a *Api) ServiceUnavailable(w http.ResponseWriter, errMsg string) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusServiceUnavailable)

  resp := &Response{
    ErrMsg: errMsg,
  }
  body, _ := json.Marshal(resp)
  body = append(body, []byte("\n")...)

  _, err := w.Write(body)
  if err != nil {
    Log.Error(err)
  }
}
