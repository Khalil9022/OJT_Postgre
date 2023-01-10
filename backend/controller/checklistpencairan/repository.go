package checklistpencairan

import (
	"github.com/khalil9022/OJT_Postgre/models"
	"gorm.io/gorm"
)

type Repository interface {
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
