package pengguna

import context "context"

type Pengguna_server struct {
}

func (t *Pengguna_server) OnePengguna(ctx context.Context, in *PenggunaData) (*PenggunaData, error) {
	id, nama, telp, alamat, email, balance, url, err := in.getOnePengguna(ctx, dbPool)
	if err != nil {
		return &PenggunaData{}, err
	}
	in.IdPengguna = id
	in.NamaPengguna = nama
	in.NomorTelepon = telp
	in.Alamat = alamat
	in.Email = email
	in.Balance = balance
	in.UrlFotoProfil = url
	return in, nil
}

func (t *Pengguna_server) LoginPengguna(ctx context.Context, in *PenggunaData) (*PenggunaData, error) {
	id, nama, err := in.loginPengguna(ctx, dbPool)
	if err != nil {
		return &PenggunaData{}, err
	}
	in.IdPengguna = id
	in.NamaPengguna = nama
	return in, nil
}

func (t *Pengguna_server) RegisterPengguna(ctx context.Context, in *PenggunaData) (*PenggunaData, error) {
	id, err := in.registerPengguna(ctx, dbPool)
	if err != nil {
		return &PenggunaData{}, err
	}
	in.IdPengguna = id
	return in, nil
}
