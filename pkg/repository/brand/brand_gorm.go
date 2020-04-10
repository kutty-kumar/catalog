package brand

import (
	"catalog/pkg/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BrandGORMRepository struct {
	Db *gorm.DB
}

func (orm *BrandGORMRepository) getBrand(externalId string) domain.Brand {
	base := domain.Base{ExternalId: externalId}
	return domain.Brand{Base: base}
}

func (orm *BrandGORMRepository) Create(brand *domain.Brand) {
	orm.Db.Save(&brand)
}

func (orm BrandGORMRepository) Update(brand *domain.Brand) {
	orm.Db.Model(&brand).Update(&brand)
}

func (orm BrandGORMRepository) Delete(brand *domain.Brand) {
	brand.Status = 1
	orm.Db.Model(&brand).Update(&brand)
}

func (orm BrandGORMRepository) GetByExternalId(externalId string)  (*domain.Brand, error) {
	brand := orm.getBrand(externalId)
	orm.Db.Where(&brand).First(&brand)
	if brand.ID == 0 {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brand %v not found.", externalId))
	}
	return &brand, nil
}

func (orm BrandGORMRepository) MultiGetByExternalId(externalIds []string) ([]domain.Brand, error) {
	var brands []domain.Brand
	orm.Db.Where("external_id IN (?)", externalIds).Find(&brands)
	if brands == nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brands with external_ids %v not found", externalIds))
	}
	return brands, nil
}

func (orm BrandGORMRepository) GetAttributes(externalId string) ([]domain.BrandAttribute, error) {
	var brandAttributes []domain.BrandAttribute
	brand := orm.getBrand(externalId)
	orm.Db.Model(&brand).Related(&brandAttributes)
	if brandAttributes == nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brand Attributes with ids %v not found", externalId))
	}
	return brandAttributes, nil
}