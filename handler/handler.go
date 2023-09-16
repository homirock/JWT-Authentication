package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



type UserDetails struct {
	User     string
	Password string
}

func Registration(w http.ResponseWriter, r *http.Request) {
	Map:= make(map[string]string)
	userdeatils := &UserDetails{}
	err := json.NewDecoder(r.Body).Decode(&userdeatils)
	fmt.Println(userdeatils.User)
	if err != nil {
		log.Println("Unsupported format")
	} else if r.Method=="POST" {
		//store in cache
		Map[userdeatils.User] = userdeatils.Password
	}
	w.WriteHeader(http.StatusOK)
}

func Server() {
	fmt.Println("Starting the server")
	http.HandleFunc("/registration", Registration)
	http.ListenAndServe(":8080", nil)
}
