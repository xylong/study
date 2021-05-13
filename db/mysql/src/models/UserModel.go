package models

import "study/db/mysql/src/config"

type UserModel struct {
	ID   int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) LoadByID(id int) *UserModel {
	config.DB.First(u, id)
	return u
}
