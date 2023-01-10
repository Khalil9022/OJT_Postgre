package checklistpencairan

import "github.com/khalil9022/OJT_Postgre/models"

type Service interface {
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
