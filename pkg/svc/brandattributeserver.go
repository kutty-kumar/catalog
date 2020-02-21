package svc

import (
	"catalog/pkg/brand_attribute/repository"
	"catalog/pkg/builders"
	"catalog/pkg/constants"
	"catalog/pkg/domain"
	"catalog/pkg/pb"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	version = "0.0.1"
)

type BrandAttributeService struct {
	repository   repository.Repository
	brandService BrandService
}

func NewBrandAttributeService(database *gorm.DB, brandService BrandService) (BrandAttributeService, error) {
	var rep repository.Repository
	rep = &repository.BrandAttributeGORM{Db: database}
	return BrandAttributeService{repository: rep, brandService: brandService}, nil
}

func (service BrandAttributeService) validateBrand(externalId string) (*domain.Brand, error) {
	brand, err := service.brandService.GetBrand(externalId)
	if err != nil {
		return nil, err
	}
	return brand, nil
}

func (service BrandAttributeService) getBrandAttribute(externalId string) (*domain.BrandAttribute, error) {
	brandAttribute, err := service.repository.GetByExternalId(externalId)
	if err != nil {
		return nil, err
	}
	return brandAttribute, nil
}

func (service BrandAttributeService) CreateBrandAttribute(_ context.Context, request *pb.CreateBrandAttributeRequest) (*pb.CreateBrandAttributeResponse, error) {
	if brand, err := service.validateBrand(request.BrandId); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brand %v not found.", request.BrandId))
	} else {
		brandAttribute := builders.BuildBrandAttribute(*brand, *request.Payload)
		service.repository.Create(brand, &brandAttribute)
		brandAttributeDetails := builders.BuildBrandAttributeDetails(brandAttribute)
		return &pb.CreateBrandAttributeResponse{Result: &brandAttributeDetails}, nil
	}
}

func (service BrandAttributeService) UpdateBrandAttribute(_ context.Context, request *pb.UpdateBrandAttributeRequest) (*pb.UpdateBrandAttributeResponse, error) {
	if brand, err := service.validateBrand(request.BrandId); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brand %v not found.", request.BrandId))
	} else {
		brandAttribute, err := service.repository.GetByExternalId(request.BrandAttributeId)
		if err != nil {
			return nil, err
		}
		updatedBrandAttribute := builders.MergeBrandAttribute(*brandAttribute, *request)
		service.repository.Update(brand, &updatedBrandAttribute)
		brandAttributeDetails := builders.BuildBrandAttributeDetails(updatedBrandAttribute)
		return &pb.UpdateBrandAttributeResponse{Result: &brandAttributeDetails}, nil
	}
}

func (service BrandAttributeService) DeleteBrandAttribute(_ context.Context, request *pb.DeleteBrandAttributeRequest) (*pb.DeleteBrandAttributeResponse, error) {
	if brand, err := service.validateBrand(request.BrandId); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brand %v not found.", request.BrandId))
	} else {
		brandAttribute, err := service.repository.GetByExternalId(request.BrandAttributeId)
		if err != nil {
			return nil, err
		}
		brandAttribute.Status = constants.Inactive
		service.repository.Update(brand, brandAttribute)
		return &pb.DeleteBrandAttributeResponse{}, nil
	}
}

func (service BrandAttributeService) GetBrandAttributeById(_ context.Context, request *pb.GetBrandAttributeByIdRequest) (*pb.GetBrandAttributeByIdResponse, error) {
	if _, err := service.validateBrand(request.BrandId); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brand %v not found.", request.BrandId))
	} else {
		brandAttribute, err := service.repository.GetByExternalId(request.BrandAttributeId)
		if err != nil {
			return nil, err
		}
		brandAttributeDetails := builders.BuildBrandAttributeDetails(*brandAttribute)
		return &pb.GetBrandAttributeByIdResponse{Result: &brandAttributeDetails}, nil
	}
}

func (service BrandAttributeService) MultiGetBrandAttributes(_ context.Context, request *pb.MultiGetBrandAttributeRequest) (*pb.MultiGetBrandAttributeResponse, error) {
	if _, err := service.validateBrand(request.BrandId); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Brand %v not found.", request.BrandId))
	} else {
		brandAttributes, err := service.repository.MultiGetByExternalId(request.ExternalIds)
		if err != nil {
			return nil, err
		}
		var result []*pb.BrandAttributeResponseView
		for _, brandAttribute := range brandAttributes {
			brandAttributeResponseView := builders.BuildBrandAttributeDetails(brandAttribute)
			result = append(result, &brandAttributeResponseView)
		}
		return &pb.MultiGetBrandAttributeResponse{Results: result}, nil
	}
}
