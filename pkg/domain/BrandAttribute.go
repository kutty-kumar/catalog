package domain

import "catalog/pkg/constants"

type BrandAttribute struct {
	Base
	BrandId int
	Brand Brand
	Key string
	Value string
	constants.Status
}
