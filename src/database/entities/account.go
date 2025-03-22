package entities

import (
	timeUtils "go-gin-test-job/src/utils/time"

	"github.com/shopspring/decimal"
)

const AccountTable = "account"

type AccountStatus string

const (
	AccountStatusOn  AccountStatus = "On"
	AccountStatusOff AccountStatus = "Off"
)

var AccountStatusList = []string{string(AccountStatusOn), string(AccountStatusOff)}

type Account struct {
	Id        int64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Address   string          `json:"address" gorm:"uniqueIndex:account_address_unique_idx;type:varchar(64);not null"`
	Balance   decimal.Decimal `json:"balance" gorm:"type:decimal(64,8);default:0;not null"`
	Status    AccountStatus   `json:"status" gorm:"index:account_status_idx;type:enum('On','Off');not null"`
	CreatedAt int64           `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt int64           `json:"updated_at" gorm:"autoUpdateTime;index:account_updated_at_idx;not null"`
	Name      string          `json:"name" gorm:"name;not null"`
	Ranking   uint8           `json:"ranking" gorm:"ranking;not null"`
	Memo      string          `json:"memo"`
}

// Set the table name for the model
func (Account) TableName() string {
	return AccountTable
}

func CreateAccount(address string, status AccountStatus) *Account {
	return &Account{
		Address: address,
		Status:  status,
	}
}

func (a *Account) UpdateBalance(balance decimal.Decimal) map[string]interface{} {
	a.Balance = balance
	a.UpdatedAt = timeUtils.GetUnixTime()
	return map[string]interface{}{
		"Balance":   a.Balance,
		"UpdatedAt": a.UpdatedAt,
	}
}

func (a *Account) UpdateStatus(status AccountStatus) map[string]interface{} {
	a.Status = status
	a.UpdatedAt = timeUtils.GetUnixTime()
	return map[string]interface{}{
		"Status":    a.Status,
		"UpdatedAt": a.UpdatedAt,
	}
}
