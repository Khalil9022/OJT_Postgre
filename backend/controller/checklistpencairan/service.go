package checklistpencairan

import "github.com/khalil9022/OJT_Postgre/models"

type Service interface {
	GetAllCustomerAs9() ([]Respones, error)
	GetSpesifikCustomerAs9(DataRequest) ([]Respones, error)
	GetDataBranch() ([]models.Branch_Tabs, error)
	GetDataCompany() ([]models.Mst_Company_Tabs, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetDataBranch() ([]models.Branch_Tabs, error) {
	Data, err := s.repo.GetDataBranch()
	return Data, err
}

func (s *service) GetDataCompany() ([]models.Mst_Company_Tabs, error) {
	Data, err := s.repo.GetDataCompany()
	return Data, err
}

func (s *service) GetAllCustomerAs9() ([]Respones, error) {
	Data, err := s.repo.GetAllCustomerAs9()
	return Data, err
}

func (s *service) GetSpesifikCustomerAs9(req DataRequest) ([]Respones, error) {
	Data, err := s.repo.GetSpesifikCustomerAs9(req)
	return Data, err
}
