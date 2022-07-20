package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)





func jsonHandler(w http.ResponseWriter, r *http.Request){
	var jsonMap = make(map[string]interface{})
	jsonArticle := `{"id": "BM-1347", "title": "The underage storm", "Content": "The creatives' careers can easily get uncreative but yet creative...", "Summary": "Seeking freedom"}`  
	json.Unmarshal([]byte(jsonArticle), &jsonMap)
	fmt.Fprint(w, "Here's some json converted to map:\n", "ID:", jsonMap["id"], "\n","Title:",jsonMap["title"], "\n","Content:",jsonMap["Content"], "\n","Summary:",jsonMap["Summary"])
}

func main(){
	r:= mux.NewRouter()
	r.HandleFunc("/", jsonHandler).Methods("GET")
	http.Handle("/",r)
	http.ListenAndServe(":8000", nil)
}