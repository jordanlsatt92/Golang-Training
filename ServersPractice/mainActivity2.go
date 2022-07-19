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

// ASSIGNMENT 1
// func handler(w http.ResponseWriter, _ *http.Request){
// 	fmt.Fprint(w, "Hello World Thanks for connecting me")
// }

// func handler1(w http.ResponseWriter, _ *http.Request){
// 	location,_ := time.LoadLocation("America/New_York")
// 	now:=time.Now().In(location)
// 	p:=make(map[string] string)
// 	p["now"] =now.Format(time.ANSIC)
// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(p)
// }

// func handler2(w http.ResponseWriter, _ *http.Request){
// 	location,_ := time.LoadLocation("Asia/Tokyo")
// 	now:=time.Now().In(location)
// 	p:=make(map[string] string)
// 	p["now"] =now.Format(time.ANSIC)
// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(p)
// }

// func handler3(w http.ResponseWriter, _ *http.Request){
// 	location,_ := time.LoadLocation("Europe/London")
// 	now:=time.Now().In(location)
// 	p:=make(map[string] string)
// 	p["now"] =now.Format(time.ANSIC)
// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(p)
// }

// func Now(w http.ResponseWriter, r *http.Request){
// 	now:=time.Now()
// 	p:=make(map[string] string)
// 	p["now"] =now.Format(time.ANSIC)
// 	w.Header().Set("Content-Type","application/json")
// 	json.NewEncoder(w).Encode(p)
// }

// func main(){
// 	// fmt.Println("")
// 	r:=mux.NewRouter()
// 	r.HandleFunc("/", handler).Methods("GET")
// 	r.HandleFunc("/newyork", handler1).Methods("GET")
// 	r.HandleFunc("/tokyo", handler2).Methods("GET")
// 	r.HandleFunc("/london", handler3).Methods("GET")
// 	r.HandleFunc("/now", Now).Methods("GET")
// 	http.Handle("/", r)
// 	http.ListenAndServe(":8000", nil)
// }

/*
Assignment 2 )  Write  a go program to get the client IP and reflect/echo back to client page ,after this add feature to 
server so that server could block to a particular client with a message
Hint  : func ReadUserIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
    return IPAddress
}
X-Real-Ip - fetches first true IP (if the requests sits behind multiple NAT sources/load balancer)

X-Forwarded-For - if for some reason X-Real-Ip is blank and does not return response, get from X-Forwarded-For

Remote Address - last resort (usually won't be reliable as this might be the last ip or if it is a naked http request to server ie no load balancer)

*/

