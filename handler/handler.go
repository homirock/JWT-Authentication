package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	Map       = make(map[string]string)
	Secretkey = "my_secrete_key"
)

type UserDetails struct {
	User     string
	Password string
}

type Token struct {
	User string
	Tkn  string
}

func Registration(w http.ResponseWriter, r *http.Request) {
	userdeatils := &UserDetails{}
	err := json.NewDecoder(r.Body).Decode(&userdeatils)
	fmt.Println(userdeatils.User)
	if err != nil {
		log.Println("Unsupported format")
	} else if r.Method == "POST" {
		//store in cache
		Map[userdeatils.User] = userdeatils.Password
	}
	w.WriteHeader(http.StatusOK)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	user := &UserDetails{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

	}
	if _, ok := Map[user.User]; !ok {
		//not registered. Go to registarion page
	}
	//generate JWT token
	tkn, err := GenerateJWTtoken(user.User, user.Password)
	var token Token
	token.User = user.User
	token.Tkn = tkn
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func GenerateJWTtoken(user, passwd string) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
func ValidateJWTtoken(tkn string) error {
	jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		return Secretkey, nil
	})
	return fmt.Errorf("Unsupported Token")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var token string
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		fmt.Println("Unsupported token type")
	}
	if ValidateJWTtoken(token) == nil {
		//go to web page
		fmt.Println("Go to Web page")
	}
}

func Server() {
	fmt.Println("Starting the server")
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/login", Login)
	http.ListenAndServe(":8080", nil)
}
