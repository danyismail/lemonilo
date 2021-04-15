package store

import (
	"errors"
	"lemonilo/model"

	"github.com/jinzhu/gorm"
)

type UserStore interface {
	Create(user *model.User) error
	List() (*[]model.User, error)
	Detail(uint) (*model.User, error)
	Update(model.User) error
	Delete(uint) error
	Login(email string) (*model.User, error)
}

type UserConstruct struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserConstruct {
	return &UserConstruct{
		db: db,
	}
}

func (u *UserConstruct) Create(user *model.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

func (u *UserConstruct) Login(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Table("users").Where("email = ?", email).Find(&user).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserConstruct) List() (*[]model.User, error) {
	var user []model.User
	if err := u.db.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserConstruct) Delete(id uint) error {
	var user model.User
	user.ID = id
	result := u.db.Unscoped().Delete(&user)
	if result.RowsAffected < 1 {
		return errors.New("no record found")
	}
	return nil
}

func (u *UserConstruct) Detail(id uint) (*model.User, error) {
	var user model.User
	if err := u.db.Where("id = ?", id).
		Find(&user).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserConstruct) Update(user model.User) error {
	if err := u.db.Save(&user).
		Error; err != nil {
		return err
	}
	return nil
}
