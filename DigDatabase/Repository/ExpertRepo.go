package expertRepo

import (
	dao "DigDatabase/Dao"
	"database/sql"
	"log"
)

type ExpertRepo struct {
	db *sql.DB
}

func NewExpertRepo(database *sql.DB) *ExpertRepo {
	return &ExpertRepo{db: database}
}

func (e *ExpertRepo) GetAllExpertDetails() []dao.Expert {
	db := e.db
	res, err := db.Query("Select * From expert")
	if err != nil {
		log.Println("Error Occur while execute Query for select all records from database..")
	}
	var expertDetails []dao.Expert
	for res.Next() {
		var expert dao.Expert
		err := res.Scan(&expert.Id, &expert.FirstName, &expert.LastName, &expert.Email, &expert.Age)
		if err != nil {
			log.Println("Error occur while scane expert data")
			log.Panic(err)
		}
		expertDetails = append(expertDetails, expert)
	}
	defer res.Close()
	return expertDetails
}
