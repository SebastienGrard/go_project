package router

import (
	"crud_project/controller"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController *controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	router.GET("/api/book", bookController.FindAll)
	router.GET("/api/book/:bookId", bookController.FindById)
	router.POST("/api/book", bookController.Create)
	router.PATCH("/api/book/:bookId", bookController.Update)
	router.DELETE("/api/book/:bookId", bookController.Delete)

	return router
}
