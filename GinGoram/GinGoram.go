package main

import (
	"Gorrila/model"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
	All gorm function will return gorm.DB.
	that is contain an Error and response of data.
*/

func connectWithDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:smart@tcp(localhost:3306)/poc")
	if err != nil {
		log.Fatal(" Error occur while create connection with Database.. ")
	}
	/*
		Gorm is an ORM framwork which will work with the relational database.
		>> db.CreateTable(&model.Employee{})
		Above sentence is used for create new Table in database.
		it will create new table with the mentioned model.
	*/
	return db
}

func main() {
	fmt.Println("Gin Goram Example")
	router := gin.Default()
	router.GET("/poc/gin/gorm/list", getEmployeeList)
	router.GET("/poc/gin/gorm/list/:id", getEmployeeById)
	router.POST("/poc/gin/gorm/create", addEmployee)
	router.DELETE("/poc/gin/gorm/remove/:id", removeEmployee)
	router.PUT("/poc/gin/gorm/update/:id", updateEmployee)
	router.Run(":8080")

}

func getEmployeeList(c *gin.Context) {
	db := connectWithDatabase()
	defer db.Close()
	var employees []model.Employee
	selectAllDb := db.Find(&employees)
	if selectAllDb.Error != nil {
		c.JSON(400, model.GetResponse("Error while Get All Employee Details", 400, nil))
		return
	}
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
	var employees []model.Employee
	//Select Data by where condition
	selectQeury := db.Where("id=?", id).Find(&employees)
	if selectQeury.Error != nil {
		c.JSON(400, model.GetResponse("Error while Get All Employee Details", 400, nil))
		return
	}
	c.JSON(200, model.GetResponse("Employee Return SuccessFully", 200, employees))
}

func addEmployee(c *gin.Context) {
	db := connectWithDatabase()
	defer db.Close()
	var e model.Employee
	json.NewDecoder(c.Request.Body).Decode(&e)
	insertDb := db.Create(&e)
	if insertDb.Error != nil {
		c.JSON(400, model.GetResponse("Error while Insert Employees Data", 400, nil))
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
	deleteDB := db.Where("id =?", id).Delete(&model.Employee{})
	if deleteDB.Error != nil {
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
	e := model.Employee{
		Id: id,
	}
	findDb := db.Find(&e)
	if findDb.Error != nil {
		c.JSON(400, model.GetResponse("FindDB ::-- Error while Update Employee", 400, nil))
		return
	}
	json.NewDecoder(c.Request.Body).Decode(&e)
	saveDb := db.Save(&e)
	if saveDb.Error != nil {
		c.JSON(400, model.GetResponse("SaveDB ::-- Error while Update Employee", 400, nil))
		return
	}
	c.JSON(200, model.GetResponse("Employee Updated SuccessFully", 200, nil))
}
