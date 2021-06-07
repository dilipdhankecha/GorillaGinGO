package main

import (
	"Gorrila/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("==== POC with Gorrilla ====")
	router := mux.NewRouter()
	router.HandleFunc("/poc/gorrilla/create", addEmployee).Methods("POST")
	router.HandleFunc("/poc/gorrilla/remove/{id}", removeEmployee).Methods("DELETE")
	router.HandleFunc("/poc/gorrilla/update/{id}", updateEmployee).Methods("PUT")
	router.HandleFunc("/poc/gorrilla/list", getAllEmployee).Methods("GET")
	router.HandleFunc("/poc/gorrilla/list/{id}", getAllEmployeeById).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func connectWithDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:smart@tcp(localhost:3306)/poc")
	if err != nil {
		log.Fatal(" Error occur while create connection with Database.. ")
	}
	return db
}

func addEmployee(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	var e model.Employee
	json.NewDecoder(request.Body).Decode(&e)
	query := "Insert Into employee (id,fName,lName,email,position,experience) VALUES (?,?,?,?,?,?)"
	_, err := db.Exec(query, &e.Id, &e.FName, &e.LName, &e.Email, &e.Position, &e.Experience)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(model.GetResponse("Error while add Employee", 400, nil))
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(model.GetResponse("Employee Added SuccessFully", 200, nil))
}

func removeEmployee(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	variables := mux.Vars(request)
	id, err := strconv.Atoi(variables["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(model.GetResponse("Error while retrive Id from PathVariable", 400, nil))
	}
	_, err = db.Exec("Delete From employee where id = ?", id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(model.GetResponse("Error while delete Employee", 400, nil))
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(model.GetResponse("Employee Deleted SuccessFully", 200, nil))
}

func updateEmployee(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	variables := mux.Vars(request)
	id, err := strconv.Atoi(variables["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(model.GetResponse("Error while retrive Id from PathVariable", 400, nil))
	}
	var e model.Employee
	json.NewDecoder(request.Body).Decode(&e)
	_, err = db.Exec("Update employee SET fname= ?,lname=?, position=?,experience =? where id = ?", &e.FName, &e.LName, &e.Position, &e.Experience, id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(model.GetResponse("Error while Update Employee", 400, nil))
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(model.GetResponse("Employee Updated SuccessFully", 200, nil))
}

func getAllEmployee(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	query := "SELECT * FROM employee"
	result, err := db.Query(query)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(model.GetResponse("Error while Get All Employee Details", 400, nil))
	}
	var employees []model.Employee
	for result.Next() {
		var e model.Employee
		err := result.Scan(&e.Id, &e.FName, &e.LName, &e.Email, &e.Position, &e.Experience)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(model.GetResponse("Error while Scan Employee Details", 400, nil))
		}
		employees = append(employees, e)
	}
	defer result.Close()
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(model.GetResponse("Employee Return SuccessFully", 200, employees))
}

func getAllEmployeeById(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	variables := mux.Vars(request)
	id, err := strconv.Atoi(variables["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(model.GetResponse("Error while retrive Id from PathVariable", 400, nil))
	}
	result, err := db.Query("Select * FROM employee where id=?", id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(model.GetResponse("Error while Get Employee Details", 400, nil))
	}
	var employees []model.Employee
	for result.Next() {
		var e model.Employee
		err := result.Scan(&e.Id, &e.FName, &e.LName, &e.Email, &e.Position, &e.Experience)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(model.GetResponse("Error while Scan Employee Details", 400, nil))
		}
		employees = append(employees, e)
	}
	defer result.Close()
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(model.GetResponse("Employee Return SuccessFully", 200, employees))
}
