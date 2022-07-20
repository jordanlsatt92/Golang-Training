package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ASSIGNMENT 1
func handler(w http.ResponseWriter, _ *http.Request){
	fmt.Fprint(w, "Hello World Thanks for connecting me")
}

func handler1(w http.ResponseWriter, _ *http.Request){
	location,_ := time.LoadLocation("America/New_York")
	now:=time.Now().In(location)
	p:=make(map[string] string)
	p["now"] =now.Format(time.ANSIC)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(p)
}

func handler2(w http.ResponseWriter, _ *http.Request){
	location,_ := time.LoadLocation("Asia/Tokyo")
	now:=time.Now().In(location)
	p:=make(map[string] string)
	p["now"] =now.Format(time.ANSIC)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(p)
}

func handler3(w http.ResponseWriter, _ *http.Request){
	location,_ := time.LoadLocation("Europe/London")
	now:=time.Now().In(location)
	p:=make(map[string] string)
	p["now"] =now.Format(time.ANSIC)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(p)
}

func Now(w http.ResponseWriter, r *http.Request){
	now:=time.Now()
	p:=make(map[string] string)
	p["now"] =now.Format(time.ANSIC)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(p)
}

func main(){
	// fmt.Println("")
	r:=mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/newyork", handler1).Methods("GET")
	r.HandleFunc("/tokyo", handler2).Methods("GET")
	r.HandleFunc("/london", handler3).Methods("GET")
	r.HandleFunc("/now", Now).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
