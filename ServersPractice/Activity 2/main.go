package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Assignment 2

var ipMap = make(map[string]string)

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func ipHandler (w http.ResponseWriter, r *http.Request){
	ipAddress := ReadUserIP(r)
	val := ipMap[ipAddress]
	if val == ""{
		fmt.Fprint(w, "Your IP Address:",ipAddress,"; This is your first time!")
		date := time.Now().Format("2006/01/02")
		time := time.Now().Format("15:04:05")
		ipMap[ipAddress] = "Date: " + date +" Time: " + time
	}else{
		lastVisit := ipMap[ipAddress]
		date := time.Now().Format("2006/01/02")
		time := time.Now().Format("15:04:05")
		ipMap[ipAddress] = "Date: " + date +" Time: " + time
		fmt.Fprint(w, "Your IP Address: ", ipAddress, "; Your last visit:", lastVisit)		
	}
}

func main(){
	// fmt.Println("")
	r:=mux.NewRouter()
	r.HandleFunc("/", ipHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

