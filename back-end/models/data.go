package models

import "time"

type Data struct {
	ID              int       `json:"id" gorm:"primary_key:auto_increment;"`
	ProductID       int64     `json:"productID" form:"productID" gorm:"type: bigint"`
	ProductName     string    `json:"productName" form:"productName" gorm:"type: varchar(255)"`
	Amount          int       `json:"amount" form:"amount" gorm:"type: int"`
	CustomerName    string    `json:"customerName" form:"customerName" gorm:"type: varchar(255)"`
	StatusID        int       `json:"statusID"`
	Status          Status    `json:"status" gorm:"foreignKey:StatusID"`
	TransactionDate time.Time `json:"transactionDate"`
	CreateBy        string    `json:"createBy" form:"createBy" gorm:"type: varchar(255)"`
	CreateOn        time.Time `json:"createOn"`
}
