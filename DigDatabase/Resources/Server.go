package resource

import (
	service "DigDatabase/Service"
	"net/http"
	"strconv"
)

type Server struct {
	Port    int
	Service service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{
		Service: *service,
		Port:    8080,
	}
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
	mux.HandleFunc("/expert/details", server.Service.GetAllExpertDetails)
	mux.HandleFunc("/user/details", server.Service.GetAllUsers)
	return mux

}
