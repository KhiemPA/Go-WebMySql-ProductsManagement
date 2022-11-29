package main

import (
	"database/sql"
	"fmt"

	_"github.com/go-sql-driver/mysql"
)

type Product struct {
	ProductId int
	ProductName string
	ProductPrice float32
	ProductDescription string
}

// tcp(localhost:3306)

func GetConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:16825ds5230@/test1")
	if err != nil {
		fmt.Println("Khong the ket noi toi DB")
		panic(err.Error())
	}
	return db
}

func GetProducts() []Product{
	// Có thể khai báo var db *sql.DB
	database := GetConnection()
	defer database.Close()
	// Co the khai bao var rows *sql.Rows
	var rows *sql.Rows
	var err error
	rows, err = database.Query("SELECT * FROM products ORDER BY idproduct DESC")
	if err != nil {
		fmt.Println("Bug in db query get all")
	}
	product := Product{}
	products := []Product{}

	for rows.Next() {
		var productId int
		var productName string
		var productPrice float32
		var productDescription string

		err = rows.Scan(&productId, &productName, &productPrice, &productDescription)
		if err != nil {
			panic(err.Error())
		}
		product.ProductId = productId
		product.ProductName = productName
		product.ProductPrice = productPrice
		product.ProductDescription = productDescription
		products = append(products, product)
	}
	return products
}

func GetProductById(productId int) Product {
	db := GetConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products WHERE idproduct =?", productId)
	if err != nil {
		fmt.Println("Bug in db query get by id")
		panic(err.Error())
	}

	product := Product{}

	for rows.Next() {
		var productId int
		var productName string
		var productPrice float32
		var productDescription string

		err = rows.Scan(&productId, &productName, &productPrice, &productDescription)
		if err != nil {
			panic(err.Error())
		}
		product.ProductId = productId
		product.ProductName = productName
		product.ProductPrice = productPrice
		product.ProductDescription = productDescription
	}
	return product
}

func SearchProductByName(productName string) []Product {
	db := GetConnection()
	defer db.Close()

	search, err := db.Prepare("SELECT * FROM products WHERE name LIKE ?;")

	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	products := []Product{}
	var rows *sql.Rows
	rows, err = search.Query("%" + productName +"%")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var productId int
		var productName string
		var productPrice float32
		var productDescription string

		err = rows.Scan(&productId, &productName, &productPrice, &productDescription)
		if err != nil {
			panic(err.Error())
		}
		product.ProductId = productId
		product.ProductName = productName
		product.ProductPrice = productPrice
		product.ProductDescription = productDescription
		products = append(products, product)
	}
	return products
}

func InsertProduct(product Product) {
	db := GetConnection()
	defer db.Close()
	// có thể khai báo var insert *sql.Stmt
	insert, err := db.Prepare("INSERT INTO products(name, price, description) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(product.ProductName, product.ProductPrice, product.ProductDescription)
}

func UpdateProduct(product Product) {
	db := GetConnection()
	defer db.Close()

	update, err := db.Prepare("UPDATE products SET name=?, price=?, description=? WHERE idproduct=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(product.ProductName, product.ProductPrice, product.ProductDescription, product.ProductId)
}

func DeleteProduct(product Product) {
	db := GetConnection()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM products WHERE idproduct=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(product.ProductId)
}
