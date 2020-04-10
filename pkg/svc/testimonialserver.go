package svc

import (
	"catalog/pkg/builders"
	"catalog/pkg/pb"
	"catalog/pkg/repository/testimonial"
	"context"
	"github.com/jinzhu/gorm"
)

type TestimonialService struct {
	repo testimonial.TestimonialGORMRepository
	doctorService DoctorService
}

func NewTestimonialService(db *gorm.DB, doctorService DoctorService) TestimonialService{
	return TestimonialService{repo: testimonial.TestimonialGORMRepository{Db:db}, doctorService:doctorService}
}

func (ts *TestimonialService) CreateDoctorTestimonial(_ context.Context, req *pb.CreateTestimonialRequest) (*pb.CreateTestimonialResponse, error){
	if doctor, err := ts.doctorService.GetDoctor(req.DoctorId); err != nil {
		return nil, err
	}else {
		testimonialEntity := builders.BuildTestimonial(req)
		testimonialEntity = ts.repo.Create(doctor, testimonialEntity)
		return builders.BuildTestimonialResponse(testimonialEntity), nil
	}
}
