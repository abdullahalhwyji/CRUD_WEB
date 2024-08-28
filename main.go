package main

import (
	"log"
	"net/http"

	"github.com/abdullahalhwyji/crud_web_golang/config"
	"github.com/abdullahalhwyji/crud_web_golang/controllers/catagoriescontroller"
	"github.com/abdullahalhwyji/crud_web_golang/controllers/homecontroller"
	"github.com/abdullahalhwyji/crud_web_golang/controllers/productcontroller"
)

func main() {
	config.ConnectDB()

	// 1. homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//2. catagories
	http.HandleFunc("/categories", catagoriescontroller.Index)
	http.HandleFunc("/categories/add", catagoriescontroller.Add)
	http.HandleFunc("/categories/edit", catagoriescontroller.Edit)
	http.HandleFunc("/categories/delete", catagoriescontroller.Delete)

	//3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server is runing on port 8000 ...")

	http.ListenAndServe(":8000", nil)
}
