package main

import (
	"microservices-demo/internals/api"
	"microservices-demo/internals/repository"
	"microservices-demo/internals/service"
)

func main() {

	repo := repository.NewProductRepository()
	service := service.NewProductService(repo)
	api.StartServer(service)

}
