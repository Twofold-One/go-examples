package templates

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"time"
)

type Product struct {
	Name            string
	Price           string
	Description     string
	ShippingDate    time.Time
	Sale            bool
	SaleImagePath   []string
	MyFunc          func(string, string) string
	ShippingOptions []string
	Notes           [][]int
}

var notes = [][]int{{1, 2, 3}, {4, 5, 6}}

var teaPot Product = Product{Name: "Red Tea Pot", Price: "20.00", Description: "Test", ShippingDate: time.Now(), Sale: true, SaleImagePath: []string{"img path"}, MyFunc: Bar, ShippingOptions: []string{"Extra Priority", "Normal", "Low Priority"}, Notes: notes}
var PageTitle string = "Shop"
var capitalizeFirstLetter = func(text string) string { return strings.Title(text) }

func (p Product) Foo() string {
	return "Foo"
}

func (p Product) Bar(test string) string {
	return fmt.Sprintf("Bar : %s", test)
}

func Bar(a string, b string) string {
	buf := bytes.NewBufferString(a)
	buf.WriteString(b)
	return buf.String()
}

// main func of package
func TemplatesExample() {
	http.HandleFunc("/red-tea-pot", redTeaPotHandler)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}

func redTeaPotHandler(w http.ResponseWriter, r *http.Request) {
	funcs := template.FuncMap{"capitalizeFirstLetter": capitalizeFirstLetter}
	tmpl, err := template.New("product.html").Funcs(funcs).ParseFiles("/home/twofold_one/GitProjects/go/go-examples/basic/templates/views/product.html", "/home/twofold_one/GitProjects/go/go-examples/basic/templates/views/header.html", "/home/twofold_one/GitProjects/go/go-examples/basic/templates/views/footer.html")
	if err != nil {
		http.Error(w, "Something went wrong parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "product.html", teaPot)

}
