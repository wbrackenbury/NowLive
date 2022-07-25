package data

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/datatypes"
)

const (
	PREVIEW = iota
	WEEKDAY = iota
	WEEKEND = iota
)

type Adjustment struct {
	gorm.Model

	Id string `gorm:"primaryKey"`

	DiscountCode string

	Multiplier float64 `gorm:"default:1.00"`
	Additive float64 `gorm:"default:0.00"`

	TransactId string
	Transact Transact //`gorm:"foreignKey:TransactRefer,references:Id"`
}

type Transact struct {
	gorm.Model

	Id string `gorm:"primaryKey"`

	Quantity uint64
	Rate float64

	Credit uint64
	CreditType uint8

	UserId string
	User User //`gorm:"foreignKey:UserRefer,references:Id"`
	ShowId *string //`gorm:"foreignKey:ShowRefer"`
	Show *Show //Can be nil for external payments
}

type Show struct {
	gorm.Model

	Id string `gorm:"primaryKey"`

	Name string

	PreviewPrice float64
	WeekendPrice float64
	WeekdayPrice float64

	StartDate datatypes.Date
	EndDate datatypes.Date
}



type User struct {
	gorm.Model

	Id string `gorm:"primaryKey"`

	Name string
	Email string
	Phone string
	Address string

	PreviewCredits uint64 `gorm:"default:0"`
	WeekendCredits uint64 `gorm:"default:0"`
	WeekdayCredits uint64 `gorm:"default:0"`

	Notes sql.NullString

}
