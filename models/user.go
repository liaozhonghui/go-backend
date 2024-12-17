package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    ID       uint   `json:"id" gorm:"primary_key"`
    Username string `json:"username" gorm:"unique;not null"`
    Password string `json:"password" gorm:"not null"`
}

func (user *User) CreateUser(db *gorm.DB) (*User, error) {
    err := db.Create(&user).Error
    if err != nil {
        return nil, err
    }
    return user, nil
}