package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/xid"

	jwt "github.com/dgrijalva/jwt-go"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Up!\n")
}

// TokenRequest is
type TokenRequest struct {
	Name     string   `json:"name"`
	LoA      int64    `json:"loa"`
	Scope    []string `json:"scope"`
	Audience []string `json:"audience"`
}

// TokenResponse is
type TokenResponse struct {
	Token string `json:"token"`
}

var mySigningKey = []byte("secret")

// new ..
func createToken(w http.ResponseWriter, r *http.Request) {
	var msgBody TokenRequest

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&msgBody)
	if err != nil {
		fmt.Println("error")

		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(msgBody)

	fmt.Println(msgBody.Name)
	if msgBody.Name == "" {
		fmt.Println("no name")
	} else {
		fmt.Println(msgBody.Audience)
	}
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	t := time.Now()

	iat := t.Unix()
	exp := t.Add(time.Duration(10) * time.Minute).Unix()

	jti := xid.New().String()
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["loa"] = msgBody.LoA
	claims["scope"] = msgBody.Scope
	claims["audience"] = msgBody.Audience
	claims["jti"] = jti
	claims["iat"] = iat
	claims["exp"] = exp
	claims["sub"] = msgBody.Name

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	logNew(msgBody.Name, jti)

	rsp := TokenResponse{Token: tokenString}

	returnToken(w, r, rsp)

}

func returnToken(w http.ResponseWriter, r *http.Request, token TokenResponse) {

	err := json.NewEncoder(w).Encode(token)

	if err != nil {

		log.Printf("token : %s", token)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

func logNew(name string, id string) {

	t := "Login"
	actStr := fmt.Sprintf("[Authentication] by %s", name)
	ref := fmt.Sprintf("%s", id)

	logActivity(t, actStr, ref)
}
