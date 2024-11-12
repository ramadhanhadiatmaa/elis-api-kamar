package models

type Kamar struct {
	Id      string `gorm:"type:varchar(25); primaryKey" json:"id"`
	Bangsal string `gorm:"type:varchar(150)" json:"bangsal"`
	Jenis   string `gorm:"type:varchar(50)" json:"jenis"`
	Total   string `gorm:"type:varchar(250)" json:"total"`
	Isi     string `gorm:"type:varchar(50)" json:"isi"`
}
