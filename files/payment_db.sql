DROP DATABASE IF EXISTS payment_db CASCADE;

CREATE DATABASE payment_db;

CREATE TABLE payment_db.Pengguna (
    id_pengguna SERIAL PRIMARY KEY,
    nama_pengguna string NOT NULL,
    password string NOT NULL,
    nomor_telepon string NOT NULL,
    alamat string NOT NULL,
    email string NOT NULL,
    balance FLOAT NOT NULL,
    jenis_kelamin string NOT NULL,
    tanggal_lahir string NOT NULL
);

CREATE TABLE payment_db.Transaksi (
    id_transaksi SERIAL PRIMARY KEY,
    id_pengirim INT64 NOT NULL REFERENCES payment_db.Pengguna (id_pengguna),
    id_penerima INT64 NOT NULL REFERENCES payment_db.Pengguna (id_pengguna),
    nominal FLOAT NOT NULL,
    waktu_transaksi string NOT NULL
);

CREATE TABLE payment_db.Isi_saldo (
    id_isi_saldo SERIAL PRIMARY KEY,
    id_penguna INT64 NOT NULL REFERENCES payment_db.Pengguna (id_pengguna),
    nominal_isi_saldo INT NOT NULL,
    waktu_isi_saldo string NOT NULL
);