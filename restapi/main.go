package main

// import (
// 	// "fmt"
// 	"net/http"
// 	"encoding/json"
// 	"github.com/gorilla/mux"
// 	"log"
// 	// "database/sql"
// 	// _ "github.com/go-sql-driver/mysql"
// )

// //Employee is table structure
// type Employee struct {
// 	ID   string `json:"id,omitempty"`
// 	Name string `json:"name,omitempty"`
// 	Post string `json:"post,omitempty"`
// }

// //Employees is the slice which store dummy data
// var Employees []Employee

// func showAllEmployees(w http.ResponseWriter, req *http.Request)  {
// 	// fmt.Fprintf(w,"All Eployees details are mentioned\n")
// 	json.NewEncoder(w).Encode(Employees)
// }

// //recciev JSON data from postman
// func addEmployee(w http.ResponseWriter, req *http.Request)  {
// 	params := mux.Vars(req)
// 	var employee Employee
// 	_ = json.NewDecoder(req.Body).Decode(&employee)
// 	// "_" used because it _ can store values which may not be used further without giving error
// 	employee.ID = params["id"]
// 	Employees = append(Employees, employee)
// 	json.NewEncoder(w).Encode(Employees)
// }

// func deleteEmployee(w http.ResponseWriter, req *http.Request)  {
// 	params := mux.Vars(req)
// 	for i, j:=range Employees {
// 		if j.ID == params["id"] {
// 			Employees = append(Employees[:i], Employees[i+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(Employees)
// }

// func updateEmployee(w http.ResponseWriter, req *http.Request) {
// 	addEmployee(w,req)
// 	deleteEmployee(w,req)
// }

// func main1()  {
// 	router := mux.NewRouter()
// 	Employees = append(Employees, Employee{ID: "1", Name:"ayush", Post:"SDE"})
// 	Employees = append(Employees, Employee{ID: "2", Name:"dummy1", Post:"HR"})
// 	Employees = append(Employees, Employee{ID: "3", Name:"dummy3", Post:"Manager"})

// 	router.HandleFunc("/employees", showAllEmployees).Methods("GET")
// 	router.HandleFunc("/addemployee/{id}", addEmployee).Methods("POST")
// 	router.HandleFunc("/deleteemployee/{id}", deleteEmployee).Methods("DELETE")
// 	router.HandleFunc("/updateemployee/{id}", updateEmployee).Methods("PUT")
// 	log.Fatal(http.ListenAndServe(":8081",router))

// }
// //jakeer.husen