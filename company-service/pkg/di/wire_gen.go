// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"company-service/pkg/api"
	"company-service/pkg/api/service"
	"company-service/pkg/config"
	"company-service/pkg/service/random"
	"company-service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.Server, error) {
	randomGenerator, err := random.NewRandomGenerator()
	if err != nil {
		return nil, err
	}
	companyUseCase := usecase.NewCompanyUseCase(randomGenerator)
	companyServiceServer := service.NewCompanyServiceServer(companyUseCase)
	server, err := api.NewServerGRPC(cfg, companyServiceServer)
	if err != nil {
		return nil, err
	}
	return server, nil
}
