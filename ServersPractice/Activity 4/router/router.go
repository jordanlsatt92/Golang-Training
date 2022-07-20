package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type task struct {
	ID		int    `json:ID`
	Name	string `json:Name`
	Dept string `json:Department`
	BasicPay float64 `json:Basic Pay`
	HRA float64	`json:HRA`
	GrossPay float64 `json:Gross Pay`
}

type allTasks []task // list of tasks

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Dept: "IT",
		BasicPay: 50000,
		HRA: 0,
		GrossPay: 0,
	},
	{
        ID: 2,
        Name: "Toby",
        Dept: "HR",
        BasicPay: 50000,
        HRA: 0,
        GrossPay: 0,
    },
    {
        ID: 3,
        Name: "John",
        Dept: "IT/Security",
        BasicPay: 50000,
        HRA: 0,
        GrossPay: 0,
    },
	{
        ID: 4,
        Name: "Holly",
        Dept: "HR/Security",
        BasicPay: 50000,
        HRA: 0,
        GrossPay: 0,
    },
}
//HTTP GET
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// HTTP Post
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Fprintf(w, "Task with ID %v has been removed successfully", taskID)
		}
	}
}

func calculateTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err:= strconv.Atoi(vars["id"])
	var calculateTask task
	if err != nil{
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for index, task := range tasks {
		if task.ID == taskID {
			calculateTask = task
			tasks = append(tasks[:index], tasks[index+1:]...)
			calculateTask.ID = taskID
			dept:= strings.Split(calculateTask.Dept, "/")
			for _,val := range dept{
				if strings.ToLower(val) == "it"{
					calculateTask.HRA = .1
				}else if strings.ToLower(val) == "security"{
					calculateTask.HRA += 0.14
				}else{
					calculateTask.HRA = .05
				}
			}
			calculateTask.GrossPay = calculateTask.BasicPay*(1+calculateTask.HRA)
			tasks = append(tasks, calculateTask)

			fmt.Fprintf(w, "The task with id %v has been updated successfully", taskID)
			break
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updateTask task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &updateTask)

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			updateTask.ID = taskID
			tasks = append(tasks, updateTask)

			fmt.Fprintf(w, "The task with id %v has been updated successfully", taskID)
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

//ConfigureRouter setup the router
func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/calculate/{id}", calculateTask).Methods("PUT")
	router.HandleFunc("/ping", healthCheck).Methods("GET")
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}