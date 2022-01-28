package user

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(ID int) (models.User, error)
	Update(user models.User) (models.User, error)
	UpdateUser(userUpdate models.User) error
}

type repository struct {
	db *gorm.DB
}

func NewReposiotry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user models.User) (models.User, error) {

	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, err
}

func (r *repository) UpdateUser(userUpdate models.User) error {
	var user models.User
	err := r.db.Model(&user).Updates(userUpdate).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID int) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ? ", ID).Find(&user).Error

	if err != nil {
		return user, err

	}

	return user, nil
}

func (r *repository) Update(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
