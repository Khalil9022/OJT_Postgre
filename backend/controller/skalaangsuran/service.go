package skalaangsuran

import "github.com/khalil9022/OJT_Postgre/models"

type Service interface {
	GenerateSkalaAngsuran() ([]models.Customer_Data_Tabs, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GenerateSkalaAngsuran() ([]models.Customer_Data_Tabs, error) {
	Data, err := s.repo.GenerateSkalaAngsuran()
	return Data, err
}
