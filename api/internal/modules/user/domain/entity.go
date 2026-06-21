package domain

type User struct {
	ID                     uint   `gorm:"column:ID;primaryKey"`
	USERNAME               string `gorm:"column:USERNAME"`
	PASSWORD               string `gorm:"column:PASSWORD"`
	NAMA                   string `gorm:"column:NAMA"`
	NIK                    string `gorm:"column:NIK"`
	ID_FASKES              uint   `gorm:"column:ID_FASKES"`
	MASA_AKTIF             string `gorm:"column:MASA_AKTIF"`
	TOKEN                  string `gorm:"column:TOKEN"`
	ROLE                   string `gorm:"column:ROLE"`
	STATUS_TELEGRAM        string `gorm:"column:STATUS_TELEGRAM"`
	ID_TELEGRAM            string `gorm:"column:ID_TELEGRAM"`
	TERAKHIR_UBAH_PASSWORD string `gorm:"column:TERAKHIR_UBAH_PASSWORD"`
	STATUS                 int16  `gorm:"column:STATUS"`
}

// TableName overrides the default table name to match database table 'user'
func (User) TableName() string {
	return "user"
}
