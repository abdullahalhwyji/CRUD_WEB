package productcontroller

import (
	"net/http"
	"text/template"

	"github.com/abdullahalhwyji/crud_web_golang/models/productmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.Getall()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {

}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
