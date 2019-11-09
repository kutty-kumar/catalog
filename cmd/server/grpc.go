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
	s, err := svc.NewBrandService(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterCatalogServer(grpcServer, s)

	return grpcServer, nil
}

func createTables(db *gorm.DB){
	db.CreateTable(domain.Brand{})
}

func dropTables(db *gorm.DB){
	db.DropTableIfExists(domain.Brand{})
}
