package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Base struct {
	gorm.Model
	ExternalId string `gorm:"type:varchar(100);unique_index"`
}

func (entity *Base) BeforeCreate() error{
	externalId := uuid.NewV4()
	entity.ExternalId = externalId.String()
	return nil
}
