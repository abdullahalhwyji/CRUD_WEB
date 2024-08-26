package main

import (
	"log"
	"net/http"

	"github.com/abdullahalhwyji/crud_web_golang/config"
	"github.com/abdullahalhwyji/crud_web_golang/controllers/catagoriescontroller"
	"github.com/abdullahalhwyji/crud_web_golang/controllers/homecontroller"
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

	log.Println("Server is runing on port 8000 ...")

	http.ListenAndServe(":8000", nil)
}
