package models

import "time"

type Skala_Rental_Tabs struct {
	Custcode      string    `json:"custcode" gorm:"not null; type: varchar(25)"`
	Counter       int8      `json:"counter" gorm:"type:smallint"`
	Osbalance     string    `json:"osbalance"  gorm:"type:money"`
	EndBalance    string    `json:"end_balance"  gorm:"type:money"`
	DueDate       time.Time `json:"due_date"  gorm:"type:timestamp"`
	EffRate       float64   `json:"eff_rate"  gorm:"type:float"`
	Rental        string    `json:"rental"  gorm:"type:money"`
	Principle     string    `json:"principle"  gorm:"type:money"`
	Interest      string    `json:"interest"  gorm:"type:money"`
	Inputdate     time.Time `json:"inputdate"  gorm:"type:timestamp"`
	Inputby       string    `json:"inputby"  gorm:"type:varchar(50)"`
	Lastmodified  time.Time `json:"lastmodified"  gorm:"type:timestamp"`
	Modifiedby    string    `json:"modifiedby"  gorm:"type:varchar(50)"`
	PaymentDate   time.Time `json:"payment_date"  gorm:"type:timestamp"`
	Penalty       string    `json:"penalty"  gorm:"type:money"`
	PaymentAmount string    `json:"payment_amount"  gorm:"type:money"`
	PaymentType   int8      `json:"payment_type" gorm:"type:smallint"`
}
