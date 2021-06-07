package main

import (
	"Gorrila/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("===== GIN POC =====")
	router := gin.Default()
	router.GET("/poc/gin/list", getEmployeeList)
	router.GET("/poc/gin/list/:id", getEmployeeById)
	router.POST("/poc/gin/create", addEmployee)
	router.DELETE("/poc/gin/remove/:id", removeEmployee)
	router.PUT("/poc/gin/update/:id", updateEmployee)
	router.Run(":8080")
}

func connectWithDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:smart@tcp(localhost:3306)/poc")
	if err != nil {
		log.Fatal(" Error occur while create connection with Database.. ")
	}
	return db
}

func getEmployeeList(c *gin.Context) {
	db := connectWithDatabase()
	defer db.Close()
	query := "SELECT * FROM employee"
	result, err := db.Query(query)
	if err != nil {
		c.JSON(400, model.GetResponse("Error while Get All Employee Details", 400, nil))
		return
	}
	var employees []model.Employee
	for result.Next() {
		var e model.Employee
		err := result.Scan(&e.Id, &e.FName, &e.LName, &e.Email, &e.Position, &e.Experience)
		if err != nil {
			c.JSON(400, model.GetResponse("Error while Scan Employee Details", 400, nil))
			return
		}
		employees = append(employees, e)
	}
	defer result.Close()
	c.JSON(200, model.GetResponse("Employee Return SuccessFully", 200, employees))
}

func getEmployeeById(c *gin.Context) {
	db := connectWithDatabase()
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, model.GetResponse("Error while convert String to int", 400, nil))
		return
	}
	result, err := db.Query("Select * FROM employee where id=?", id)
	if err != nil {
		c.JSON(400, model.GetResponse("Error while Get All Employee Details", 400, nil))
		return
	}
	var employees []model.Employee
	for result.Next() {
		var e model.Employee
		err := result.Scan(&e.Id, &e.FName, &e.LName, &e.Email, &e.Position, &e.Experience)
		if err != nil {
			c.JSON(400, model.GetResponse("Error while Scan Employee Details", 400, nil))
			return
		}
		employees = append(employees, e)
	}
	defer result.Close()
	c.JSON(200, model.GetResponse("Employee Return SuccessFully", 200, employees))
}

func addEmployee(c *gin.Context) {
	db := connectWithDatabase()
	defer db.Close()
	var e model.Employee
	json.NewDecoder(c.Request.Body).Decode(&e)
	query := "Insert Into employee (id,fName,lName,email,position,experience) VALUES (?,?,?,?,?,?)"
	_, err := db.Exec(query, &e.Id, &e.FName, &e.LName, &e.Email, &e.Position, &e.Experience)
	if err != nil {
		c.JSON(400, model.GetResponse("Error while add Employee", 400, nil))
		return
	}
	c.JSON(200, model.GetResponse("Employee Added SuccessFully", 200, nil))
}

func removeEmployee(c *gin.Context) {
	db := connectWithDatabase()
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, model.GetResponse("Error while retrive Id from PathVariable", 400, nil))
		return
	}
	_, err = db.Exec("Delete From employee where id = ?", id)
	if err != nil {
		c.JSON(400, model.GetResponse("Error while delete Employee", 400, nil))
		return
	}
	c.JSON(200, model.GetResponse("Employee Deleted SuccessFully", 200, nil))
}

func updateEmployee(c *gin.Context) {
	db := connectWithDatabase()
	defer db.Close()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, model.GetResponse("Error while retrive Id from PathVariable", 400, nil))
		return
	}
	var e model.Employee
	json.NewDecoder(c.Request.Body).Decode(&e)
	_, err = db.Exec("Update employee SET fname= ?,lname=?, position=?,experience =? where id = ?", &e.FName, &e.LName, &e.Position, &e.Experience, id)
	if err != nil {
		c.JSON(400, model.GetResponse("Error while Update Employee", 400, nil))
		return
	}
	c.JSON(200, model.GetResponse("Employee Updated SuccessFully", 200, nil))
}
