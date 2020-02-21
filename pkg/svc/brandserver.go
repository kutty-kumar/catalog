package svc

import (
	b "catalog/pkg/brand"
	"catalog/pkg/brand/repository"
	"catalog/pkg/builders"
	"catalog/pkg/pb"
	"catalog/pkg/utils"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	version = "0.0.1"
)

type BrandService struct {
	repository b.Repository
}

func NewBrandService(database *gorm.DB) (BrandService, error) {
	var rep b.Repository
	rep = &repository.BrandGORM{Db: database}
	return BrandService{repository: rep}, nil
}

func (service BrandService) CreateBrand(ctx context.Context, req *pb.CreateBrandRequest) (*pb.CreateBrandResponse, error) {
	brand := builders.BuildBrand(req.Payload)
	service.repository.Create(&brand)
	return &pb.CreateBrandResponse{Brand: req.Payload, ExternalId: brand.ExternalId}, nil
}

func (service BrandService) GetBrandById(ctx context.Context, req *pb.GetBrandByIdRequest) (*pb.GetBrandByIdResponse, error) {
	if err := utils.HandleEmptyRequest(&req.ExternalId, "Brand Id"); err != nil {
		return nil, err
	}
	brand := service.repository.GetByExternalId(req.ExternalId)
	return &pb.GetBrandByIdResponse{Brand: builders.BuildBrandDetails(&brand)}, nil
}

func (service BrandService) DeleteBrand(ctx context.Context, req *pb.DeleteBrandRequest) (*pb.DeleteBrandResponse, error) {
	if err := utils.HandleEmptyRequest(&req.ExternalId, "Brand Id"); err != nil {
		return nil, err
	}
	brand := service.repository.GetByExternalId(req.ExternalId)
	if &brand != nil && &brand.ID == nil {
		service.repository.Delete(&brand)
		return &pb.DeleteBrandResponse{}, nil
	}
	return nil, status.Error(codes.Internal, fmt.Sprintf("Brand %v not found.", req.ExternalId))
}

func (service BrandService) UpdateBrand(ctx context.Context, req *pb.UpdateBrandRequest) (*pb.UpdateBrandResponse, error) {
	if err := utils.HandleEmptyRequest(&req.ExternalId, "Brand Id"); err != nil {
		return nil, err
	}
	brand := service.repository.GetByExternalId(req.ExternalId)
	builders.MergeBrand(&brand, req.Payload)
	service.repository.Update(&brand)
	return &pb.UpdateBrandResponse{Brand: builders.BuildBrandDetails(&brand), ExternalId: brand.ExternalId}, nil
}

func (service BrandService) MultiGetBrands(ctx context.Context, req *pb.MultiGetBrandsRequest) (*pb.MultiGetBrandsResponse, error) {
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

func (service BrandService) GetBrandAttributes(ctx context.Context, ){

}
