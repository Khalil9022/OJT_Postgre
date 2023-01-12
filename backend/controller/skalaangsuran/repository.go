package skalaangsuran

import (
	"strconv"
	"strings"
	"time"

	"github.com/khalil9022/OJT_Postgre/models"
	"gorm.io/gorm"
)

type Repository interface {
	GenerateSkalaAngsuran() ([]models.Customer_Data_Tabs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GenerateSkalaAngsuran() ([]models.Customer_Data_Tabs, error) {
	var loan_data_tabs models.Loan_Data_Tabs
	var customer_data_tabs []models.Customer_Data_Tabs

	res := r.db.Where("approval_status = ?", "0").Find(&customer_data_tabs)
	if res.Error != nil {
		return nil, res.Error
	}

	r.db.Exec("UPDATE customer_data_tabs SET approval_status = 1 WHERE approval_status = 0")
	for _, item := range customer_data_tabs {
		res = r.db.Where("custcode = ? ", item.Custcode).First(&loan_data_tabs)
		if res.Error != nil {
			return nil, res.Error
		}

		Counter, _ := strconv.Atoi(loan_data_tabs.LoanPeriod)
		loanamount := loan_data_tabs.LoanAmount
		loanamount = strings.ReplaceAll(loanamount, "Rp", "")
		loanamount = strings.ReplaceAll(loanamount, ".", "")
		loanamount = strings.ReplaceAll(loanamount, "$", "")
		loanamount = strings.ReplaceAll(loanamount, ",00", "")

		loanamountInt, _ := strconv.Atoi(loanamount)
		monthlypayment := loan_data_tabs.MonthlyPayment
		monthlypayment = strings.ReplaceAll(monthlypayment, "Rp", "")
		monthlypayment = strings.ReplaceAll(monthlypayment, ".", "")
		monthlypayment = strings.ReplaceAll(monthlypayment, "$", "")
		monthlypayment = strings.ReplaceAll(monthlypayment, ",00", "")
		monthlypaymentInt, _ := strconv.Atoi(monthlypayment)

		endbalance, principle, osbalance := 0, 0, 0
		duedate := loan_data_tabs.InputDate
		interest := 0

		for i := 0; i <= Counter; i++ {
			if i == 0 {
				osbalance = loanamountInt
				endbalance = osbalance
			} else {
				osbalance = endbalance
				interest = (int(osbalance) * int(loan_data_tabs.InterestEffective) * 30) / 36000
				principle = monthlypaymentInt - int(interest)
				endbalance = osbalance - principle
				duedate = duedate.AddDate(0, 0, 30)
				if endbalance < 0 {
					endbalance = 0
					osbalance = osbalance + (-endbalance)
				}
			}

			r.db.Create(&models.Skala_Rental_Tabs{
				Custcode:      item.Custcode,
				Counter:       int8(i),
				Osbalance:     strconv.Itoa(osbalance),
				EndBalance:    strconv.Itoa(endbalance),
				DueDate:       duedate,
				EffRate:       float64(loan_data_tabs.InterestEffective),
				Rental:        strconv.Itoa(monthlypaymentInt),
				Interest:      strconv.Itoa(interest),
				Principle:     strconv.Itoa(principle),
				Inputdate:     loan_data_tabs.InputDate,
				Inputby:       "system",
				Lastmodified:  time.Now(),
				Modifiedby:    "system",
				PaymentDate:   time.Now(),
				Penalty:       "0",
				PaymentAmount: "0",
				PaymentType:   1,
			})
		}

	}

	return customer_data_tabs, nil
}
