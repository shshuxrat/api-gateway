package services

import (
	"api-gateway/config"
	"api-gateway/genproto/position_service"
	"api-gateway/genproto/profession_service"
	"fmt"

	"google.golang.org/grpc"
)

type ServicesI interface {
	ProfessionService() position_service.ProfessionServiseClient
	AttributeService() position_service.AttributeServiceClient
	CompanyService() position_service.CompanyServiceClient
	PositionService() profession_service.PositionServiceClient
	PositionAttributeService() profession_service.PositionAttributeServiceClient
}

type servicesRepo struct {
	professionService        position_service.ProfessionServiseClient
	attributeService         position_service.AttributeServiceClient
	companyService           position_service.CompanyServiceClient
	positionService          profession_service.PositionServiceClient
	positionAttributeService profession_service.PositionAttributeServiceClient
}

func NewServicesRepo(c *config.Config) (ServicesI, error) {
	connPositionService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", c.PositionServiceHost, c.PositionServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connProfessionService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", c.ProfessionServiceHost, c.ProfessionServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return &servicesRepo{
		professionService:        position_service.NewProfessionServiseClient(connPositionService),
		attributeService:         position_service.NewAttributeServiceClient(connPositionService),
		companyService:           position_service.NewCompanyServiceClient(connPositionService),
		positionService:          profession_service.NewPositionServiceClient(connProfessionService),
		positionAttributeService: profession_service.NewPositionAttributeServiceClient(connProfessionService),
	}, nil
}

func (s *servicesRepo) ProfessionService() position_service.ProfessionServiseClient {
	return s.professionService
}
func (s *servicesRepo) AttributeService() position_service.AttributeServiceClient {
	return s.attributeService
}
func (s *servicesRepo) CompanyService() position_service.CompanyServiceClient {
	return s.companyService
}

func (s *servicesRepo) PositionService() profession_service.PositionServiceClient {
	return s.positionService
}

func (s *servicesRepo) PositionAttributeService() profession_service.PositionAttributeServiceClient {
	return s.positionAttributeService
}
