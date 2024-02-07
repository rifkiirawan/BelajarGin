package models

type Car struct {
	CarID string `gorm:"primaryKey"`
	Brand string `gorm:"not null;type:varchar(191)"`
	Model string `gorm:"not null;type:varchar(191)"`
	Price int    `gorm:"not null"`
}
