package domain

import (
	"catalog/pkg/constants"
	"time"
)

type Doctor struct {
	Base
	FirstName string
	LastName string
	DateOfBirth time.Time
	Status constants.Status
	//EmailAddresses commons.Set
	//PhoneNumbers commons.Set
	ThumbnailUrl string
	ImageUrl string
	PreSalutation string
	PostSalutation string
	MetaKeywords string
	MetaDescription string
	DisplayOrder int
	CapabilityBitMask int
	Display bool
	//Capabilities commons.Set
	//Gender constants.Gender
	Testimonials []Testimonial `gorm:"many2many:doctor_testimonials"`
}
