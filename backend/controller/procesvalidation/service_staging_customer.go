package procesvalidation

import "github.com/khalil9022/OJT_Postgre/models"

type Service interface {
	GetDataCustomer() ([]models.Staging_Customers, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetDataCustomer() ([]models.Staging_Customers, error) {
	DataCustomer, err := s.repo.GetDataCustomer()
	return DataCustomer, err
}
