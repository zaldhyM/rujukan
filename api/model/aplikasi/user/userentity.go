package user

type User struct {
	ID                     uint
	USERNAME               string
	PASSWORD               string
	NAMA                   string
	NIK                    string
	ID_FASKES              uint
	MASA_AKTIF             string
	TOKEN                  string
	ROLE                   string
	STATUS_TELEGRAM        string
	ID_TELEGRAM            string
	TERAKHIR_UBAH_PASSWORD string
	STATUS                 int16
}

// override nama tabel default
func (User) TableName() string {
	return "user"
}
