package repository

import (
	"catalog/pkg/domain"
	"github.com/jinzhu/gorm"
)

type BrandGORM struct {
	Db *gorm.DB
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
	base := domain.Base{ExternalId: externalId}
	brand := domain.Brand{Base: base}
	orm.Db.First(&brand)
	return brand
}

func (orm BrandGORM) MultiGetByExternalId(externalIds []string) []domain.Brand {
	var brands []domain.Brand
	orm.Db.Where("external_id IN (?)", externalIds).Find(&brands)
	return brands
}
