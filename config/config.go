package config

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
  "os"
)

type EnvironmentConfiguration struct {
  DbHost    string `json:"dbHost"`
  DbName    string `json:"dbName"`
  KeySecret string `json:"keySecret"`
  ServerPort string `json:"serverPort"`
}

type Configuration struct {
  development EnvironmentConfiguration
  production EnvironmentConfiguration
}

const AppName string = "godo"
var CurrentConfiguration EnvironmentConfiguration


func LoadConfiguration(env string) {
  file, e := ioutil.ReadFile("./config/config.json")
  if e != nil {
      fmt.Printf("Configuration file not found: %v\n", e)
      os.Exit(1)
  } 
  var config Configuration
  json.Unmarshal(file, &config)
  if(env == "production") {
    CurrentConfiguration = config.production
  } else {
    CurrentConfiguration = config.development
  }
  return
}