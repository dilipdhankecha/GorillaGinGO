package main

import (
	"Gorrila/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func connectWithDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:smart@tcp(localhost:3306)/poc")
	if err != nil {
		log.Fatal(" Error occur while create connection with Database.. ")
	}
	return db
}

func main() {
	fmt.Println("GOrilla Goram Example")
	router := mux.NewRouter()
	router.HandleFunc("/poc/gorilla/gorm/list", getEmployeeList).Methods("GET")
	router.HandleFunc("/poc/gorilla/gorm/list/{id}", getEmployeeById).Methods("GET")
	router.HandleFunc("/poc/gorilla/gorm/create", addEmployee).Methods("POST")
	router.HandleFunc("/poc/gorilla/gorm/remove/{id}", removeEmployee).Methods("DELETE")
	router.HandleFunc("/poc/gorilla/gorm/update/{id}", updateEmployee).Methods("PUT")
	http.ListenAndServe(":8080", router)

}

func getEmployeeList(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	var employees []model.Employee
	selectAllDb := db.Find(&employees)
	if selectAllDb.Error != nil {
		json.NewEncoder(response).Encode(model.GetResponse("Error while Get All Employee Details", 400, nil))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(model.GetResponse("Employee Return SuccessFully", 200, employees))
}

func getEmployeeById(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	variables := mux.Vars(request)
	id, err := strconv.Atoi(variables["id"])
	if err != nil {
		json.NewEncoder(response).Encode(model.GetResponse("Error while convert String to int", 400, nil))
		return
	}
	var employees []model.Employee
	//Select Data by where condition
	selectQeury := db.Where("id=?", id).Find(&employees)
	if selectQeury.Error != nil {
		json.NewEncoder(response).Encode(model.GetResponse("Error while Get All Employee Details", 400, nil))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(model.GetResponse("Employee Return SuccessFully", 200, employees))
}

func addEmployee(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	var e model.Employee
	json.NewDecoder(request.Body).Decode(&e)
	insertDb := db.Create(&e)
	if insertDb.Error != nil {
		json.NewEncoder(response).Encode(model.GetResponse("Error while Insert Employees Data", 400, nil))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(model.GetResponse("Employee Added SuccessFully", 200, nil))
}

func removeEmployee(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	variables := mux.Vars(request)
	id, err := strconv.Atoi(variables["id"])
	if err != nil {
		json.NewEncoder(response).Encode(model.GetResponse("Error while retrive Id from PathVariable", 400, nil))
		return
	}
	deleteDB := db.Where("id =?", id).Delete(&model.Employee{})
	if deleteDB.Error != nil {
		json.NewEncoder(response).Encode(model.GetResponse("Error while delete Employee", 400, nil))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(model.GetResponse("Employee Deleted SuccessFully", 200, nil))
}

func updateEmployee(response http.ResponseWriter, request *http.Request) {
	db := connectWithDatabase()
	defer db.Close()
	variables := mux.Vars(request)
	id, err := strconv.Atoi(variables["id"])
	if err != nil {
		json.NewEncoder(response).Encode(model.GetResponse("Error while retrive Id from PathVariable", 400, nil))
		return
	}
	e := model.Employee{
		Id: id,
	}
	findDb := db.Find(&e)
	if findDb.Error != nil {
		json.NewEncoder(response).Encode(model.GetResponse("FindDB ::-- Error while Update Employee", 400, nil))
		return
	}
	json.NewDecoder(request.Body).Decode(&e)
	saveDb := db.Save(&e)
	if saveDb.Error != nil {
		json.NewEncoder(response).Encode(model.GetResponse("SaveDB ::-- Error while Update Employee", 400, nil))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(model.GetResponse("Employee Updated SuccessFully", 200, nil))
}
