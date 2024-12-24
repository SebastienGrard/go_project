package service

import (
	"context"
	"crud_project/data"
)

type BookService interface {
	Create(ctx context.Context, request data.BookCreateRequest)
	Update(ctx context.Context, request data.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) data.BookResponse
	FindAll(ctx context.Context) []data.BookResponse
}
