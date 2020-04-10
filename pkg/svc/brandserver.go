package svc

import (
	"catalog/pkg/builders"
	"catalog/pkg/domain"
	"catalog/pkg/pb"
	bRepo "catalog/pkg/repository/brand"
	"catalog/pkg/utils"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BrandService struct {
	repository bRepo.Repository
}

func NewBrandService(database *gorm.DB) (BrandService, error) {
	var rep bRepo.Repository
	rep = &bRepo.BrandGORMRepository{Db: database}
	return BrandService{repository: rep}, nil
}

func (service BrandService) CreateBrand(_ context.Context, req *pb.CreateBrandRequest) (*pb.CreateBrandResponse, error) {
	brand := builders.BuildBrand(req.Payload)
	service.repository.Create(&brand)
	return &pb.CreateBrandResponse{Brand: req.Payload, ExternalId: brand.ExternalId}, nil
}

func (service BrandService) GetBrandById(_ context.Context, req *pb.GetBrandByIdRequest) (*pb.GetBrandByIdResponse, error) {
	if err := utils.HandleEmptyRequest(&req.ExternalId, "Brand Id"); err != nil {
		return nil, err
	}
	brand, err := service.repository.GetByExternalId(req.ExternalId)
	if err != nil{
		return nil, err
	}
	return &pb.GetBrandByIdResponse{Brand: builders.BuildBrandDetails(brand)}, nil
}

func (service BrandService) DeleteBrand(_ context.Context, req *pb.DeleteBrandRequest) (*pb.DeleteBrandResponse, error) {
	if err := utils.HandleEmptyRequest(&req.ExternalId, "Brand Id"); err != nil {
		return nil, err
	}
	brand, err  := service.repository.GetByExternalId(req.ExternalId)
	if err == nil {
		service.repository.Delete(brand)
		return &pb.DeleteBrandResponse{}, nil
	}
	return nil, status.Error(codes.Internal, fmt.Sprintf("Brand %v not found.", req.ExternalId))
}

func (service BrandService) UpdateBrand(_ context.Context, req *pb.UpdateBrandRequest) (*pb.UpdateBrandResponse, error) {
	if err := utils.HandleEmptyRequest(&req.ExternalId, "Brand Id"); err != nil {
		return nil, err
	}
	brand, err := service.repository.GetByExternalId(req.ExternalId)
	if err != nil {
		return nil, err
	}
	builders.MergeBrand(brand, req.Payload)
	service.repository.Update(brand)
	return &pb.UpdateBrandResponse{Brand: builders.BuildBrandDetails(brand), ExternalId: brand.ExternalId}, nil
}

func (service BrandService) MultiGetBrands(_ context.Context, req *pb.MultiGetBrandsRequest) (*pb.MultiGetBrandsResponse, error) {
	brands, err := service.repository.MultiGetByExternalId(req.ExternalIds)
	if err != nil {
		return nil, err
	}
	var result []*pb.GetBrandByIdResponse
	for _, brand := range brands {
		result = append(result, &pb.GetBrandByIdResponse{Brand: builders.BuildBrandDetails(&brand), ExternalId: brand.ExternalId})
	}
	return &pb.MultiGetBrandsResponse{Result: result}, nil
}

func (service BrandService) GetBrandAttributes(_ context.Context, req *pb.GetAttributesForBrandRequest) (*pb.GetAttributesForBrandResponse, error) {
	brandAttributes, err := service.repository.GetAttributes(req.BrandId)
	if err != nil {
		return nil, err
	}
	var result pb.GetAttributesForBrandResponse
	var attributesView []*pb.BrandAttributeResponseView
	for _, brandAttribute := range brandAttributes {
		attributesView = append(attributesView, &pb.BrandAttributeResponseView{ExternalId: brandAttribute.ExternalId, Key: brandAttribute.Key, Value: brandAttribute.Value})
	}
	result.Results = attributesView
	return &result, nil
}

func (service BrandService) GetBrand(externalId string) (*domain.Brand, error) {
	return service.repository.GetByExternalId(externalId)
}
