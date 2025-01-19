package entities

//menampung data pasien yang diambil dari database
type Pasien struct{
	ID int64
	NamaLengkap string
	NIK string
	JenisKelamin string
	TempatLahir string
	TanggalLahir string
	Alamat string
	NoHp string
}