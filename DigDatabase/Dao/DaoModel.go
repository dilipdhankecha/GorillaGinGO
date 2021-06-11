package dao

type Expert struct {
	Id        int    `json:"id"`
	FirstName string `json:"fName"`
	LastName  string `json:"lName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type User struct {
	Id       int    `json:"id"`
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	Country  string `json:"country"`
}
