package testimonial

import (
	"catalog/pkg/domain"
	"github.com/jinzhu/gorm"
)

type TestimonialGORMRepository struct {
	Db *gorm.DB
}

func (tgr *TestimonialGORMRepository) Create(doctor *domain.Doctor, testimonial *domain.Testimonial) *domain.Testimonial {
	testimonial.DoctorID = doctor.ID
	tgr.Db.Save(testimonial)
	return testimonial
}
