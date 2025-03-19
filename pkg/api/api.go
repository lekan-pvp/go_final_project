package api

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var formatDate string

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		delTaskHandler(w, r)
	}
}

func JWTVerify(jwtData string) (claims *Claims, token *jwt.Token, err error) {
	claims = &Claims{}
	if len(jwtData) > 0 {
		token, err = jwt.ParseWithClaims(jwtData, claims,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				jwthash := md5.Sum([]byte(Secret))
				return jwthash[:], nil
			})
	}
	return
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var jwt string
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}
			claims, token, err := JWTVerify(jwt)
			var valid bool
			if err == nil && token != nil && token.Valid {
				if claims.CRCPass > 0 {
					valid = CRC64([]byte(pass)) == claims.CRCPass
				}
			}
			if !valid {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

func Init(format string) {

	formatDate = format

	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", auth(taskHandler))
	http.HandleFunc("/api/tasks", auth(tasksHandler))
	http.HandleFunc("/api/task/done", auth(doneTaskHandler))
	http.HandleFunc("/api/signin", signinHandler)
}
