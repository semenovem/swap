package build

import (
  "encoding/json"
  "os"
)

const UNKNOWN = "UNKNOWN"

var (
  version, hashCommit, branch, time string
)

type BuildInfo struct {
  Version    string
  HashCommit string
  Branch     string
  Time       string
}

func (b *BuildInfo) String() string {
  body, _ := json.Marshal(b)
  return string(body)
}

func Ver() string {
  if version != "" {
    return version
  }
  v := os.Getenv("ENV_PKG_APP_BUILD_DEV_MODE_VERSION")
  if v != "" {
    return v
  }
  return UNKNOWN
}

func HashCommit() string {
  if hashCommit != "" {
    return hashCommit
  }
  v := os.Getenv("ENV_PKG_APP_BUILD_DEV_MODE_HASH_COMMIT")
  if v != "" {
    return v
  }
  return UNKNOWN
}

func Branch() string {
  if branch != "" {
    return branch
  }
  v := os.Getenv("ENV_PKG_APP_BUILD_DEV_MODE_BRANCH")
  if v != "" {
    return v
  }
  return UNKNOWN
}

func Time() string {
  if time != "" {
    return time
  }
  v := os.Getenv("ENV_PKG_APP_BUILD_DEV_MODE_TIME")
  if v != "" {
    return v
  }
  return UNKNOWN
}

func Info() *BuildInfo {
  return &BuildInfo{
    Version:    Ver(),
    HashCommit: HashCommit(),
    Branch:     Branch(),
    Time:       Time(),
  }
}

func ToString() string {
  return Info().String()
}
