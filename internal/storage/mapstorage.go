package storage

import (
	"errors"
	"fmt"

	"github.com/RobinSoGood/bookly/internal/domain/models"
	"github.com/RobinSoGood/bookly/internal/logger"
	"github.com/RobinSoGood/bookly/internal/storage/storageerrors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type MapStorage struct {
	stor  map[string]models.User
	bStor map[string]models.Book
}

func New() *MapStorage {
	return &MapStorage{stor: make(map[string]models.User)}
}

func (ms *MapStorage) SaveUser(user models.User) (string, error) {
	log := logger.Get()
	for _, usr := range ms.stor {
		if user.Email == usr.Email {
			return ``, errors.New("user alredy exist")
		}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return ``, err
	}
	user.Password = string(hash)
	UID := uuid.New()
	user.UID = UID
	ms.stor[user.UID.String()] = user
	log.Debug().Any("storage", ms.stor).Msg("check storage")
	return UID.String(), nil
}

func (ms *MapStorage) ValidateUser(user models.UserLogin) (string, error) {
	for key, usr := range ms.stor {
		if user.Email == usr.Email {
			if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(user.Password)); err != nil {
				return ``, errors.New("invalid user password")
			}
			return key, nil
		}
	}
	return ``, errors.New("user no exist")
}

func (ms *MapStorage) SaveBook(book models.Book) (string, error) {
	log := logger.Get()
	for _, b := range ms.bStor {
		if book.Lable == b.Lable && book.Author == b.Author {
			return ``, storageerrors.ErrBookAlreadyExist
		}
	}
	bID := uuid.New()
	book.BID = bID
	ms.bStor[book.BID.String()] = book
	log.Debug().Any("book storage", ms.bStor).Msg("check storage")
	return bID.String(), nil
}

func (ms *MapStorage) GetBooks() ([]models.Book, error) {
	if len(ms.bStor) == 0 {
		return nil, storageerrors.ErrEmptyStorage
	}
	var books []models.Book
	for _, book := range ms.bStor {
		books = append(books, book)
	}
	return books, nil
}

func (ms *MapStorage) GetBook(bid string) (models.Book, error) {
	book, ok := ms.bStor[bid]
	if !ok {
		return models.Book{}, storageerrors.ErrBookNoFound
	}
	return book, nil
}

func (ms *MapStorage) DeleteBook(bid string) error {
	_, ok := ms.bStor[bid]
	if !ok {
		return fmt.Errorf("book not found")
	}
	delete(ms.bStor, bid)
	return nil
}
