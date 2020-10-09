package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Employee struct {
	ID   int `json:"id"`
	Name string `json:"name,omitempty"`
	Post string `json:"post,omitempty"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
    dbPass := ""
    dbName := "testdb"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func showEmployees(w http.ResponseWriter, req *http.Request)  {
	// fmt.Fprintf(w,"Below are all available employees\n\n")

	db := dbConn()
	mydb,err1 := db.Query("SELECT * FROM employees")
	if(err1 != nil) {
		fmt.Println("\nerror coming from showEmployees func\n")
		panic(err1.Error())
	}


	for mydb.Next() {
		var id int
		var name, post string
		err1 = mydb.Scan(&id,&name,&post)
		var employee Employee
		employee.ID=id
		employee.Name=name
		employee.Post=post
		temp,_ := json.Marshal(employee)
		fmt.Fprintf(w,"%s\n",string(temp))
		// fmt.Fprintf(w,"%d %s %s\n",id,name,post)
	}
}

func add(w http.ResponseWriter, req *http.Request) {
	// param := mux.Vars(req)
	var employee Employee
	_ = json.NewDecoder(req.Body).Decode(&employee)
	name := employee.Name
	post := employee.Post
	
	db := dbConn()
	mydb,err := db.Prepare("INSERT INTO `employees` (`name`, `post`) VALUES (?, ?)")
	if(err != nil) {
		fmt.Println("\nerror coming from add func\n")
		panic(err.Error())
	}
	mydb.Exec(name,post)

	fmt.Fprintf(w,"data name=%s and post=%s added successfully\n",name,post)
}

func delete(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	db := dbConn()
	id := params["id"]

	mydb,err := db.Prepare("DELETE FROM `employees` WHERE id=?")
	if(err != nil) {
		fmt.Println("\nerror coming from delete func\n")
		panic(err.Error())
	}

	mydb.Exec(id)
	fmt.Fprintf(w,"Employee with id = %s deleted successfully\n",id)
}

func update(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	db := dbConn()
	id := params["id"]

	var employee Employee
	_ = json.NewDecoder(req.Body).Decode(&employee)
	name := employee.Name
	post := employee.Post

	mydb,err := db.Prepare("UPDATE `employees` SET name=?,post=? WHERE id=?")
	if(err != nil) {
		fmt.Println("\nerror coming from delete func\n")
		panic(err.Error())
	}

	mydb.Exec(name,post,id)
	fmt.Fprintf(w,"Employee with id = %s updated successfully\n",id)
}

func main() {
	db,err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testdb")
	// err := db.Query("INSERT INTO `employees` (`id`, `name`, `post`) VALUES ('1', 'ayush', 'SDE')")

	if(err != nil) {
		fmt.Println("\n3\n")
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/employees", showEmployees).Methods("GET")
	router.HandleFunc("/add", add).Methods("POST")
	router.HandleFunc("/delete/{id}", delete).Methods("DELETE")
	router.HandleFunc("/update/{id}",update).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081",router))

}