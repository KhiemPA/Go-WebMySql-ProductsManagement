package main

import (
	"fmt"
	"net/http"
	"strconv"

	"text/template"
)


var tpl = template.Must(template.ParseGlob("template/*"))


func main() {
	fmt.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	http.HandleFunc("/search", Search)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	products := GetProducts()
	tpl.ExecuteTemplate(w, "Home", products)
}

func Create(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	product := Product{}
	product.ProductName = r.FormValue("productname")

	// convert string thanh float 
	// price, err := strconv.ParseFloat(r.FormValue("productprice"), 64)
	// if err != nil {
	// 	fmt.Println("Yeu cau nhap dung so")
	// 	panic(err.Error())
	// }
	var price float64
	priceStr := r.FormValue("productprice")
	fmt.Sscanf(priceStr, "%d", &price)
	product.ProductPrice = float32(price)
	product.ProductDescription = r.FormValue("productdescription")

	InsertProduct(product)
	products := GetProducts()
	tpl.ExecuteTemplate(w, "Home", products)
}

func Alter(w http.ResponseWriter, r *http.Request) {
	product := Product{}
	var productId int
	productIdStr := r.FormValue("id")
	fmt.Sscanf(productIdStr, "%d", &productId)
	product.ProductId = productId
	product.ProductName = r.FormValue("productname")
	var err error
	var price float64
	price, err = strconv.ParseFloat(r.FormValue("productprice"), 64)
	if err != nil {
		panic(err.Error())
	}
	product.ProductPrice = float32(price)
	product.ProductDescription = r.FormValue("productdescription")

	UpdateProduct(product)
	products := GetProducts()
	tpl.ExecuteTemplate(w, "Home", products)
}


func Update(w http.ResponseWriter, r *http.Request) {

	id, err:= strconv.Atoi(r.FormValue("id"))
	if err != nil {
		panic(err.Error())
	}
	product := GetProductById(id)
	tpl.ExecuteTemplate(w, "Update", product)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	// lấy thông tin giống như hàm hiển thị update

	id, err:= strconv.Atoi(r.FormValue("id"))
	if err != nil {
		panic(err.Error())
	}
	product := GetProductById(id)
	DeleteProduct(product)
	products := GetProducts()
	tpl.ExecuteTemplate(w, "Home", products)
}

func View(w http.ResponseWriter, r *http.Request) {
	id, err:= strconv.Atoi(r.FormValue("id"))
	if err != nil {
		panic(err.Error())
	}
	product := GetProductById(id)
	fmt.Println(product)
	tpl.ExecuteTemplate(w, "View", product)
}

func Search(w http.ResponseWriter, r *http.Request) {
	productName := r.FormValue("productname")
	products := SearchProductByName(productName)
	
	tpl.ExecuteTemplate(w, "Home", products)
}


