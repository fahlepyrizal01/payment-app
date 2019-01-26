package transaksi

import (
	context "context"
	"database/sql"
	"fmt"
)

type transaksiDataModel struct {
	IDTransaksi    int64
	IDPengirim     int64
	IDPenerima     int64
	Nominal        float32
	WaktuTransaksi string
	NamaPengirim   string
	NamaPenerima   string
}

func (t *transaksiDataModel) getAllTransaksi(ctx context.Context, db *sql.DB) ([]TransaksiData, error) {
	transaksiDatas := []TransaksiData{}

	query := `SELECT a.id_transaksi, a.id_pengirim, a.id_penerima, a.nominal, a.waktu_transaksi,
		b.nama_pengguna AS nama_pengirim, c.nama_pengguna AS nama_penerima 
    	FROM payment_db.transaksi a
    	INNER JOIN payment_db.pengguna b ON a.id_pengirim = b.id_pengguna 
    	INNER JOIN payment_db.pengguna c ON a.id_penerima = c.id_pengguna
		WHERE id_penerima = $1 OR id_pengirim = $1 ORDER BY id_transaksi DESC;`

	rows, err := db.QueryContext(ctx, fmt.Sprintf(query), t.IDPengirim)
	if err != nil {
		return transaksiDatas, err
	}
	for rows.Next() {
		transaksiData := TransaksiData{}
		rows.Scan(&transaksiData.IdTransaksi, &transaksiData.IdPengirim,
			&transaksiData.IdPenerima, &transaksiData.Nominal, &transaksiData.WaktuTransaksi,
			&transaksiData.NamaPengirim, &transaksiData.NamaPenerima)
		transaksiDatas = append(transaksiDatas, transaksiData)
	}
	rows.Close()
	return transaksiDatas, nil
}

func (t *TransaksiData) addTransaksi(ctx context.Context, db *sql.DB) (int64, error) {
	var id int64 = 0
	query := `INSERT INTO payment_db.transaksi (
		id_pengirim,id_penerima,nominal,waktu_transaksi) 
		VALUES ($1,$2,$3,$4) RETURNING id_transaksi`

	err := db.QueryRowContext(ctx, fmt.Sprintf(query), t.IdPengirim,
		t.IdPenerima, t.Nominal, t.WaktuTransaksi).Scan(&id)

	if err != nil {
		return id, err
	}

	queryUpdate := `UPDATE payment_db.pengguna SET balance = (CASE 
		WHEN id_pengguna = $1 THEN balance-$3 
		WHEN id_pengguna = $2 THEN balance+$3
		ELSE balance END) WHERE id_pengguna IN ($1, $2);`
	db.QueryRowContext(ctx, fmt.Sprintf(queryUpdate), t.IdPengirim, t.IdPenerima, t.Nominal)

	return id, nil
}
