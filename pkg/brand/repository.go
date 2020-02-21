package brand

import "catalog/pkg/domain"

type Repository interface {
	Create(brand *domain.Brand)
	Update(brand *domain.Brand)
	Delete(brand *domain.Brand)
	GetByExternalId(externalId string) domain.Brand
	MultiGetByExternalId(externalIds []string) ([]domain.Brand, error)
	GetAttributes(externalId string) ([]domain.BrandAttribute, error)
}
