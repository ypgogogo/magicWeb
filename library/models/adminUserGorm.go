package models

import "time"

//AdminUser Table
type AdminUser struct {
	ID       string `gorm:"primary_key;type:varchar(36);not null;"`
	Account  string `gorm:"type:varchar(32);not null;index:account_idx"`
	Password string `gorm:"type:varchar(32);not null;"`
	Secret   string `gorm:"type:varchar(8);not null;"`
	Nick     string `gorm:"type:varchar(16);not null;"`
	Level    int    `gorm:"not null;default 0"`
	Email    string `gorm:"type:varchar(128);not null;"`
	Mobile   string `gorm:"type:varchar(20);not null;"`

	Profile   AdminProfile `gorm:"ForeignKey:ProfileID"`
	ProfileID string       `grom:"type:varchar(36);index"`

	Backstage       uint8  `gorm:"not null;default 0"`
	State           uint8  `gorm:"not null;default 0"`
	Fail            uint16 `gorm:"not null;default 0"`
	FailLastTime    time.Time
	Logined         uint16 `gorm:"not null;default 0"`
	LoginedLastTime time.Time
	LoginedIP       string `gorm:"type:varchar(32);not null;"`
	Source          string `gorm:"type:varchar(64);not null;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreateIP        string `gorm:"type:varchar(32);not null;"`
}

//TableName Returns table name
func (AdminUser) TableName() string {
	return "adminusers"
}
