package doctor

import (
	"catalog/pkg/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DoctorGORMRepository struct {
	Db *gorm.DB
}

func (dgr DoctorGORMRepository) Create(doctor *domain.Doctor) *domain.Doctor {
	dgr.Db.Save(doctor)
	return doctor
}

func (dgr DoctorGORMRepository) GetByExternalId(externalId string) (*domain.Doctor, error) {
	doctor := domain.Doctor{Base: domain.Base{ExternalId:externalId}}
	dgr.Db.Where(&doctor).First(&doctor)
	if doctor.ID == 0 {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Doctor %v not found", externalId))
	}
	return &doctor, nil
}

func (dgr DoctorGORMRepository) UpdateDoctor(doctor *domain.Doctor) *domain.Doctor {
	dgr.Db.Model(doctor).Update(doctor)
	return doctor
}

func (dgr DoctorGORMRepository) MultiGet(doctorIds []string) ([]*domain.Doctor, error){
	var doctors []*domain.Doctor
	dgr.Db.Where("external_id IN (?)", doctorIds).Find(&doctors)
	if doctors == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("doctors with %v external ids not found.", doctorIds))
	}
	return doctors, nil
}