package internal

import "fmt"

type Docker struct {
}

// Запуск контейнера
func (d Docker) StartDocker() {
	fmt.Println("Docker started")
}

type Server struct {
}

// Запуск сервера
func (s Server) StartServer() {
	fmt.Println("Server started")
}

type Database struct {
}

// Запуск базы данных
func (s Database) StartDatabase() {
	fmt.Println("Database started")
}

// Фасад
type Facade struct {
	docker   *Docker
	server   *Server
	database *Database
}

// Конструктор фасада
func NewService() *Facade {
	return &Facade{
		docker:   &Docker{},
		server:   &Server{},
		database: &Database{},
	}
}

// Запуск всех модулей
func (f *Facade) StartService() {
	f.docker.StartDocker()
	f.server.StartServer()
	f.database.StartDatabase()
	fmt.Println("Service started")
}

func FacadePattern() {
	service := NewService()
	service.StartService()
}
