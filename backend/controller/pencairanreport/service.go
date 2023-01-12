package pencairanreport

import "github.com/khalil9022/OJT_Postgre/models"

type Service interface {
	GetAllCustomerAs9() ([]Respones, error)
	GetSpesifikCustomerAs9(DataRequest) ([]Respones, error)
	GetAllCustomerAs01() ([]Respones, error)
	GetSpesifikCustomerAs01(DataRequest) ([]Respones, error)
	GetDataBranch() ([]models.Branch_Tabs, error)
	GetDataCompany() ([]models.Mst_Company_Tabs, error)
	UpdateApprovalStatus(ReqPpk) error
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

func (s *service) GetAllCustomerAs01() ([]Respones, error) {
	Data, err := s.repo.GetAllCustomerAs01()
	return Data, err
}

func (s *service) GetSpesifikCustomerAs01(req DataRequest) ([]Respones, error) {
	Data, err := s.repo.GetSpesifikCustomerAs01(req)
	return Data, err
}

func (s *service) UpdateApprovalStatus(ppk ReqPpk) error {
	err := s.repo.UpdateApprovalStatus(ppk)
	return err
}
