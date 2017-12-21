package config

import (
  "encoding/json"
  "log"
  "os"

  g "github.com/xanzy/go-gitlab"
)

type GitlabCfg_t struct {
  Url                   string
  Token                 string
  Group                 string
  TLSInsecureSkipVerify bool
}

type GitlabState_t struct {
  Client    *g.Client
  GroupID   int
}

type Configuration struct {
  GitlabCfg     GitlabCfg_t
  GitlabState   GitlabState_t
  Log           *log.Logger
}

func ReadFile(configFilePath string) *Configuration {

  file, err := os.Open(configFilePath)
  if err != nil {
    log.Fatalf("Can't open config file(%s), %s", configFilePath, err)
  }
  defer file.Close()

  decoder := json.NewDecoder(file)
  config := Configuration{GitlabCfg: GitlabCfg_t{TLSInsecureSkipVerify: true,
                                                  Group: "ansible-roles"}}

  err = decoder.Decode(&config)

  if err != nil {
    log.Fatalf("Can't decode config file(%s), %s", configFilePath, err)
  }

  config.Log = log.New(os.Stderr, "roles-ws: ", log.Llongfile)

  return &config
}
