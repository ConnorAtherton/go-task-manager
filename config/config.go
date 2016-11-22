package config

import (
  "os"
  "io/ioutil"
  "fmt"
  "errors"
  "path/filepath"
  "github.com/BurntSushi/toml"
)

const configFile string = ".tasks.toml"

type Config struct {
	Path      string
	Formatter string
}

func userDir() string {
  return os.Getenv("HOME")
}

func New() (*Config, error) {
  conf := Config{}

  // NOTE: I could use DecodeFile here but I like having to ioutil when still learning
  data, readErr := ioutil.ReadFile(filepath.Join(userDir(), configFile))
  _, decodeErr := toml.Decode(string(data), &conf)

  if readErr != nil || decodeErr != nil {
    fmt.Println(errors.New("Malformed config file"), readErr, decodeErr)
    os.Exit(1)
  }

  // NOTE: Need to figure out how to set default values for fields here
  return &conf, nil
}
