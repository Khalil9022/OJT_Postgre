package checklistpencairan

import (
	"github.com/khalil9022/OJT_Postgre/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllCustomerAs9() ([]Respones, error)
	GetSpesifikCustomerAs9(DataRequest) ([]Respones, error)
	GetDataBranch() ([]models.Branch_Tabs, error)
	GetDataCompany() ([]models.Mst_Company_Tabs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetDataBranch() ([]models.Branch_Tabs, error) {

	var branch_tabs []models.Branch_Tabs

	res := r.db.Find(&branch_tabs)
	if res.Error != nil {
		return nil, res.Error
	}

	return branch_tabs, nil
}

func (r *repository) GetDataCompany() ([]models.Mst_Company_Tabs, error) {

	var mst_company_tabs []models.Mst_Company_Tabs

	res := r.db.Find(&mst_company_tabs)
	if res.Error != nil {
		return nil, res.Error
	}

	return mst_company_tabs, nil
}

func (r *repository) GetAllCustomerAs9() ([]Respones, error) {

	var customer_data_tabs []Respones

	res := r.db.Table("customer_data_tabs").Select("customer_data_tabs.ppk, customer_data_tabs.name,customer_data_tabs.channeling_company, customer_data_tabs.drawdown_date, loan_data_tabs.loan_amount,loan_data_tabs.loan_period,loan_data_tabs.interest_effective").Joins("left join loan_data_tabs on loan_data_tabs.custcode = customer_data_tabs.custcode").Where("approval_status=?", "9").Scan(&customer_data_tabs)
	if res.Error != nil {
		return nil, res.Error
	}

	return customer_data_tabs, nil
}

func (r *repository) GetSpesifikCustomerAs9(req DataRequest) ([]Respones, error) {
	branch := ""
	company := ""
	if req.Branch == "000" {
		branch = "%%"
	} else {
		branch = req.Branch
	}
	if req.Company == "000" {
		company = "%%"
	} else {
		company = req.Company
	}

	var customer_data_tabs []Respones
	res := r.db.Table("customer_data_tabs").Select("customer_data_tabs.ppk, customer_data_tabs.name,customer_data_tabs.channeling_company, customer_data_tabs.drawdown_date, loan_data_tabs.loan_amount,loan_data_tabs.loan_period,loan_data_tabs.interest_effective").Joins("left join loan_data_tabs on loan_data_tabs.custcode = customer_data_tabs.custcode").Where("approval_status=? AND branch LIKE ? AND channeling_company LIKE ? AND drawdown_date BETWEEN ? AND ?", "9", branch, company, req.Start, req.End).Scan(&customer_data_tabs)
	if res.Error != nil {
		return nil, res.Error
	}

	return customer_data_tabs, nil
}
