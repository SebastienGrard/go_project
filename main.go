package main

import (
	"crud_project/controller"
	"crud_project/dbconfig"
	"crud_project/helper"
	"crud_project/repository"
	"crud_project/router"
	"crud_project/service"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Start server")
	// database
	db := dbconfig.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicError(err)

}
