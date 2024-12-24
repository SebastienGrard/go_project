package controller

import (
	"crud_project/data"
	"crud_project/helper"
	"crud_project/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (controller *BookController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookCreateRequest := data.BookCreateRequest{}
	helper.ReadRequestBody(requests, &bookCreateRequest)

	controller.BookService.Create(requests.Context(), bookCreateRequest)
	webResponse := data.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookUpdateRequest := data.BookUpdateRequest{}
	helper.ReadRequestBody(requests, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicError(err)
	bookUpdateRequest.Id = id

	controller.BookService.Update(requests.Context(), bookUpdateRequest)
	webResponse := data.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *BookController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicError(err)

	controller.BookService.Delete(requests.Context(), id)
	webResponse := data.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *BookController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.BookService.FindAll(requests.Context())
	webResponse := data.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicError(err)

	result := controller.BookService.FindById(requests.Context(), id)
	webResponse := data.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)

}
