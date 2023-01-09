package procesvalidation

import (
	"regexp"

	"github.com/khalil9022/OJT_Postgre/models"
)

func (r *repository) ValidasiCustomerPPK(CustomerPPK string) (string, string) {
	var customer_data_tabs []models.Customer_Data_Tabs
	sc_flag := "0"
	errorMessage := ""

	r.db.Find(&customer_data_tabs)
	//validasi 1
	for _, item2 := range customer_data_tabs {
		if item2.PPK == CustomerPPK {
			sc_flag = "8"
			errorMessage = "Customer PPK Tidak memenuhi syarat"
			break
		}
	}
	return sc_flag, errorMessage
}

func (r *repository) ValidasiSCCompany(ScCompany string) (string, string, string) {
	var mst_company_tabs []models.Mst_Company_Tabs
	sc_flag := "0"
	errorMessage := ""
	validasicompanyname := 0
	companycode := ""

	r.db.Find(&mst_company_tabs)

	for _, item2 := range mst_company_tabs {
		if item2.CompanyShortName == ScCompany {
			companycode = item2.CompanyCode
			validasicompanyname = 1
			break
		}
	}

	if validasicompanyname == 0 {
		sc_flag = "8"
		errorMessage = "Company tidak terdaftar"
	}
	return sc_flag, errorMessage, companycode
}

func (r *repository) ValidasiScBranchCode(ScBranchCode string) (string, string) {
	var branch_tabs []models.Branch_Tabs
	sc_flag := "0"
	errorMessage := ""
	validasibranchcode := 0

	r.db.Find(&branch_tabs)

	for _, item2 := range branch_tabs {
		if item2.Code == ScBranchCode {
			validasibranchcode = 1
			break
		}
	}

	if validasibranchcode == 0 {
		sc_flag = "8"
		errorMessage = "Branch Code tidak terdaftar"
	}
	return sc_flag, errorMessage
}

func (r *repository) ValidasiWaktu(LoanTglPk string, tglSekarangString string) (string, string) {
	sc_flag := "0"
	errorMessage := ""

	if LoanTglPk[0:7] != tglSekarangString[0:7] {
		sc_flag = "8"
		errorMessage = "Bulan Dan Tahun Tidak boleh berbeda"
	}
	return sc_flag, errorMessage
}

func (r *repository) ValidasiCustomerIDType(CustomerIDType string, CustomerIDNumber string) (string, string) {
	sc_flag := "0"
	errorMessage := ""
	if CustomerIDType == "1" {
		if CustomerIDNumber == "" {
			sc_flag = "8"
			errorMessage = "Customer ID Number Kosong"
		}
	}
	return sc_flag, errorMessage
}

func (r *repository) ValidasiNamaDebitur(CustomerName string) (string, string) {
	sc_flag := "0"
	errorMessage := ""
	regex := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]+`)
	if matched := regex.MatchString(CustomerName); matched {
		sc_flag = "8"
		errorMessage = "Nama Debitur Tidak boleh mengandung special karakter"
	}
	return sc_flag, errorMessage
}

func (r *repository) ValidasiVehicleBPKB(VehicleBpkb string) (string, string) {
	var vehicle_data_tabs []models.Vehicle_Data_Tabs
	sc_flag := "0"
	errorMessage := ""
	r.db.Find(&vehicle_data_tabs)

	if VehicleBpkb == "" {
		sc_flag = "8"
		errorMessage = "BPKB tidak boleh kosong"
	}

	return sc_flag, errorMessage
}

func (r *repository) ValidasiVehicleSTNK(VehicleStnk string) (string, string) {
	var vehicle_data_tabs []models.Vehicle_Data_Tabs
	sc_flag := "0"
	errorMessage := ""
	r.db.Find(&vehicle_data_tabs)

	if VehicleStnk == "" {
		sc_flag = "8"
		errorMessage = "STNK tidak boleh kosong"
	}
	return sc_flag, errorMessage
}

func (r *repository) ValidasiVehicleEngineNo(VehicleEngineNo string) (string, string) {
	sc_flag := "0"
	errorMessage := ""

	if VehicleEngineNo == "" {
		sc_flag = "8"
		errorMessage = "Engine Nomor Tidak boleh kosong"
	}
	return sc_flag, errorMessage
}

func (r *repository) ValidasiVehicleDuplicateEngine(VehicleEngineNo string) (string, string) {
	var vehicle_data_tabs []models.Vehicle_Data_Tabs
	sc_flag := "0"
	errorMessage := ""
	r.db.Find(&vehicle_data_tabs)

	for _, item2 := range vehicle_data_tabs {
		if item2.EngineNo == VehicleEngineNo {
			sc_flag = "8"
			errorMessage = "Engine Nomor Tidak boleh duplikat"
			break
		}
	}
	return sc_flag, errorMessage
}

func (r *repository) ValidasiVehicleChasisNo(VehicleChasisNo string) (string, string) {
	sc_flag := "0"
	errorMessage := ""

	if VehicleChasisNo == "" {
		sc_flag = "8"
		errorMessage = "Chasis No tidak boleh kosong"
	}
	return sc_flag, errorMessage
}

func (r *repository) ValidasiVehicleDuplicateChasis(VehicleChasisNo string) (string, string) {
	var vehicle_data_tabs []models.Vehicle_Data_Tabs
	sc_flag := "0"
	errorMessage := ""
	r.db.Find(&vehicle_data_tabs)

	for _, item2 := range vehicle_data_tabs {
		if item2.ChasisNo == VehicleChasisNo {
			sc_flag = "8"
			errorMessage = "Chasis No tidak boleh duplikat"
			break
		}
	}
	return sc_flag, errorMessage
}
