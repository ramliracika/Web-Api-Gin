package service

import (
	"github.com/ramliracika/Web-Api-Go/laptop"
	"github.com/ramliracika/Web-Api-Go/repository"
)

type Service interface {
	FindAll() ([]laptop.Gaming, error)
	FindById(id int) (laptop.Gaming, error)
	Create(laptop laptop.Gaming) (laptop.Gaming, error)
	Update(laptop laptop.Gaming, id int) (laptop.Gaming, error)
	Delete(id int) (laptop.Gaming, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]laptop.Gaming, error) {
	laptop, err := s.repository.FindAll()

	return laptop, err

	//return s.repository.FindAll()
}

func (s *service) FindById(id int) (laptop.Gaming, error) {
	laptop, err := s.repository.FindById(id)

	return laptop, err
}

func (s *service) Create(laptop laptop.Gaming) (laptop.Gaming, error) {
	newLaptop, err := s.repository.Create(laptop)

	return newLaptop, err
}

func (s *service) Update(laptop laptop.Gaming, id int) (laptop.Gaming, error) {
	updateLaptop, err := s.repository.Update(laptop, id)

	return updateLaptop, err
}

func (s *service) Delete(id int) (laptop.Gaming, error) {
	delete, err := s.repository.Delete(id)

	return delete, err
}
