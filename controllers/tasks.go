package tasks

import (
  //"fmt"
  "net/http"
  "github.com/mb-dev/godo/models"
  "github.com/mb-dev/godo/util"
)

func GetAllTasks(rw http.ResponseWriter, req *http.Request) { 
  tasks, err := models.DatabaseInstance(util.ContextGetUserId(req)).Tasks.GetAll()
  if err != nil { panic(err) }
  util.WriteJson(rw, tasks)
  return
}

func CreateTask(rw http.ResponseWriter, req *http.Request) {
  var (
    task models.Task
    err error
  )
  err = util.ReadJson(req, &task)
  if err != nil { panic(err) }

  err = models.DatabaseInstance(util.ContextGetUserId(req)).Tasks.Create(&task)
  if err != nil { panic(err) }
  util.WriteJson(rw, task)
  return
}

func UpdateTask(rw http.ResponseWriter, req *http.Request) {
  var (
    task models.Task
    err error
  )
  err = util.ReadJson(req, &task)
  if err != nil { panic(err) }

  err = models.DatabaseInstance(util.ContextGetUserId(req)).Tasks.Update(&task)
  if err != nil { panic(err) }
  util.WriteJson(rw, task)
  return
}
