package transaksi

import context "context"

type Transaksi_server struct {
}

func (t *Transaksi_server) AllTransaksi(in *TransaksiData, stream TransaksiService_AllTransaksiServer) error {
	transaksiDatas, err := (&transaksiDataModel{IDPengirim: in.IdPengirim}).getAllTransaksi(stream.Context(), dbPool)
	if err != nil {
		return err
	}
	for _, transaksi := range transaksiDatas {
		stream.Send(&transaksi)
	}
	return nil
}

func (t *Transaksi_server) AddTransaksi(ctx context.Context, in *TransaksiData) (*TransaksiData, error) {
	id, err := in.addTransaksi(ctx, dbPool)
	if err != nil {
		return &TransaksiData{}, err
	}
	in.IdTransaksi = id
	return in, nil
}
