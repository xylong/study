package models

import "github.com/jinzhu/gorm"

type CompareFunc func(db *gorm.DB) *gorm.DB

const (
	GraterThan = 10 // >
	LessThan   = 20 // <
	Equal      = 30 // =
)
