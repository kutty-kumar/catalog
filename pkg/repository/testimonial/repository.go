package testimonial

import (
	"catalog/pkg/domain"
)

type Repository interface {
	Create(doctor *domain.Doctor, testimonial *domain.Testimonial) *domain.Testimonial
	Update(testimonial *domain.Testimonial) *domain.Testimonial
	GetByExternalId(testimonialId string) (*domain.Testimonial, error)
	MultiGetByExternalId(doctorId string) ([]*domain.Testimonial, error)
	MultiGetTestimonialsForDoctors(doctorIds []string) ([]*domain.Testimonial, error)
}
