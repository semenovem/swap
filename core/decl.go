package core

type BaseConfig struct {
  DevMode     bool   `yaml:"dev-mode"`
  TestMode    bool   `yaml:"test-mode"`
  LogLevel    string `yaml:"log-level"`
  LogMode     string `yaml:"log-mode"`
  LogDest     string `yaml:"log-dest"`
  CmdPort     int    `yaml:"cmd-port"`
  CmdLogLevel string `yaml:"cmd-log-level"`
}
