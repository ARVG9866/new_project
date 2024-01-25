package service

import "github.com/ARVG9866/new_project/internal/storage"

type IService interface {
}

type service struct {
	storage storage.IStorage
}

func NewService(storage storage.IStorage) IService {
	return &service{
		storage: storage,
	}
}
