package main

import (
	"catalog/pkg/domain"
	"catalog/pkg/pb"
	"catalog/pkg/svc"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewGRPCServer(logger *logrus.Logger, dbConnectionString string) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    time.Duration(viper.GetInt("config.keepalive.time")) * time.Second,
				Timeout: time.Duration(viper.GetInt("config.keepalive.timeout")) * time.Second,
			},
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// logging middleware
				grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),

				// Request-Id interceptor
				requestid.UnaryServerInterceptor(),

				// validation middleware
				grpc_validator.UnaryServerInterceptor(),

				// collection operators middleware
				gateway.UnaryServerInterceptor(),
			),
		),
	)

	// create new postgres database
	db, err := gorm.Open("mysql", dbConnectionString)
	// TODO this is just for dev purposes
	dropTables(db)
	createTables(db)
	if err != nil {
		return nil, err
	}
	// register service implementation with the grpcServer
	brandService, err := svc.NewBrandService(db)
	brandAttributeService, err := svc.NewBrandAttributeService(db, brandService)
	doctorService := svc.NewDoctorService(db)
	doctorTestimonialService := svc.NewTestimonialService(db, doctorService)
	if err != nil {
		return nil, err
	}
	pb.RegisterBrandServiceServer(grpcServer, brandService)
	pb.RegisterBrandAttributeServiceServer(grpcServer, brandAttributeService)
	pb.RegisterDoctorServiceServer(grpcServer, &doctorService)
	pb.RegisterDoctorTestimonialServiceServer(grpcServer, &doctorTestimonialService)
	return grpcServer, nil
}

func createTables(db *gorm.DB){
	db.CreateTable(domain.Brand{}, domain.BrandAttribute{}, domain.Doctor{}, domain.Testimonial{})
}

func dropTables(db *gorm.DB){
	db.DropTableIfExists(domain.Brand{}, domain.BrandAttribute{}, domain.Doctor{}, domain.Testimonial{})
}
