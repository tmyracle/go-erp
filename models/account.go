package models

import (
	"time"
)

type Account struct {
	ID 				string `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	AcctName string `json:"acctName"`
	AcctNumber string `json:"acctNumber"`
	IsInactive bool `json:"isInactive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeleteAt *time.Time `json:"deleteAt" gorm:"index"`
}
