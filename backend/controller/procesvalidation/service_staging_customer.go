package procesvalidation

import "github.com/khalil9022/OJT_Postgre/models"

type Service interface {
	PencairanKredit() ([]models.Staging_Customers, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) PencairanKredit() ([]models.Staging_Customers, error) {
	DataCustomer, err := s.repo.PencairanKredit()
	return DataCustomer, err
}
