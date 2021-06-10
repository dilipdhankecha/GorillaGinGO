package config

import (
	model "Dig/ConfigModel/Model"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	config DataBaseConfig
	repo   PersonRepository
	Port   int
}

type PersonRepository struct {
	database *sql.DB
}

func (repo *PersonRepository) getAllEmployee() []model.Employee {
	db := repo.database
	query := "SELECT * FROM employee"
	result, err := db.Query(query)
	if err != nil {
		return nil
	}
	var employees []model.Employee
	for result.Next() {
		var e model.Employee
		err := result.Scan(&e.Id, &e.FName, &e.LName, &e.Email, &e.Position, &e.Experience)
		if err != nil {
			return nil
		}
		employees = append(employees, e)
	}
	defer result.Close()
	return employees
}

type DataBaseConfig struct {
	DialectName  string
	DataBaseName string
	User         string
	Password     string
	Port         int
}

//"root:smart@tcp(localhost:3306)/poc"

func NewDataBaseConfig() *DataBaseConfig {
	return &DataBaseConfig{
		DialectName:  "mysql",
		DataBaseName: "poc",
		User:         "root",
		Password:     "smart",
		Port:         3306,
	}
}

func ConnectDatabase(databaseConfig *DataBaseConfig) *sql.DB {
	connection := databaseConfig.User + ":" + databaseConfig.Password + "@tcp(localhost:" + strconv.Itoa(databaseConfig.Port) + ")/" + databaseConfig.DataBaseName
	db, err := sql.Open(databaseConfig.DialectName, connection)
	if err != nil {
		log.Println("=== Error Occure while connect Database ====")
		log.Panic(err)
	}
	return db
}

func (server *Server) findPeople(writer http.ResponseWriter, request *http.Request) {
	people := server.repo.getAllEmployee()
	bytes, _ := json.Marshal(people)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}

func NewServer(config *DataBaseConfig, repo *PersonRepository) *Server {
	return &Server{
		config: *config,
		repo:   *repo,
		Port:   8080,
	}
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}

func (server *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(server.Port),
		Handler: server.HandlerFunc(),
	}

	httpServer.ListenAndServe()
}

func (server *Server) HandlerFunc() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/people", server.findPeople)
	return mux

}
