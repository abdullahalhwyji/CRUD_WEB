package main

import (
	"log"
	"net/http"

	"github.com/abdullahalhwyji/crud_web_golang/config"
)

func main() {
	config.ConnectDB()
	log.Println("Server is runing on port 8000 ...")

	http.ListenAndServe(":8000", nil)
}
