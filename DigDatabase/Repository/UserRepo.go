package expertRepo

import (
	dao "DigDatabase/Dao"
	"database/sql"
	"log"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(database *sql.DB) *UserRepo {
	return &UserRepo{db: database}
}

func (e *UserRepo) GetAllUsers() []dao.User {
	db := e.db
	res, err := db.Query("Select * From user")
	if err != nil {
		log.Println("Error Occur while execute Query for select all records from database..")
	}
	var userDetails []dao.User
	for res.Next() {
		var u dao.User
		err := res.Scan(&u.Id, &u.FullName, &u.Country, &u.Age)
		if err != nil {
			log.Println("Error occur while scane user data")
			log.Panic(err)
		}
		userDetails = append(userDetails, u)
	}
	defer res.Close()
	return userDetails
}
