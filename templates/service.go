package templates

const ServiceTemplate = `package services

import (
	"MODULE_NAME/repositories"
)

// SERVICE_NAMEService interface define
type SERVICE_NAMEService interface {
	
}

type service_nameService struct {
	repository_nameRepo repositories.REPOSITORY_NAMERepository
}

// NewSERVICE_NAMEService create new campaign service handler
func NewSERVICE_NAMEService(
	repository_nameRepo repositories.REPOSITORY_NAMERepository,
) SERVICE_NAMEService {
	return &service_nameService{
		repository_nameRepo: repository_nameRepo,
	}
}

// TODO: Implement DoSomething
func (s *service_nameService) DoSomething() error {
	return nil
}
`
