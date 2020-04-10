package builders

import (
	"catalog/pkg/domain"
	"catalog/pkg/pb"
)

func GetTestimonialFromDto(payload *pb.Testimonial) *domain.Testimonial {
	return &domain.Testimonial{
		Attribute:    payload.Attribute,
		Value:        payload.Value,
		Scale:        payload.Scale,
		Comments:     payload.Comments,
		DisplayOrder: int(payload.DisplayOrder),
	}
}

func GetDtoFromTestimonial(testimonial *domain.Testimonial) *pb.Testimonial {
	return &pb.Testimonial{
		Attribute:    testimonial.Attribute,
		Value:        testimonial.Value,
		Scale:        testimonial.Scale,
		Comments:     testimonial.Comments,
		DisplayOrder: int64(testimonial.DisplayOrder),
	}
}

func BuildTestimonial(testimonialReq *pb.CreateTestimonialRequest) *domain.Testimonial {
	return GetTestimonialFromDto(testimonialReq.Payload)
}

func BuildTestimonialResponse(testimonial *domain.Testimonial) *pb.CreateTestimonialResponse {
	return &pb.CreateTestimonialResponse{
		Response: GetDtoFromTestimonial(testimonial),
	}
}
