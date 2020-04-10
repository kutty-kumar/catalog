package domain

type Testimonial struct {
	Base
	Attribute string
	Value string
	Scale string
	Comments string
	DisplayOrder int
	DoctorID uint
}
