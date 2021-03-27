package usecase

import (
	"context"

	"github.com/gmhafiz/go8/internal/domain/book"
	"github.com/gmhafiz/go8/internal/models"
)

type BookUseCase struct {
	bookRepo book.Repository
}

func New(bookRepo book.Repository) *BookUseCase {
	return &BookUseCase{
		bookRepo: bookRepo,
	}
}

func (u *BookUseCase) Create(ctx context.Context, book *models.Book) (*models.Book, error) {
	bookID, err := u.bookRepo.Create(ctx, book)
	if err != nil {
		return nil, err
	}
	bookFound, err := u.bookRepo.Read(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return bookFound, err
}

func (u *BookUseCase) List(ctx context.Context, f *book.Filter) ([]*models.Book, error) {
	return u.bookRepo.List(ctx, f)
}

func (u *BookUseCase) Read(ctx context.Context, bookID int64) (*models.Book, error) {
	return u.bookRepo.Read(ctx, bookID)
}

func (u *BookUseCase) Update(ctx context.Context, book *models.Book) (*models.Book, error) {
	err := u.bookRepo.Update(ctx, book)
	if err != nil {
		return nil, err
	}
	return u.bookRepo.Read(ctx, book.BookID)
}

func (u *BookUseCase) Delete(ctx context.Context, bookID int64) error {
	return u.bookRepo.Delete(ctx, bookID)
}

func (u *BookUseCase) Search(ctx context.Context, req *book.Filter) ([]*models.Book, error) {
	return u.bookRepo.Search(ctx, req)
}
