package model

import "gorm.io/gorm"

type Version struct {
	gorm.Model
	SoftName  string
	SoftVersion  string
	SoftPath string
	IsChoose int
}