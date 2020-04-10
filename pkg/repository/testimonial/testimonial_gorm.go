package testimonial

import (
	"catalog/pkg/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestimonialGORMRepository struct {
	Db *gorm.DB
}

func (tgr *TestimonialGORMRepository) GetByExternalId(testimonialId string) (*domain.Testimonial, error) {
	testimonial := domain.Testimonial{Base: domain.Base{ExternalId: testimonialId}}
	tgr.Db.Model(&testimonial).First(&testimonial)
	if testimonial.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Testimonial with %v external_id not found.", testimonialId))
	}
	return &testimonial, nil
}

func (tgr *TestimonialGORMRepository) MultiGetByExternalId(doctorId string) ([]*domain.Testimonial, error) {
	var testimonials []*domain.Testimonial
	doctor := domain.Doctor{Base: domain.Base{ExternalId: doctorId}}
	tgr.Db.Model(&doctor).Related(&testimonials)
	if testimonials == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Testimonials for doctor %v not found.", doctorId))
	}
	return testimonials, nil
}

func (tgr *TestimonialGORMRepository) MultiGetTestimonialsForDoctors(doctorIds []string) ([]*domain.Testimonial, error) {
	var testimonials []*domain.Testimonial
	var doctors []*domain.Doctor
	tgr.Db.Where("external_id IN (?)", doctorIds).Find(&doctors).Related(&testimonials)
	if testimonials == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Testimonials for doctors %v not found", doctorIds))
	}
	return testimonials, nil
}

func (tgr *TestimonialGORMRepository) Create(doctor *domain.Doctor, testimonial *domain.Testimonial) *domain.Testimonial {
	testimonial.DoctorID = doctor.ID
	tgr.Db.Save(testimonial)
	return testimonial
}

func (tgr *TestimonialGORMRepository) Update(testimonial *domain.Testimonial) *domain.Testimonial {
	tgr.Db.Model(testimonial).Update(testimonial)
	return testimonial
}
