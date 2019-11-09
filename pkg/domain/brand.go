package domain

import (
	"catalog/pkg/constants"
)

type Brand struct {
	Base
	Name        string
	Description string
	Keywords    string
	Status      constants.Status
}
