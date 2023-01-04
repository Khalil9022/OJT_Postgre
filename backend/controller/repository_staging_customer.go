package controller

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/khalil9022/OJT_Postgre/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetDataCustomer() ([]models.Staging_Customers, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func generateCustCode(companyCode string, tglSekarang string, custCode string) string {
	AppCustCode := "006"
	tglSekarang = strings.ReplaceAll(tglSekarang, "-", "")

	tambahan := "0000000000" + custCode
	length := len([]rune(tambahan))

	appCustCodeSeq := tambahan[length-10:]

	NewCustCode := AppCustCode + companyCode + tglSekarang + appCustCodeSeq
	return NewCustCode
}

func (r *repository) GetDataCustomer() ([]models.Staging_Customers, error) {
	currentTime := time.Now()
	var staging_customer []models.Staging_Customers
	var customer_data_tabs []models.Customer_Data_Tabs
	var mst_company_tabs []models.Mst_Company_Tabs
	var branch_tabs []models.Branch_Tabs
	var vehicle_data_tabs []models.Vehicle_Data_Tabs

	res := r.db.Where("sc_create_date = ? AND sc_flag = ? ", currentTime.Format("2006-01-02"), "0").Find(&staging_customer)

	if res.Error != nil {
		return []models.Staging_Customers{}, res.Error
	}

	r.db.Find(&customer_data_tabs)
	r.db.Find(&mst_company_tabs)
	r.db.Find(&branch_tabs)
	r.db.Find(&vehicle_data_tabs)
	//validasi
	for _, item := range staging_customer {
		sc_flag := "0"
		log.Println(sc_flag)

		// //validasi 1
		// for _, item2 := range customer_data_tabs {
		// 	if item2.PPK == item.CustomerPpk {
		// 		sc_flag = "8"
		// 		break
		// 	}
		// }

		// //validasi 2
		// for _, item2 := range mst_company_tabs {
		// 	if item2.CompanyShortName == item.ScCompany {
		// 		continue
		// 	} else {
		// 		sc_flag = "8"
		// 		break
		// 	}
		// }

		// // validasi 3
		// for _, item2 := range branch_tabs {
		// 	if item2.Code == item.ScBranchCode {
		// 		continue
		// 	} else {
		// 		sc_flag = "8"
		// 		break
		// 	}
		// }

		tglSekarang := time.Now()
		tglSekarangString := tglSekarang.String()
		// //validasi 4
		// if item.LoanTglPk[0:7] != tglSekarangString[0:7] {
		// 	sc_flag = "8"
		// }

		// //validasi 5
		// if item.CustomerIDType == "1" {
		// 	if item.CustomerIDNumber == "" {
		// 		sc_flag = "8"
		// 	}
		// }

		//validasi 6
		regex := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]+`)
		if matched := regex.MatchString(item.CustomerName); matched {
			sc_flag = "8"
		}

		// //validasi 7
		// if item.VehicleBpkb == "" {
		// 	sc_flag = "8"
		// }

		// //validasi 8
		// if item.VehicleStnk == "" {
		// 	sc_flag = "8"
		// }

		// //validasi 9
		// if item.VehicleEngineNo == "" {
		// 	sc_flag = "8"
		// }

		// //validasi 10
		// for _, item2 := range vehicle_data_tabs {
		// 	if item2.EngineNo == item.VehicleEngineNo {
		// 		sc_flag = "8"
		// 		break
		// 	}
		// }

		// //validasi 11
		// if item.VehicleChasisNo == "" {
		// 	sc_flag = "8"
		// }

		// //validasi 12
		// for _, item2 := range vehicle_data_tabs {
		// 	if item2.ChasisNo == item.VehicleChasisNo {
		// 		sc_flag = "8"
		// 		break
		// 	}
		// }

		custCode := generateCustCode(item.ScCompany, tglSekarangString[0:7], strconv.FormatInt(item.ID, 10))
		if sc_flag == "0" {
			sc_flag = "1"
			r.db.Model(&staging_customer).Where("id=?", item.ID).Update("sc_flag", sc_flag)
			r.db.Create(&models.Skala_Rental_Tabs{
				Custcode: custCode,
			})
			continue
		}
		r.db.Model(&staging_customer).Where("id=?", item.ID).Update("sc_flag", sc_flag)

	}

	return staging_customer, nil
}
