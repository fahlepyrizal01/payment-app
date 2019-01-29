package pengguna

import (
	context "context"
	"database/sql"
	"fmt"
)

func (t *PenggunaData) getOnePengguna(ctx context.Context, db *sql.DB) (int64, string, string,
	string, string, float32, string, error) {
	var id int64 = 0
	var nama string = ""
	var telp string = ""
	var alamat string = ""
	var email string = ""
	var balance float32 = 0
	var url string = ""
	query := `SELECT id_pengguna, nama_pengguna, nomor_telepon, alamat, 
		email, balance, url_foto_profil FROM payment_db.pengguna 
		WHERE id_pengguna = $1`

	err := db.QueryRowContext(ctx, fmt.Sprintf(query), t.IdPengguna).Scan(&id, &nama,
		&telp, &alamat, &email, &balance, &url)

	if err != nil {
		return id, nama, telp, alamat, email, balance, url, err
	}
	return id, nama, telp, alamat, email, balance, url, nil
}

func (t *PenggunaData) loginPengguna(ctx context.Context, db *sql.DB) (int64, error) {
	var id int64 = 0
	query := `SELECT id_pengguna FROM payment_db.pengguna 
	 	WHERE email = $1 AND password = $2`

	err := db.QueryRowContext(ctx, fmt.Sprintf(query), t.Email, t.Password).Scan(&id)

	if err != nil {
		return id, err
	}
	return id, nil
}

func (t *PenggunaData) registerPengguna(ctx context.Context, db *sql.DB) (string, error) {
	var email string = ""
	query := `INSERT INTO payment_db.pengguna (
		nama_pengguna, password, nomor_telepon, alamat, email, balance, url_foto_profil) 
		VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING email`

	err := db.QueryRowContext(ctx, fmt.Sprintf(query), t.NamaPengguna,
		t.Password, t.NomorTelepon, t.Alamat, t.Email, t.Balance, t.UrlFotoProfil).Scan(&email)

	if err != nil {
		return email, err
	}
	return email, nil
}
