package domain

type Wilayah struct {
	ID           string `gorm:"column:ID;primaryKey" json:"id"`
	Jenis        int8   `gorm:"column:JENIS" json:"jenis"`
	Deskripsi    string `gorm:"column:DESKRIPSI" json:"deskripsi"`
	Kota         int8   `gorm:"column:KOTA;default:0" json:"kota"`
	Status       int8   `gorm:"column:STATUS;default:1" json:"status"`
	RemoteChange int8   `gorm:"column:REMOTE_CHANGE;default:0" json:"remote_change"`
}

func (Wilayah) TableName() string {
	return "master.wilayah"
}

type JenisWilayah struct {
	ID        int8   `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	Deskripsi string `gorm:"column:DESKRIPSI" json:"deskripsi"`
	Digit     int8   `gorm:"column:DIGIT;default:2" json:"digit"`
	Delimiter string `gorm:"column:DELIMITER" json:"delimiter"`
}

func (JenisWilayah) TableName() string {
	return "master.jenis_wilayah"
}
