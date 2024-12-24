package repository

import (
	"context"
	"crud_project/helper"
	"crud_project/model"
	"database/sql"
	"errors"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

// Delete implements BookRepository
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from book where id =$1"
	_, errExec := tx.ExecContext(ctx, SQL, bookId)
	helper.PanicError(errExec)
}

// FindAll implements BookRepository
func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	tx, err := b.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id,name from book"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicError(errQuery)
	defer result.Close()

	var books []model.Book

	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicError(err)

		books = append(books, book)
	}

	return books
}

// FindById implements BookRepository
func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (model.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id,name from book where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicError(errQuery)
	defer result.Close()

	book := model.Book{}

	if result.Next() {
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicError(err)
		return book, nil
	} else {
		return book, errors.New("book id not found")
	}
}

// Save implements BookRepository
func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into book(name) values ($1)"
	_, err = tx.ExecContext(ctx, SQL, book.Name)
	helper.PanicError(err)
}

// Update implements BookRepository
func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update book set name=$1 where id=$2"
	_, err = tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicError(err)
}
