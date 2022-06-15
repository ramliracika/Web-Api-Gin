package repository

import (
	"github.com/ramliracika/Web-Api-Go/laptop"
	"gorm.io/gorm"
)

//declare repository contract
type Repository interface {
	FindAll() ([]laptop.Gaming, error)
	FindById(id int) (laptop.Gaming, error)
	Create(laptop laptop.Gaming) (laptop.Gaming, error)
	Update(laptop laptop.Gaming, id int) (laptop.Gaming, error)
	Delete(id int) (laptop.Gaming, error)
}

// declare struct to connect database
type repository struct {
	db *gorm.DB
}

// delcare repository to hit on main
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// declare all function on contract

func (r *repository) FindAll() ([]laptop.Gaming, error) {
	var laptop []laptop.Gaming

	err := r.db.Find(&laptop).Error // select * from table

	return laptop, err
}

func (r *repository) FindById(id int) (laptop.Gaming, error) {
	var laptop laptop.Gaming

	err := r.db.Where("id = ?", id).Find(&laptop).Error

	return laptop, err

}
func (r *repository) Create(laptop laptop.Gaming) (laptop.Gaming, error) {
	err := r.db.Create(&laptop).Error

	return laptop, err
}

func (r *repository) Delete(id int) (laptop.Gaming, error) {
	var laptop laptop.Gaming
	err := r.db.Where("id = ?", id).Delete(&laptop).Error

	return laptop, err
}

func (r *repository) Update(laptop laptop.Gaming, id int) (laptop.Gaming, error) {

	err := r.db.Model(&laptop).Where("id=?", id).Updates(&laptop).Error

	return laptop, err

}
