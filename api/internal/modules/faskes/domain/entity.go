package domain

import "time"

type Faskes struct {
	ID                  int16     `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	Nama                string    `gorm:"column:NAMA" json:"nama"`
	Kode                *string   `gorm:"column:KODE;unique" json:"kode"`
	Alamat              string    `gorm:"column:ALAMAT" json:"alamat"`
	Kontak              string    `gorm:"column:KONTAK" json:"kontak"`
	Email               string    `gorm:"column:EMAIL" json:"email"`
	Prov                string    `gorm:"column:PROV" json:"prov"`
	Kab                 string    `gorm:"column:KAB" json:"kab"`
	KelompokFaskes      int8      `gorm:"column:KELOMPOK FASKES" json:"kelompok_faskes"`
	Lat                 *string   `gorm:"column:LAT" json:"lat"`
	Lng                 *string   `gorm:"column:LNG" json:"lng"`
	Tanggal             time.Time `gorm:"column:TANGGAL;autoUpdateTime" json:"tanggal"`
	StatusKerjasamaBPJS int8      `gorm:"column:STATUS_KERJASAMA_BPJS" json:"status_kerjasama_bpjs"`
	KodeFaskesBPJS      *string   `gorm:"column:KODE_FASKES_BPJS" json:"kode_faskes_bpjs"`
	Status              *int8     `gorm:"column:STATUS;default:1" json:"status"`
}

func (Faskes) TableName() string {
	return "master.faskes"
}
