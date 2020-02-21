package repository

import (
	"catalog/pkg/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BrandGORM struct {
	Db *gorm.DB
}

func (orm *BrandGORM) getBrand(externalId string) domain.Brand {
	base := domain.Base{ExternalId: externalId}
	return domain.Brand{Base: base}
}

func (orm *BrandGORM) Create(brand *domain.Brand) {
	orm.Db.Save(&brand)
}

func (orm BrandGORM) Update(brand *domain.Brand) {
	orm.Db.Model(&brand).Update(&brand)
}

func (orm BrandGORM) Delete(brand *domain.Brand) {
	brand.Status = 1
	orm.Db.Model(&brand).Update(&brand)
}

func (orm BrandGORM) GetByExternalId(externalId string) domain.Brand {
	brand := orm.getBrand(externalId)
	orm.Db.First(&brand)
	return brand
}

func (orm BrandGORM) MultiGetByExternalId(externalIds []string) ([]domain.Brand, error) {
	var brands []domain.Brand
	orm.Db.Where("external_id IN (?)", externalIds).Find(&brands)
	if brands == nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brands with external_ids %v not found", externalIds))
	}
	return brands, nil
}

func (orm BrandGORM) GetAttributes(externalId string) ([]domain.BrandAttribute, error) {
	var brandAttributes []domain.BrandAttribute
	brand := orm.getBrand(externalId)
	orm.Db.Model(&brand).Related(&brandAttributes)
	if brandAttributes == nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brand Attributes with ids %v not found", externalId))
	}
	return brandAttributes, nil
}