package domain

import "time"

type Pasien struct {
	ID           string     `gorm:"column:ID;primaryKey" json:"id"`
	IDFaskes     *int16     `gorm:"column:ID_FAKSES" json:"id_faskes"`
	Norm         string     `gorm:"column:NORM" json:"norm"`
	Nama         string     `gorm:"column:NAMA" json:"nama"`
	Kontak       string     `gorm:"column:KONTAK" json:"kontak"`
	Alamat       string     `gorm:"column:ALAMAT" json:"alamat"`
	TempatLahir  string     `gorm:"column:TEMPAT_LAHIR" json:"tempat_lahir"`
	TanggalLahir *time.Time `gorm:"column:TANGGAL_LAHIR" json:"tanggal_lahir"`
	JenisKelamin *int16     `gorm:"column:JENIS_KELAMIN" json:"jenis_kelamin"`
	NoJKN        string     `gorm:"column:NO_JKN" json:"no_jkn"`
	Nik          string     `gorm:"column:NIK" json:"nik"`
	Status       *int8      `gorm:"column:STATUS;default:1" json:"status"`
	Tanggal      *time.Time `gorm:"column:TANGGAL;autoUpdateTime" json:"tanggal"`
}

func (Pasien) TableName() string {
	return "master.pasien"
}
