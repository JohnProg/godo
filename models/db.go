package models

import (
  "strings"
  "labix.org/v2/mgo"
  "github.com/mb-dev/godo/config"
)

type Database struct {
  Tasks TaskCollection
}

var (
  mgoDatabase *mgo.Database
)

func fullTableName(tableName string, userId string) (string) {
  return strings.Join([]string{"data", userId, config.AppName, tableName}, "_")
}

func DatabaseInstance(userId string) (*Database) {
  tasksCollection := mgoDatabase.C(fullTableName("tasks", userId))
  tasksModel := TaskCollection{c: tasksCollection}
  return &Database{Tasks: tasksModel}
}

func ConnectToDatabase() {
  mongoSession, err := mgo.Dial(config.CurrentConfiguration.DbHost)
  if err != nil { panic(err) }

  mgoDatabase = mongoSession.DB(config.CurrentConfiguration.DbName)
}