package util

import (
  "net/http"
  "github.com/gorilla/context"
)

const userIdKey = "userId"

func ContextSetUserId(req *http.Request, userId string) {
  context.Set(req, userIdKey, userId)
}

func ContextGetUserId(req *http.Request) (string) {
  return context.Get(req, userIdKey).(string)
}