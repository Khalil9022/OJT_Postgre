package procesvalidation

import (
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

func (r *repository) generateCustCode(companyCode string, tglSekarang string) string {
	var customer_data_tabs models.Customer_Data_Tabs
	appCustCodeSeq := ""
	res := r.db.Last(&customer_data_tabs)
	AppCustCode := "006"

	tglSekarang = strings.ReplaceAll(tglSekarang, "-", "")

	if res.Error != nil {
		appCustCodeSeq = "0000000001"
	} else {
		custCode := customer_data_tabs.Custcode
		custCode = custCode[len(custCode)-10:]
		idNew, _ := strconv.Atoi(custCode)
		idNew = idNew + 1

		tambahan := "0000000000" + strconv.Itoa(idNew)
		length := len([]rune(tambahan))
		appCustCodeSeq = tambahan[length-10:]
	}

	// tambahan := "0000000000" + custCode
	// length := len([]rune(tambahan))

	// AppCustCodeSeq1 := customer_data_tabs.Custcode
	// appCustCodeSeq := tambahan[length-10:]

	NewCustCode := AppCustCode + companyCode + tglSekarang + appCustCodeSeq
	return NewCustCode
}

func (r *repository) GetDataCustomer() ([]models.Staging_Customers, error) {
	currentTime := time.Now()
	var staging_customer []models.Staging_Customers

	res := r.db.Where("sc_create_date = ? AND sc_flag = ? ", currentTime.Format("2006-01-02"), "0").Find(&staging_customer)

	if res.Error != nil {
		return []models.Staging_Customers{}, res.Error
	}

	//validasi
	for _, item := range staging_customer {
		sc_flag := "0"
		companycode := ""
		tglSekarang := time.Now()
		tglSekarangString := tglSekarang.String()

		sc_flag, errorMessage := r.ValidasiCustomerPPK(item.CustomerPpk) //validasi 1
		if sc_flag != "8" {
			sc_flag, errorMessage, companycode = r.ValidasiSCCompany(item.ScCompany)
		}
		if sc_flag != "8" {
			sc_flag, errorMessage = r.ValidasiScBranchCode(item.ScBranchCode)
		}
		if sc_flag != "8" {
			sc_flag, errorMessage = r.ValidasiWaktu(item.LoanTglPk, tglSekarangString)
		}
		if sc_flag != "8" {
			sc_flag, errorMessage = r.ValidasiCustomerIDType(item.CustomerIDType, item.CustomerIDNumber)
		}
		if sc_flag != "8" {
			sc_flag, errorMessage = r.ValidasiNamaDebitur(item.CustomerName)
		}
		if sc_flag != "8" {
			sc_flag, errorMessage = r.ValidasiVehicleBPKB(item.VehicleBpkb)
		}
		if sc_flag != "8" {
			sc_flag, errorMessage = r.ValidasiVehicleSTNK(item.VehicleStnk)
		}
		if sc_flag != "8" {
			sc_flag, errorMessage = r.ValidasiVehicleEngineNo(item.VehicleEngineNo)
		}
		if sc_flag != "8" {
			sc_flag, errorMessage = r.ValidasiVehicleDuplicateEngine(item.VehicleEngineNo)
		}
		if sc_flag != "8" {
			sc_flag, errorMessage = r.ValidasiVehicleChasisNo(item.VehicleChasisNo)
		}

		drawdowndate, _ := time.Parse("2006-01-02", item.LoanTglPk)
		tglpkchanneling, _ := time.Parse("2006-01-02", item.LoanTglPkChanneling)
		customerbirthdate, _ := time.Parse("2006-01-02", item.CustomerBirthDate)
		customeridtype, _ := strconv.ParseInt(item.CustomerIDType, 10, 8)
		loaninterestflatchanneling, _ := strconv.ParseFloat(item.LoanInterestFlatChanneling, 32)
		loaninteresteffectivechanneling, _ := strconv.ParseFloat(item.LoanInterestEffectiveChanneling, 32)
		loaneffectivepaymenttype, _ := strconv.ParseInt(item.LoanEffectivePaymentType, 10, 8)
		vehiclestatus, _ := strconv.ParseInt(item.VehicleStatus, 10, 8)
		vehicledealerID, _ := strconv.ParseInt(item.VehicleDealerID, 10, 8)
		custCode := r.generateCustCode(companycode, tglSekarangString[0:7])
		vehicletglstnk, _ := time.Parse("2006-01-02", item.VehicleTglStnk)
		vehicletglbpkb, _ := time.Parse("2006-01-02", item.VehicleTglBpkb)
		collateraltypeID, _ := strconv.ParseInt(item.CollateralTypeID, 10, 64)

		if sc_flag == "0" {
			sc_flag = "1"
			r.db.Model(&staging_customer).Where("id=?", item.ID).Update("sc_flag", sc_flag)
			r.db.Create(&models.Customer_Data_Tabs{
				Custcode:          custCode,
				PPK:               item.CustomerPpk,
				Name:              item.CustomerName,
				Address1:          item.CustomerAddress1,
				Address2:          item.CustomerAddress2,
				City:              item.CustomerCity,
				Zip:               item.CustomerZip,
				BirthPlace:        item.CustomerBirthPlace,
				BirthDate:         customerbirthdate,
				IdType:            int8(customeridtype),
				IdNumber:          item.CustomerIDNumber,
				MobileNo:          item.CustomerMobileNo,
				DrawdownDate:      drawdowndate,
				TglPkChanneling:   tglpkchanneling,
				MotherMaidenName:  item.CustomerMotherMaidenName,
				ChannelingCompany: item.ScCompany,
				ApprovalStatus:    "9",
			})
			r.db.Create(&models.Loan_Data_Tabs{
				Custcode:             custCode,
				Branch:               item.ScBranchCode,
				OTR:                  item.LoanOtr,
				DownPayment:          item.LoanDownPayment,
				LoanAmount:           item.LoanLoanAmountChanneling,
				LoanPeriod:           item.LoanLoanPeriodChanneling,
				InterestType:         0,
				InterestFlat:         float32(loaninterestflatchanneling),
				InterestEffective:    float32(loaninteresteffectivechanneling),
				EffectivePaymentType: int8(loaneffectivepaymenttype),
				AdminFee:             "0",
				MonthlyPayment:       item.LoanMonthlyPaymentChanneling,
				InputDate:            tglSekarang,
				LastModified:         tglSekarang,
				ModifiedBy:           "system",
				InputDate2:           tglSekarang,
				LastModified2:        tglSekarang,
				ModifiedBy2:          "system",
				InputBy:              "system",
			})
			r.db.Create(&models.Vehicle_Data_Tabs{
				Custcode:       custCode,
				Brand:          item.VehicleBrand,
				Type:           item.VehicleType,
				Year:           item.VehicleYear,
				Golongan:       1,
				Jenis:          item.VehicleJenis,
				Status:         int8(vehiclestatus),
				Color:          item.VehicleColor,
				PoliceNo:       item.VehiclePoliceNo,
				EngineNo:       item.VehicleEngineNo,
				ChasisNo:       item.VehicleChasisNo,
				Bpkb:           item.VehicleBpkb,
				RegisterNo:     "source",
				Stnk:           item.VehicleStnk,
				StnkAddress1:   "source",
				StnkAddress2:   "source",
				StnkCity:       item.VehicleCityDealer,
				DealerID:       int(vehicledealerID),
				Inputdate:      tglSekarang,
				Inputby:        "system",
				Lastmodified:   tglSekarang,
				Modifiedby:     "system",
				TglStnk:        vehicletglstnk,
				TglBpkb:        vehicletglbpkb,
				TglPolis:       time.Now(), //cari
				PolisNo:        item.VehiclePoliceNo,
				CollateralID:   collateraltypeID,
				Ketagunan:      "source",
				AgunanLbu:      "source",
				Dealer:         item.VehicleDealer,
				AddressDealer1: item.VehicleAddressDealer1,
				AddressDealer2: item.VehicleAddressDealer2,
				CityDealer:     item.VehicleCityDealer,
			})
			continue
		}
		r.db.Model(&staging_customer).Where("id=?", item.ID).Update("sc_flag", sc_flag)
		r.db.Create(&models.Staging_Errors{
			SeReff:       item.ScReff,
			SeCreateDate: item.ScCreateDate,
			BranchCode:   item.ScBranchCode,
			Company:      item.ScCompany,
			Ppk:          item.CustomerPpk,
			Name:         item.CustomerName,
			ErrorDesc:    errorMessage,
		})
	}

	return staging_customer, nil
}
