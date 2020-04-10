package builders

import (
	"catalog/pkg/domain"
	"catalog/pkg/pb"
)

func BuildTestimonial(testimonialReq *pb.CreateTestimonialRequest) *domain.Testimonial {
	testimonial := testimonialReq.Payload
	return &domain.Testimonial{
		Attribute:            testimonial.Attribute,
		Value:                testimonial.Value,
		Scale:                testimonial.Scale,
		Comments:             testimonial.Comments,
		DisplayOrder:         int(testimonial.DisplayOrder),
	}
}

func BuildTestimonialResponse(testimonial *domain.Testimonial) *pb.CreateTestimonialResponse {
	return &pb.CreateTestimonialResponse{
		Response: &pb.Testimonial{
			Attribute:            testimonial.Attribute,
			Value:                testimonial.Value,
			Scale:                testimonial.Scale,
			Comments:             testimonial.Comments,
			DisplayOrder:         int64(testimonial.DisplayOrder),
		},
	}
}
