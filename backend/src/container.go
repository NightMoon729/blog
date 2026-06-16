package main

import (
	"blog/backend/src/Service"
	"blog/backend/src/handler"
	"blog/backend/src/repository"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Container struct {
	SolutionHandler *handler.SolutionHandler
}

func NewContainer(db *sqlx.DB) *Container {
	solutionRepo := repository.NewSolutionRepo(db)
	solutionSvc := Service.NewSolutionService(solutionRepo)
	solutionHandler := handler.NewSolutionHandler(solutionSvc)

	return &Container{
		SolutionHandler: solutionHandler,
	}
}

func registerRoutes(c *Container) {
	http.HandleFunc("/api/solutions", c.SolutionHandler.GetSolutions)
}
