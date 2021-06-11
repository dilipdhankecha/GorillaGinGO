package core

import (
	repo "DigDatabase/Repository"
	"encoding/json"
	"net/http"
)

type Service struct {
	ExpertRe repo.ExpertRepo
	UserRe   repo.UserRepo
}

func NewService(expertRepo *repo.ExpertRepo, userRepo *repo.UserRepo) *Service {
	return &Service{
		ExpertRe: *expertRepo,
		UserRe:   *userRepo,
	}
}

func (service *Service) GetAllExpertDetails(response http.ResponseWriter, request *http.Request) {
	experts := service.ExpertRe.GetAllExpertDetails()
	bytes, _ := json.Marshal(&experts)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(bytes)

}

func (service *Service) GetAllUsers(response http.ResponseWriter, request *http.Request) {
	users := service.UserRe.GetAllUsers()
	bytes, _ := json.Marshal(&users)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(bytes)

}
