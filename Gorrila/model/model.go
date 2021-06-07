package model

type Response struct {
	Message string     `json:"message"`
	Status  int        `json:"status"`
	Payload []Employee `json:"payload"`
}

type Employee struct {
	Id         int    `json:"id"`
	FName      string `json:"fName"`
	LName      string `json:"lName"`
	Email      string `json:"email"`
	Position   string `json:"position"`
	Experience int    `json:"experience"`
}

func GetResponse(message string, status int, payload []Employee) *Response {
	return &Response{
		Message: message,
		Status:  status,
		Payload: payload,
	}
}
