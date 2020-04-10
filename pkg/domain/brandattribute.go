package domain

import "catalog/pkg/constants"

type BrandAttribute struct {
	Base
	BrandId uint
	Brand Brand `gorm:"foreignkey:BrandId"`
	Key string
	Value string
	constants.Status
}
