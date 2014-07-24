package util

import (
  "net/http"
)

func CORSMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  if origin := req.Header.Get("Origin"); origin != "" {
      rw.Header().Set("Access-Control-Allow-Origin", origin)
      rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
      rw.Header().Set("Access-Control-Allow-Headers",
          "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  }
  // Stop here if its Preflighted OPTIONS request
  if req.Method == "OPTIONS" {
      return
  }

  next(rw, req)
}