package main

import (
	"database/sql"
	"fmt"
	"log"
	"todo-api/config"
	"todo-api/controller"
	"todo-api/repository"
	"todo-api/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// server ini menghubungkan semua komponen
type Server struct {
	uS      service.UserService
	tS      service.TodoService
	engine  *gin.Engine //untuk start engine gin
	portApp string
}

// method untuk memanggil route yang di controller
func (s *Server) initiateRoute() {
	//bisa menambah grouping lagi disini
	routerGroup := s.engine.Group("/api/v1")
	controller.NewUserController(s.uS, routerGroup).Route()
	controller.NewTodoController(s.tS, routerGroup).Route()
}

// func untuk running
func (s *Server) Start() {
	s.initiateRoute()
	s.engine.Run(s.portApp)
}

// constructur, agar dipanggil main.go
func NewServer() *Server {
	//memanggil hasil config .env
	co, _ := config.NewConfig()

	//melakukan koneksi database
	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", co.Host, co.Port, co.User, co.Password, co.Name)

	db, err := sql.Open(co.Driver, urlConnection)
	if err != nil {
		log.Fatal(err)
	}

	portApp := co.AppPort
	userRepo := repository.NewUserRepository(db)
	todoRepo := repository.NewTodoRepository(db)

	userService := service.NewUserService(userRepo)
	todoService := service.NewTodoService(todoRepo)

	//menginject repo ke service
	return &Server{
		uS:      userService,
		tS:      todoService,
		portApp: portApp,
		engine:  gin.Default(),
	}
}
