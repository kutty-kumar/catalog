package svc

import (
	"catalog/pkg/builders"
	"catalog/pkg/domain"
	"catalog/pkg/pb"
	"catalog/pkg/repository/doctor"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DoctorService struct {
	repo doctor.DoctorGORMRepository
}

func NewDoctorService(db *gorm.DB) DoctorService{
	var doctorRepo doctor.DoctorGORMRepository
	doctorRepo = doctor.DoctorGORMRepository{Db:db}
	return DoctorService{repo:doctorRepo}
}

func (ds *DoctorService) CreateDoctor(_ context.Context, req *pb.CreateDoctorRequest) (*pb.CreateDoctorResponse, error){
	doctorEntity := builders.BuildDoctor(req.Payload)
	ds.repo.Create(doctorEntity)
	return builders.BuildCreateDoctorResponse(doctorEntity), nil
}

func (ds *DoctorService) GetDoctor(externalId string) (*domain.Doctor, error){
	return ds.repo.GetByExternalId(externalId)
}

func (ds *DoctorService) UpdateDoctor(_ context.Context, req *pb.UpdateDoctorRequest) (*pb.UpdateDoctorResponse, error){
	doctorEntity := builders.BuildDoctor(req.Payload)
	updatedDoctorEntity := ds.repo.UpdateDoctor(doctorEntity)
	return builders.BuildUpdateDoctorResponse(updatedDoctorEntity), nil
}

func (ds *DoctorService) GetDoctorByExternalId(_ context.Context, req *pb.GetDoctorByIdRequest) (*pb.GetDoctorByIdResponse, error){
	d, err := ds.GetDoctor(req.ExternalId)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("doctor with %v external id not found", req.ExternalId))
	}
	return &pb.GetDoctorByIdResponse{Response: builders.GetDoctorResponse(d)}, nil
}

func (ds *DoctorService) MultiGetDoctors(_ context.Context, req *pb.MultiGetDoctorsRequest) (*pb.MultiGetDoctorsResponse, error){
	doctors, err := ds.repo.MultiGet(req.DoctorIds)
	if err != nil {
		return nil, err
	}
	return builders.BuildMultiGetDoctorsResponse(doctors), nil
}
