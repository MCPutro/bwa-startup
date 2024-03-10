package app

import "github.com/gofiber/fiber/v2"

type Server struct {
	*fiber.App
}

func NewServer() *Server {
	return &Server{fiber.New()}

}

func (s *Server) Listen(port string) {
	s.App.Listen(port)
}
