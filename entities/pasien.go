package entities

//menampung data pasien yang diambil dari database
type Pasien struct{
	ID int64
	NamaLengkap string `validate:"required" label:"Nama lengkap"`
	NIK string `validate:"required"`
	JenisKelamin string `validate:"required" label:"Jenis Kelamin"`
	TempatLahir string `validate:"required" label:"Tempat Lahir"`
	TanggalLahir string `validate:"required" label:"Tanggal Lahir"`
	Alamat string `validate:"required"`
	NoHp string `validate:"required" label:"No Hp"`
}