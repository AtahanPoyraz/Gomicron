package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/AtahanPoyraz/cmd"
	dbproc"github.com/AtahanPoyraz/db/dbases"

	"github.com/gorilla/mux"
)

// Logger struct oluşturuyoruz.
type Products struct {
	l *log.Logger
}
type KeyProduct struct {
}

// Yeni Productlar Bu Fonksiyon Ile Products Struct ı Kullanılarak Yeni Bir Product Oluşturuluyor.
func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l}
}

// Veri Dosyasından Alınan Productlar Json Nesnesine Cevriliyor.
func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("%s[METHOD]%s: %sGET%s >> %s\n",cmd.BGREEN_WHITE, cmd.TRESET, cmd.TGREEN, cmd.TRESET, r.URL.Path)
	lp := dbproc.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError) //Geri Bildirim
		p.l.Printf("%s[ERROR]%s: Unable to marshal JSON: %v",cmd.BRED_WHITE, cmd.TRESET, err) //Geri Bildirim

	}
	for _, prod := range lp {
		p.l.Printf("%s[INFO]%s : %v", cmd.BGRAY_BLACK, cmd.TRESET, *prod)
	}
}

// Yeni Product Olusturmak Istendiginde Struct Tanımlanıp Json Formatına Donusturuluyor.
func (p Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("%s[METHOD]%s: %sPOST%s >> %s\n",cmd.BYELLOW_WHITE, cmd.TRESET, cmd.TYELLOW, cmd.TRESET, r.URL.Path)

	prod := r.Context().Value(KeyProduct{}).(dbproc.Product)
	dbproc.AddProduct(&prod)

	p.l.Printf("%s[INFO]%s : %v\n", cmd.BGRAY_BLACK, cmd.TRESET, prod)
	p.l.Println("OBJECT ADDING")
}

// Var Olan bir Product ı Guncelleniyor.
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("%s[METHOD]%s: %sPUT%s >> %s\n",cmd.BBLUE_WHITE, cmd.TRESET, cmd.TDBLUE, cmd.TRESET, r.URL.Path)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		p.l.Printf("%s[ERROR]%s : Unable to Convert ID: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) //Geri Bildirim
		return
	}
	prod := r.Context().Value(KeyProduct{}).(dbproc.Product)

	err = dbproc.UpdateProduct(id, &prod)
	if err == dbproc.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		p.l.Printf("%s[ERROR]%s : Product Not Found: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) 

		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		p.l.Printf("%s[ERROR]%s : Product Not Found: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) 

		return
	}
	p.l.Printf("%s[INFO]%s : %v", cmd.BGRAY_BLACK, cmd.TRESET, prod)
}

func (p Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("%s[METHOD]%s: %sDELETE%s >> %s\n",cmd.BORANGE_WHITE, cmd.TRESET, cmd.TORANGE, cmd.TRESET, r.URL.Path)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		p.l.Printf("%s[ERROR]%s : Unable to Convert ID: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) //Geri Bildirim

		return
	}
	prod := r.Context().Value(KeyProduct{}).(dbproc.Product)

	err = dbproc.DeleteProduct(id, &prod) 
	if err == dbproc.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		p.l.Printf("%s[ERROR]%s : Product Not Found: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) 

		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		p.l.Printf("%s[ERROR]%s : Product Not Found: (%e)", cmd.BRED_WHITE, cmd.TRESET, err) 
		return
	}

	p.l.Printf("%s[INFO]%s: Product Delete Proccess been Succsessful\n", cmd.BGRAY_BLACK, cmd.TRESET)
}


func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := dbproc.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Printf("%s[ERROR]%s: Deserializing Data: %v", cmd.BRED_WHITE, cmd.TRESET , err)
			http.Error(rw, "Error Reading Product", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Printf("%s[ERROR]%s : Validating Product : %v", cmd.BRED_WHITE, cmd.TRESET, err)
			http.Error(rw, "Error Validating Product", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}