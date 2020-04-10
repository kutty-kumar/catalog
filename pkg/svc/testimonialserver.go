package svc

import (
	"catalog/pkg/builders"
	"catalog/pkg/domain"
	"catalog/pkg/pb"
	"catalog/pkg/repository/testimonial"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestimonialService struct {
	repo          testimonial.TestimonialGORMRepository
	doctorService DoctorService
}

type testimonialFunc func(doctor *domain.Doctor, req *pb.Testimonial) (*pb.Testimonial, error)

func NewTestimonialService(db *gorm.DB, doctorService DoctorService) TestimonialService {
	return TestimonialService{repo: testimonial.TestimonialGORMRepository{Db: db}, doctorService: doctorService}
}

func (ts *TestimonialService) doctorPresent(doctorId string) (*domain.Doctor, error) {
	return ts.doctorService.GetDoctor(doctorId)
}

func (ts *TestimonialService) testimonialOperation(doctorId string, req *pb.Testimonial, f testimonialFunc) (*pb.Testimonial, error) {
	doctor, err := ts.doctorPresent(doctorId)
	if err != nil {
		return nil, err
	}
	return f(doctor, req)
}

func (ts *TestimonialService) CreateDoctorTestimonial(_ context.Context, req *pb.CreateTestimonialRequest) (*pb.CreateTestimonialResponse, error) {
	function := func(doctor *domain.Doctor, req *pb.Testimonial) (*pb.Testimonial, error) {
		testimonialEntity := builders.GetTestimonialFromDto(req)
		ts.repo.Create(doctor, testimonialEntity)
		return builders.GetDtoFromTestimonial(testimonialEntity), nil

	}
	tl, err := ts.testimonialOperation(req.DoctorId, req.Payload, function)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTestimonialResponse{Response: tl}, nil
}

func (ts *TestimonialService) UpdateDoctorTestimonial(_ context.Context, req *pb.UpdateDoctorTestimonialRequest) (*pb.UpdateDoctorTestimonialResponse, error) {
	function := func(doctor *domain.Doctor, req *pb.Testimonial) (*pb.Testimonial, error) {
		testimonialEntity := builders.GetTestimonialFromDto(req)
		ts.repo.Update(testimonialEntity)
		return builders.GetDtoFromTestimonial(testimonialEntity), nil

	}
	tl, err := ts.testimonialOperation(req.DoctorId, req.Payload, function)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateDoctorTestimonialResponse{Response: tl}, nil
}

func (ts *TestimonialService) GetDoctorTestimonialByExternalId(_ context.Context, req *pb.GetDoctorTestimonialByIdRequest) (*pb.GetDoctorTestimonialByIdResponse, error) {
	if _, err := ts.doctorPresent(req.DoctorId); err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Doctor with %v external_id not found.", req.DoctorId))
	}
	if tl, err := ts.repo.GetByExternalId(req.TestimonialId); err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Testimonial with %v external_id not found.", req.TestimonialId))
	} else {
		return &pb.GetDoctorTestimonialByIdResponse{Response: builders.GetDtoFromTestimonial(tl)}, nil
	}
}

func (ts *TestimonialService) MultiGetDoctorTestimonials(_ context.Context, req *pb.MultiGetDoctorTestimonialsRequest) (*pb.MultiGetDoctorTestimonialsResponse, error) {
	var tlsResponse []*pb.Testimonial
	testimonials, err := ts.repo.MultiGetByExternalId(req.DoctorId)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("No testimonials found for doctor %v.", req.DoctorId))
	}
	for _, tl := range testimonials {
		tlsResponse = append(tlsResponse, builders.GetDtoFromTestimonial(tl))
	}
	return &pb.MultiGetDoctorTestimonialsResponse{Testimonials: tlsResponse}, nil
}

func (ts *TestimonialService) GetTestimonialsForDoctors(_ context.Context, req *pb.GetTestimonialsForDoctorsRequest) (*pb.GetTestimonialsForDoctorsResponse, error) {
	var responseMap = make(map[uint64][]*pb.Testimonial)
	var responseList []*pb.DoctorTestimonial
	testimonials, err := ts.repo.MultiGetTestimonialsForDoctors(req.DoctorIds)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("No testimonials found for doctors %v.", req.DoctorIds))
	}
	for _, tl := range testimonials {
		responseMap[uint64(tl.DoctorID)] = append(responseMap[uint64(tl.DoctorID)], builders.GetDtoFromTestimonial(tl))
	}
	for doctorId, testimonialList := range responseMap {
		doctorTestimonial := pb.DoctorTestimonial{DoctorId: doctorId, Testimonials: testimonialList}
		responseList = append(responseList, &doctorTestimonial)
	}
	return &pb.GetTestimonialsForDoctorsResponse{DoctorTestimonials: responseList}, nil
}
