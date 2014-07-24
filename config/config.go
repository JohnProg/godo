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
  Development EnvironmentConfiguration  `json:"development"`
  Production EnvironmentConfiguration   `json:"production"`
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
  e = json.Unmarshal(file, &config)
  if e != nil {
      fmt.Printf("Error Unmarshal config file: %v\n", e)
      os.Exit(1)
  } 
  if(env == "production") {
    CurrentConfiguration = config.Production
  } else {
    CurrentConfiguration = config.Development
  }
  return
}