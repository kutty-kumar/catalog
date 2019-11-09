package builders

import (
	"catalog/pkg/constants"
	"catalog/pkg/domain"
	"catalog/pkg/pb"
)

func BuildBrand(brandDetails *pb.BrandDetails) domain.Brand {
	return domain.Brand{Name: brandDetails.GetName(), Description: brandDetails.GetDescription(), Keywords: brandDetails.GetKeywords(), Status: constants.Status(brandDetails.Status)}
}

func BuildBrandDetails(brand *domain.Brand) *pb.BrandDetails {
	status := int32(brand.Status)
	return &pb.BrandDetails{Name: brand.Name, Keywords: brand.Keywords, Description: brand.Description, Status: status}
}


func MergeBrand(brandOld *domain.Brand, updateReq *pb.BrandDetails) {
	if &updateReq.Status != nil {
		brandOld.Status = constants.Status(updateReq.Status)
	}
	if &brandOld.Name != nil {
		brandOld.Name = updateReq.Name
	}
	if &brandOld.Keywords != nil {
		brandOld.Keywords = updateReq.Keywords
	}
	if &brandOld.Description != nil {
		brandOld.Description = updateReq.Description
	}
}