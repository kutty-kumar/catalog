package builders

import (
	"catalog/pkg/constants"
	"catalog/pkg/domain"
	"catalog/pkg/pb"
)

func BuildBrandAttribute(brand domain.Brand, brandAttributeView pb.BrandAttributeView) domain.BrandAttribute {
	return domain.BrandAttribute{
		BrandId: brand.ID,
		Brand:   brand,
		Key:     brandAttributeView.Key,
		Value:   brandAttributeView.Value,
		Status:  constants.Active,
	}
}

func BuildBrandAttributeDetails(brandAttribute domain.BrandAttribute) pb.BrandAttributeResponseView {
	return pb.BrandAttributeResponseView{
		Key:        brandAttribute.Key,
		Value:      brandAttribute.Value,
		ExternalId: brandAttribute.ExternalId,
	}
}

func MergeBrandAttribute(old domain.BrandAttribute, new pb.UpdateBrandAttributeRequest) domain.BrandAttribute {
	old.Key = new.Payload.Key
	old.Value = new.Payload.Value
	return old
}
