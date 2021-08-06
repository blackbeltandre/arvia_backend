package main

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	 _ "github.com/go-sql-driver/mysql"
	"backend/transaction"
	"flag"
	"fmt"
)
var(
	port = "3004"
)
func newRouter() *mux.Router {
	r := mux.NewRouter()
	//get stok awal route
	r.HandleFunc("/transaction/get_stock", transaction.GetStockHandler).Methods("GET")

	//get all transaction route
	r.HandleFunc("/transaction/get_transaction", transaction.GetTransactionHandler).Methods("GET")

	//Transaction route (post json from node.js to api)
	r.HandleFunc("/transaction/insert_transaction", transaction.CreateTransactionHandler).Methods("POST")

	//Last stok route
	r.HandleFunc("/transaction/last_stock", transaction.GetLastStockHandler).Methods("GET")
	return r
}
func main() {
	flag.Parse()
	 db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/arvia")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	storeTransaction := &transaction.DbStore{Db: db}
	transaction.Regis(storeTransaction)
		r := newRouter()
	fmt.Println("Starting Arvia Backend at port : ", port)
	http.ListenAndServe(fmt.Sprintf(":%v",port), r)

}
