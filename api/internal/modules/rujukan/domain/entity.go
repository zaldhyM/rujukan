package domain

import "time"

type Rujukan struct {
	ID             int        `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	NoRujukan      string     `gorm:"column:NO_RUJUKAN;unique" json:"no_rujukan"`
	IDPasien       string     `gorm:"column:ID_PASIEN" json:"id_pasien"`
	IDFaskesAsal   string     `gorm:"column:ID_FASKES_ASAL" json:"id_faskes_asal"`
	IDFaskesTujuan string     `gorm:"column:ID_FASKES_TUJUAN" json:"id_faskes_tujuan"`
	KodeICD        *string    `gorm:"column:KODE_ICD" json:"kode_icd"`
	Diagnosa       *string    `gorm:"column:DIAGNOSA" json:"diagnosa"`
	Catatan        *string    `gorm:"column:CATATAN" json:"catatan"`
	Status         string     `gorm:"column:STATUS;default:menunggu" json:"status"`
	IDUser         uint       `gorm:"column:ID_USER" json:"id_user"`
	Tanggal        *time.Time `gorm:"column:TANGGAL;autoCreateTime" json:"tanggal"`
}

func (Rujukan) TableName() string {
	return "rujukan.rujukan"
}
