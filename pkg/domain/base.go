package domain

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Base struct {
	gorm.Model
	ExternalId string `gorm:"type:varchar(100);unique_index"`
}

func (entity *Base) BeforeCreate() error{
	externalId, err := uuid.NewV4()
	if err != nil {
		return errors.New(fmt.Sprintf("An error occurred while generating external id for entity %v", entity))
	}
	entity.ExternalId = externalId.String()
	return nil
}
