package util

import (
  "net/http"
  "github.com/dgrijalva/jwt-go"
  "github.com/mb-dev/godo/config"
)

func JWTMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) ([]byte, error) {
    return []byte(config.CurrentConfiguration.KeySecret), nil
  })

  if err != nil || !token.Valid {
    rw.WriteHeader(http.StatusForbidden)
    return
  } 
  ContextSetUserId(req, token.Claims["id"].(string))
  next(rw, req)
} 