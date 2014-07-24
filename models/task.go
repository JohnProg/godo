package models

import (
  //"fmt"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "github.com/mb-dev/godo/util"
)

type TaskCollection struct {
  c *mgo.Collection
}

type Task struct {
  Id        bson.ObjectId `json:"id"           bson:"_id"`
  Title     string        `json:"title"        bson:"title"`
  Dates     []int         `json:"dates"        bson:"dates"`
  Comment   string        `json:"comment"      bson:"comment"`
  Completed bool          `json:"completed"    bson:"completed"`
}

func (model TaskCollection) GetAll() (tasks []Task, err error) {
  err = model.c.Find(nil).All(&tasks)
  if tasks == nil {
    tasks = []Task{}
  }
  return
}

func (model TaskCollection) Create(newTask *Task) (err error) {
  if newTask.Id.Hex() == "" {
      newTask.Id = bson.NewObjectId()
  }
  if newTask.Dates == nil || len(newTask.Dates) == 0 {
    newTask.Dates = []int{util.TodayAsInteger()}
  }
  newTask.Comment = ""
  newTask.Completed = false
  err = model.c.Insert(newTask)
  return
}

func (model TaskCollection) Update(updatedTask *Task) (err error) {
  err = model.c.UpdateId(updatedTask.Id, updatedTask)
  return
}
