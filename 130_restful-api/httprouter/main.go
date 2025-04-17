package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func GetProductsHandler(wrt http.ResponseWriter, req *http.Request, prms httprouter.Params) {
	keywords := req.URL.Query()["keyword"]
	if len(keywords) == 0 {
		fmt.Fprintf(wrt, "Delivering all products...\n")
	} else {
		fmt.Fprintf(wrt, "Delivering products with keyword(s) %v\n", keywords)
	}
}

var nextId = 1000

var products = make(map[int]Product)

func CreateProductHandler(w http.ResponseWriter, r *http.Request, pathPrms httprouter.Params) {
	// get and increment id
	id := nextId
	nextId++
	p := Product{
		ProductID: id,
		Name:      "Product Name",
	}
	products[p.ProductID] = p
	respond(w, p, http.StatusCreated)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request, pathPrms httprouter.Params) {
	productID, _ := strconv.Atoi(pathPrms.ByName("id"))
	p, exists := products[productID]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	respond(w, p, http.StatusOK)
}

// respond formats the response based on the Accept header
func respond(w http.ResponseWriter, data Product, httpStatus int) {
	w.WriteHeader(httpStatus)
	// current implementation always encodes to JSON
	var err error
	switch {
	default:
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
	}
	if err != nil {
		http.Error(w, "Encoding failed", http.StatusInternalServerError)
	}
}

// Product represents the API response structure
type Product struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
}

func main() {
	fmt.Println("🚀 Starting server on :8080")
	router := httprouter.New()
	router.GET("/products", GetProductsHandler)
	router.POST("/products", CreateProductHandler)
	router.DELETE("/products/:id", DeleteProductHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
