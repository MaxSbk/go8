package server

import (
	bookHandler "github.com/gmhafiz/go8/internal/domain/book/handler/http"
	bookRepo "github.com/gmhafiz/go8/internal/domain/book/repository/mysql"
	bookUseCase "github.com/gmhafiz/go8/internal/domain/book/usecase"
	healthHandler "github.com/gmhafiz/go8/internal/domain/health/handler/http"
	healthRepo "github.com/gmhafiz/go8/internal/domain/health/repository/mysql"
	healthUseCase "github.com/gmhafiz/go8/internal/domain/health/usecase"
)

func (s *Server) initDomains() {
	s.initHealth()
	s.initBook()
}

func (s *Server) initHealth() {
	newHealthRepo := healthRepo.New(s.DB())
	newHealthUseCase := healthUseCase.New(newHealthRepo)
	healthHandler.RegisterHTTPEndPoints(s.router, newHealthUseCase)
}

func (s *Server) initBook() {
	newBookRepo := bookRepo.New(s.DB())
	newBookUseCase := bookUseCase.New(newBookRepo)
	bookHandler.RegisterHTTPEndPoints(s.router, newBookUseCase)
}
