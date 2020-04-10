package builders

import (
	"catalog/pkg/constants"
	"catalog/pkg/domain"
	"catalog/pkg/pb"
	"github.com/golang/protobuf/ptypes"
)

func GetDoctorResponse(doctor *domain.Doctor) *pb.Doctor {
	dateOfBirth, _ := ptypes.TimestampProto(doctor.DateOfBirth)
	return &pb.Doctor{
		FirstName:       doctor.FirstName,
		LastName:        doctor.LastName,
		DateOfBirth:     dateOfBirth,
		Status:          int64(doctor.Status),
		PreSalutation:   doctor.PreSalutation,
		PostSalutation:  doctor.PostSalutation,
		ImageUrl:        doctor.ImageUrl,
		ThumbnailUrl:    doctor.ThumbnailUrl,
		MetaKeywords:    doctor.MetaKeywords,
		MetaDescription: doctor.MetaDescription,
		Display:         doctor.Display,
		DisplayOrder:    int64(doctor.DisplayOrder),
		ExternalId: doctor.ExternalId,
	}
}

func BuildUpdateDoctorResponse(doctor *domain.Doctor) *pb.UpdateDoctorResponse {
	return &pb.UpdateDoctorResponse{
		Response: GetDoctorResponse(doctor),
	}
}

func BuildCreateDoctorResponse(doctor *domain.Doctor) *pb.CreateDoctorResponse {
	return &pb.CreateDoctorResponse{
		Response: GetDoctorResponse(doctor),
	}
}

func BuildMultiGetDoctorsResponse(doctors []*domain.Doctor) *pb.MultiGetDoctorsResponse {
	var doctorsResponse []*pb.Doctor
	for _, doctor := range doctors {
		doctorsResponse = append(doctorsResponse, GetDoctorResponse(doctor))
	}
	return &pb.MultiGetDoctorsResponse{Doctors:doctorsResponse}
}

func BuildDoctor(payload *pb.Doctor) *domain.Doctor {
	dateOfBirth, _ := ptypes.Timestamp(payload.DateOfBirth)
	return &domain.Doctor{
		FirstName:         payload.FirstName,
		LastName:          payload.LastName,
		DateOfBirth:       dateOfBirth,
		Status:            constants.Status(payload.Status),
		MetaDescription:   payload.MetaDescription,
		MetaKeywords:      payload.MetaKeywords,
		DisplayOrder:      int(payload.DisplayOrder),
		ThumbnailUrl:      payload.ThumbnailUrl,
		ImageUrl:          payload.ImageUrl,
		PreSalutation:     payload.PreSalutation,
		PostSalutation:    payload.PostSalutation,
		CapabilityBitMask: int(payload.CapabilityBitMask),
		Display:           payload.Display,
	}
}
