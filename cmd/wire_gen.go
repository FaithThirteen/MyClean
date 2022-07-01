// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fhz/api"
	"fhz/api/handlers"
	"fhz/app"
	"fhz/repo"
	"fhz/service"
)

// Injectors from wire.go:

func InitServer() *app.Server {
	engine := app.NewGinEngine()
	db := app.InitDB()
	iArticleRepo := repo.NewArticleRepository(db)
	iArticleService := service.NewArticleService(iArticleRepo)
	articleHandel := handlers.NewArticleHandel(iArticleService)
	router := api.NewRouter(articleHandel)
	server := app.NewServer(engine, router)
	return server
}