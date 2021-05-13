package models

import (
	"github.com/jinzhu/gorm"
	"study/db/mysql/src/config"
)

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

func (u *UserModel) LoadByName(name string) *UserModel {
	u.Name = name
	config.DB.Where(u).First(u)
	return u
}

func (u *UserModel) AgeCompare(age int, opt int) CompareFunc {
	return func(db *gorm.DB) *gorm.DB {
		switch opt {
		case GraterThan:
			return db.Where("age>?", age)
		case LessThan:
			return db.Where("age<?", age)
		default:
			return db.Where("age=", age)
		}
	}
}

func (u *UserModel) Filter(f ...func(db *gorm.DB) *gorm.DB) {
	config.DB.Scopes(f...).First(u)
}
