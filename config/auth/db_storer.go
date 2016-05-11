package auth

import (
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"gopkg.in/authboss.v0"
)

type AuthStorer struct {
}

func (s AuthStorer) Create(key string, attr authboss.Attributes) error {
	var user models.User
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s AuthStorer) Put(key string, attr authboss.Attributes) error {
	s.Create(key, attr)
	return nil
}

func (s AuthStorer) Get(key string) (result interface{}, err error) {
	var user models.User
	if err := db.DB.Where("email = ?", key).First(&user).Error; err != nil {
		return nil, authboss.ErrUserNotFound
	}
	return &user, nil
}
