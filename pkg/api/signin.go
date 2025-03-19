package api

import (
	"crypto/md5"
	"hash/crc64"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	CRCPass uint64
	jwt.RegisteredClaims
}

type Login struct {
	Password string `json:"password"`
}

type Signin struct {
	Token string `json:"token"`
}

const Secret = `secret word`

func AuthToken(crc uint64) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &Claims{
		CRCPass: crc,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwthash := md5.Sum([]byte(Secret))
	return tok.SignedString(jwthash[:])
}

func CRC64(data []byte) uint64 {
	return crc64.Checksum(data, crc64.MakeTable(crc64.ECMA))
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var token string
	var login Login

	err := parseJson(r, &login)
	if err != nil {
		errorJson(w, err)
		return
	}
	original := os.Getenv("TODO_PASSWORD")
	if original != login.Password {
		errorJson(w, "Неверный пароль")
		return
	}

	token, err = AuthToken(CRC64([]byte(login.Password)))
	if err != nil {
		errorJson(w, err)
		return
	}
	//	md := md5.Sum([]byte(token))
	//	stoken := hex.EncodeToString(md[:])
	writeJson(w, Signin{Token: token})
}
