package doctor

import (
	"catalog/pkg/domain"
)

type Repository interface {
	Create(request *domain.Doctor) *domain.Doctor
	GetByExternalId(externalId string) (*domain.Doctor, error)
	UpdateDoctor(request *domain.Doctor) *domain.Doctor
	MultiGet(doctorIds []string) ([]*domain.Doctor, error)
}