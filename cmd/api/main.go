package main

import (
	"log"
	"net/http"

	"github.com/gguedes14/mundo-invest/internal/controller"
	"github.com/gguedes14/mundo-invest/internal/handler"
	"github.com/gguedes14/mundo-invest/internal/infrastructure/database"
	"github.com/gguedes14/mundo-invest/internal/repository"
	"github.com/gguedes14/mundo-invest/internal/routes"
	"github.com/gguedes14/mundo-invest/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	db := database.ConnectDb()

	repo := repository.NewClientRepository(db)

	svc := service.NewService(repo)

	ctrl := controller.NewController(svc)

	h := handler.NewHandler(ctrl)

	router := chi.NewRouter()

	routes.ClientRoutes(router, h)

	log.Println("🚀 Server running on :3000")

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
