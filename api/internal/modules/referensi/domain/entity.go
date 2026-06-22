package domain

type Referensi struct {
	Jenis     int8   `gorm:"column:JENIS;primaryKey" json:"jenis"`
	ID        int16  `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	Deskripsi string `gorm:"column:DESKRIPSI" json:"deskripsi"`
	Status    int8   `gorm:"column:STATUS;default:1" json:"status"`
}

func (Referensi) TableName() string {
	return "master.referensi"
}

type JenisReferensi struct {
	ID        int8   `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	Deskripsi string `gorm:"column:DESKRIPSI" json:"deskripsi"`
	Singkatan string `gorm:"column:SINGKATAN" json:"singkatan"`
	Aplikasi  int8   `gorm:"column:APLIKASI;default:0" json:"aplikasi"`
}

func (JenisReferensi) TableName() string {
	return "master.jenis_referensi"
}
