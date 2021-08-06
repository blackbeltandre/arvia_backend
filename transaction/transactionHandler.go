package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)
//response header / meta
type Result struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
	Status  int         `:"status"`
}

//interface data last stok
type StokAwal struct {
	IDStok        string    `json:"id_stok" `
	TglBeli       string    `json:"tgl_beli"`
	KdBrg       string    `json:"kd_brg"`
	Qty  string `json:"qty"`
	Harga  string `json:"harga"`
	Lokasi       string    `json:"lokasi"`
	}


//interface data last stok
type LastStock struct {
	Sisa        string    `json:"sisa" `
	LabaKotor       string    `json:"laba_kotor"`
	HargaJual       string    `json:"harga_jual"`
	Qty  string `json:"qty"`
	Penjualan  string `json:"penjualan"`
	Lokasi  string `json:"lokasi"`
	KdBrg       string    `json:"kd_brg"`
	StokAwal       string    `json:"stok_awal"`
	LokAwal       string    `json:"lokasi_awal"`
	ModalAwal       string    `json:"modal_awal"`
	}
//interface data transaksi get & post	
type Transaction struct {
	IDTransaksi        string    `json:"id" `
	TglTransaksi  string `json:"tgl_transaksi"`
	Penjualan  string `json:"penjualan"`
	Qty  string `json:"qty"`
	HargaJual       string    `json:"harga_jual"`
	Lokasi       string    `json:"lokasi"`
	Modal       string    `json:"modal"`
	LabaKotor       string    `json:"laba_kotor"`
	}
	func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	requestBody, _ := ioutil.ReadAll(r.Body)
	var transaction = Transaction{}
	json.Unmarshal(requestBody, &transaction)
	
	//tadinya saya mau cek field yg dikirim tp tidak jadi ,saya lempar smwnya sebagai string dari front end dan juga saya tangkap sebagai string ,namun saya convert string ke integer dan konvert lagi dari integer ke string untuk proses insert data
	//extract body dan decode ,pilih yang mau di parsing atau sent all field yang ada di interface Transaction
	// values := map[string]string{"tgl_transaksi": transaction.TglTransaksi, "penjualan": transaction.Penjualan, "qty": transaction.Qty, "harga_jual": transaction.HargaJual, "lokasi": transaction.Lokasi}
	// fmt.Println(values)
	// var modal int = 20
 // 	modal_awalnya, err := strconv.Atoi(transaction.HargaJual)
 //    if err != nil {
 //        log.Fatal(err)
 //    }
	// modal_awalss := modal_awalnya - modal
	// modal_awals := strconv.Itoa(modal_awalss)
	//  fmt.Println(modal_awal)
	// var laba int = 100
 // 	laba_kotornya, err := strconv.Atoi(transaction.HargaJual)
 //    if err != nil {
 //        log.Fatal(err)
 //    }
 //    laba_kotorss := laba_kotornya - laba
 //    laba_kotors := strconv.Itoa(laba_kotorss)
 //    var modal_awal = modal_awals
	// var laba_kotor = laba_kotors
	// fmt.Println(laba_kotor)

	var tgl_transaksi = transaction.TglTransaksi
	var penjualan = transaction.Penjualan
	var qty = transaction.Qty
	var harga_jual = transaction.HargaJual
	var lokasi = transaction.Lokasi
	err = store.CreateTransaction(tgl_transaksi,penjualan,qty,harga_jual,lokasi)
	if err != nil {
		fmt.Println(err)
	}
	result := &Result{Result: "Process Success", Message: "Data Added", Status: 1}
	fmt.Println(result)
	status_ok, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(status_ok)
	json.NewEncoder(w).Encode(transaction)
}

func GetStockHandler(w http.ResponseWriter, r *http.Request) {
	if store == nil {
		result := &Result{Result: "SOMETHING WRONG", Message: "PLEASE REGISTER DB TO MAIN CLASS !", Status: 0}
		stokListBytes, err := json.Marshal(result)
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(stokListBytes)
		return
	}
	stokList, err := store.GetStokAwal()
	result := &Result{Result: stokList, Message: "Success Display Data", Status: 1}
	stokListBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(stokListBytes)

}


func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if store == nil {
		result := &Result{Result: "SOMETHING WRONG", Message: "PLEASE REGISTER DB TO MAIN CLASS !", Status: 0}
		transactionListBytes, err := json.Marshal(result)
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(transactionListBytes)
		return
	}
	transactionList, err := store.GetTransaction()
	result := &Result{Result: transactionList, Message: "Success Display Data", Status: 1}
	transactionListBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(transactionListBytes)

}

func GetLastStockHandler(w http.ResponseWriter, r *http.Request) {
	if store == nil {
		result := &Result{Result: "SOMETHING WRONG", Message: "PLEASE REGISTER DB TO MAIN CLASS !", Status: 0}
		laststokListBytes, err := json.Marshal(result)
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(laststokListBytes)
		return
	}
	laststokList, err := store.GetLastStocks()
	result := &Result{Result: laststokList, Message: "Success Display Data", Status: 1}
	laststokListBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(laststokListBytes)

}