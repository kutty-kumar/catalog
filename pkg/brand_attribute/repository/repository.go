package repository

import "catalog/pkg/domain"

type Repository interface {
	Create(brand *domain.Brand, brandAttribute *domain.BrandAttribute)
	Update(brand *domain.Brand, brandAttribute *domain.BrandAttribute)
	Delete(brand *domain.Brand, brandAttribute *domain.BrandAttribute)
	GetByExternalId(externalId string) (*domain.BrandAttribute, error)
	MultiGetByExternalId(externalIds []string) ([]domain.BrandAttribute, error)
}
