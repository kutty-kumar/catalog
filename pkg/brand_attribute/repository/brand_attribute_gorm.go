package repository

import (
	"catalog/pkg/constants"
	"catalog/pkg/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BrandAttributeGORM struct {
	Db *gorm.DB
}

func (orm BrandAttributeGORM) getBrandAttribute(externalId string) domain.BrandAttribute {
	base := domain.Base{ExternalId: externalId}
	return domain.BrandAttribute{Base: base}
}

func (orm BrandAttributeGORM) Create(brand *domain.Brand, brandAttribute *domain.BrandAttribute){
	brandAttribute.Brand = *brand
	orm.Db.Model(&brandAttribute).Save(&brandAttribute)
}

func (orm BrandAttributeGORM) Update(brand *domain.Brand, brandAttribute *domain.BrandAttribute){
	orm.Db.Model(&brandAttribute).Updates(&brandAttribute)
}

func (orm BrandAttributeGORM) Delete(brand *domain.Brand, brandAttribute *domain.BrandAttribute){
	brandAttribute.Status = constants.Inactive
	orm.Db.Model(&brandAttribute).Updates(&brandAttribute)
}

func (orm BrandAttributeGORM) GetByExternalId(externalId string) (*domain.BrandAttribute, error) {
	brandAttribute := orm.getBrandAttribute(externalId)
	orm.Db.First(&brandAttribute)
	if brandAttribute.ID == 0 {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("Brand Attribute with Id %v not found", externalId))
	}
	return &brandAttribute, nil
}

func (orm BrandAttributeGORM) MultiGetByExternalId(externalIds []string) ([]domain.BrandAttribute, error){
	var brandAttributes []domain.BrandAttribute
	orm.Db.Where("external_id IN (?)", externalIds).Find(&brandAttributes)
	if brandAttributes == nil {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("Brand Attributes %v not found", externalIds))
	}
	return brandAttributes, nil
}
