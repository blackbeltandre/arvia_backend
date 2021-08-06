package transaction

import (
	"database/sql"
	"fmt"
)
type Store interface {
	 CreateTransaction(tgl_transaksi,penjualan,qty,harga_jual,lokasi string) error
	 GetTransaction() ([]*Transaction, error)
	 GetLastStocks() ([]*LastStock, error)
	 GetStokAwal() ([]*StokAwal, error)

}
type DbStore struct {
	Db *sql.DB
}

func Regis(storeDB *DbStore) {
	store = storeDB
}

var store Store

func (store *DbStore) CreateTransaction(tgl_transaksi,penjualan,qty,harga_jual,lokasi string) error {
	_, err := store.Db.Query(
		"INSERT INTO transactions(tgl_transaksi,penjualan,qty,harga_jual,lokasi) VALUES (?,?,?,?,?)",
		 tgl_transaksi,penjualan,qty,harga_jual,lokasi)
	return err
}

//insert data transaksi
func (store *DbStore) GetTransaction() ([]*Transaction, error) {
	rows, err := store.Db.Query("select id,tgl_transaksi,penjualan,qty,harga_jual,lokasi from transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transactionList := []*Transaction{}
	for rows.Next() {
		transaction := &Transaction{}
		if err := rows.Scan(&transaction.IDTransaksi, &transaction.TglTransaksi, &transaction.Penjualan, &transaction.Qty, &transaction.HargaJual,&transaction.Lokasi); err != nil {

			return nil, err
		}
		transactionList = append(transactionList, transaction)
		fmt.Println(transaction)
	}
	return transactionList, nil
}


//data stok awal
func (store *DbStore) GetStokAwal() ([]*StokAwal, error) {
	rows, err := store.Db.Query("select * from stoks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stokawalList := []*StokAwal{}
	for rows.Next() {
		stokawal := &StokAwal{}
		if err := rows.Scan(&stokawal.IDStok, &stokawal.TglBeli, &stokawal.KdBrg, &stokawal.Qty, &stokawal.Harga,&stokawal.Lokasi); err != nil {

			return nil, err
		}
		stokawalList = append(stokawalList, stokawal)
		fmt.Println(stokawal)
	}
	return stokawalList, nil
}

//data stok akhir
func (store *DbStore) GetLastStocks() ([]*LastStock, error) {
	rows, err := store.Db.Query("select * from last_stok")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	laststokList := []*LastStock{}
	for rows.Next() {
		laststok := &LastStock{}
		if err := rows.Scan(&laststok.Sisa, &laststok.LabaKotor, &laststok.HargaJual, &laststok.Qty, &laststok.Penjualan,&laststok.Lokasi,&laststok.KdBrg,&laststok.StokAwal,&laststok.LokAwal,&laststok.ModalAwal); err != nil {

			return nil, err
		}
		laststokList = append(laststokList, laststok)
		fmt.Println(laststok)
	}
	return laststokList, nil
}
