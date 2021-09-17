package tinkoffapi

import (
  "encoding/json"
  "fmt"
  "github.com/pkg/errors"
  "io"
  "io/ioutil"
  "net/http"
)

func (t *TfApi) entryRequest(ent *entry, body io.Reader, reply interface{}) error {
  url := t.makeUrl(ent.path)

  t.log.Infof("Запрос к api: %s", url)

  req, err := http.NewRequest(ent.typ, url, body)
  if err != nil {
    err = errors.Wrap(err, "Can't create register http request")
    t.log.Error(err)
    return err
  }

  req.Header.Add("Authorization", "Bearer "+t.token)
  resp, err := t.httpClient.Do(req)
  if err != nil {
    err = errors.Wrap(err, "Can't send register request")
    t.log.Error(err)
    return err
  }

  defer func() {
    err := resp.Body.Close()
    if err != nil {
      t.log.Warn(err)
    }
  }()

  if resp.StatusCode >= 200 && resp.StatusCode < 300 {
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      err = errors.Wrapf(err, "Can't read register response: %s", err)
      t.log.Error(err)
      return err
    }

    err = json.Unmarshal(respBody, &reply)
    if err != nil {
      err = errors.Wrapf(err, "Can't unmarshal register response: '%s' \nwith error: %s", string(respBody), err)
      t.log.Error(err)
      return err
    }
    return nil
  }

  err = fmt.Errorf("register, bad response code '%s' from '%s'", resp.Status, url)
  t.log.Error(err)
  t.log.Debugf("Body = %+v\nHeader = %+v\nStatus = %+v\n", resp.Body, resp.Header, resp.Status)

  return err
}

func (t *TfApi) makeUrl(path string) string {
  return fmt.Sprintf("%s%s", t.cfg.UrlEntry, path)
}
