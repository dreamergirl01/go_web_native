package pasiencontroller

import (
	"go_web_native/entities"
	"go_web_native/libraries"
	"go_web_native/models"
	"html/template"
	"net/http"
)

var validation = libraries.NewValidation()
var PasienModel = models.NewPasienModel()

func Index(response http.ResponseWriter, request *http.Request) {
	//panggil method findAll yang ada pada pasienmodel
	pasien, _ := PasienModel.FindAll()
	data := map[string]interface{}{
		"pasien": pasien,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		//tampilkan data
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	}else if request.Method == http.MethodPost{
		//proses database
		request.ParseForm()

		// meneriima inputan dari form yang sudah dibuat (views/pasien/add.html),kemudian inputan yang sudah diterima akan disimpan pada NamaLengkap di struc pasien yang sudah dibuat pada entities pasien
		var pasien entities.Pasien
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)

		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors
		}else{
			data["pesan"] = "Data Pasien Berhasil Disimpan"
			//proses simpan data ke database
			PasienModel.Create(pasien)
		}
		
		//setelah proses simpan data ke database telah berhasil kembalikan view untuk tambah data
		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
