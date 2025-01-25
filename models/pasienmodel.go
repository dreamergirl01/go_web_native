package models

import (
	"database/sql"
	"fmt"
	"go_web_native/config"
	"go_web_native/entities"
	"time"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	//koneksi ke database
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PasienModel{
		conn: conn,
	}
}

//untuk mengambil semua data pasien
func (p *PasienModel) FindAll() ([]entities.Pasien, error) {
	rows, err := p.conn.Query("select * from pasien")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	//untuk menampung data dari iterasi rows, err := p.conn.Query("select * from pasien")
	var dataPasien []entities.Pasien
	for rows.Next(){
		//menampung satu data pasien
		var pasien entities.Pasien
		rows.Scan(&pasien.ID, 
			&pasien.NamaLengkap, 
			&pasien.NIK, 
			&pasien.JenisKelamin, 
			&pasien.TempatLahir, 
			&pasien.TanggalLahir, 
			&pasien.Alamat,
			&pasien.NoHp)

			if pasien.JenisKelamin == "1" {
				pasien.JenisKelamin = "Laki-Laki"
			}else{
				pasien.JenisKelamin = "Perempuan"
			}

			tgl_lahir, _ := time.Parse("2006-01-02", pasien.TanggalLahir)
			pasien.TanggalLahir = tgl_lahir.Format("02-01-2006")  //format time pada golang

		dataPasien = append(dataPasien, pasien)
	}
	return dataPasien, nil
}

// proses menyimpan data ke database
// create menerima parameter dengan tipe struct pasien
func (p *PasienModel) Create(pasien entities.Pasien) bool {
	result, err := p.conn.Exec("insert into pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values(?,?,?,?,?,?,?)",
		pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
