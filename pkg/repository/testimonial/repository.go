package testimonial

import (
	"catalog/pkg/domain"
)

type Repository interface {
	Create(doctor *domain.Doctor, testimonial *domain.Testimonial) *domain.Testimonial
}
