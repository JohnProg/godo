package main

import (
  "flag"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "github.com/mb-dev/godo/config"
  "github.com/mb-dev/godo/models"
  "github.com/mb-dev/godo/util"
  "github.com/mb-dev/godo/controllers"
)

func routeTasks(r *mux.Router) {
  r.HandleFunc("/tasks", tasks.GetAllTasks).Methods("GET")
  r.HandleFunc("/tasks", tasks.CreateTask).Methods("POST")
  r.HandleFunc("/tasks/{id}", tasks.UpdateTask).Methods("PUT")
}


func setupRoutes(r *mux.Router) {
  routeTasks(r)
}

func main() {
  var env = flag.String("env", "development", "ENV - development or production")
  flag.Parse()

  config.LoadConfiguration(*env)
  models.ConnectToDatabase()

  router := mux.NewRouter()
  setupRoutes(router)
  
  n := negroni.Classic()
  n.Use(negroni.HandlerFunc(util.CORSMiddleware)) // handle CORS
  n.Use(negroni.HandlerFunc(util.JWTMiddleware)) // handle authorization
  n.UseHandler(router)
  n.Run(config.CurrentConfiguration.ServerPort)
}