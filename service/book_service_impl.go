package service

import (
	"context"
	"crud_project/data"
	"crud_project/helper"
	"crud_project/model"
	"crud_project/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

// Create implements BookService
func (b *BookServiceImpl) Create(ctx context.Context, request data.BookCreateRequest) {
	book := model.Book{
		Name: request.Name,
	}
	b.BookRepository.Save(ctx, book)
}

// Delete implements BookService
func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicError(err)
	b.BookRepository.Delete(ctx, book.Id)
}

// FindAll implements BookService
func (b *BookServiceImpl) FindAll(ctx context.Context) []data.BookResponse {
	books := b.BookRepository.FindAll(ctx)

	var bookResp []data.BookResponse

	for _, value := range books {
		book := data.BookResponse{Id: value.Id, Name: value.Name}
		bookResp = append(bookResp, book)
	}
	return bookResp

}

// FindById implements BookService
func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) data.BookResponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicError(err)
	return data.BookResponse(book)
}

// Update implements BookService
func (b *BookServiceImpl) Update(ctx context.Context, request data.BookUpdateRequest) {
	book, err := b.BookRepository.FindById(ctx, request.Id)
	helper.PanicError(err)

	book.Name = request.Name
	b.BookRepository.Update(ctx, book)
}
