package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username" gorm:"unique"`
	Active    bool      `json:"active"`
	Password  string    `json:"password"`
	UpdatedAt string    `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
}

func Get(db *gorm.DB, userID uint) (*User, error) {
	var user User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Create(db *gorm.DB) (uint, error) {
	result := db.Create(u)
	if result.Error != nil {
		return 0, result.Error
	}
	return u.ID, nil
}

func (u *User) Update(db *gorm.DB) error {
	err := db.Model(&User{}).Where("id = ?", u.ID).Updates(
		User{
			Email:    u.Email,
			Username: u.Username,
			Active:   u.Active,
			Password: u.Password,
		}).Error

	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(db *gorm.DB, ID uint) error {
	result := db.Delete(&User{}, ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
