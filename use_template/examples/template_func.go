package examples

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

var capitalizeFirstLetter = func(text string) string { return strings.Title(text) }

func TmplFunc() {

	teaPot := Product{
		Name:          "Red Tea Pot",
		Description:   "250ml",
		Price:         19.99,
		ShippingDate:  time.Now(),
		Sale:          true,
		SaleImagePath: []string{"test_image"},

		// templdate function use call action
		MyFunc:          Myfunc,
		ShippingOptions: []string{"Extra Priority", "Normal", "Low Priority"},
	}

	hd := path.Join("templates", "header.html")
	ft := path.Join("templates", "footer.html")
	pd := path.Join("templates", "product.html")

	funcs := template.FuncMap{"capitalizeFirstLetter": capitalizeFirstLetter}

	// https://stackoverflow.com/questions/49043292/error-template-is-an-incomplete-or-empty-template
	// not work: tmpl, err := template.New("random_name").Funcs(funcs).ParseFiles(pd, hd, ft)
	// Make sure the argument you pass to template.New is the base name of one of the files in the list you pass to ParseFiles
	// tmpl, err := template.New("header.html").Funcs(funcs).ParseFiles(pd, hd, ft)
	tmpl, err := template.New("product.html").Funcs(funcs).ParseFiles(pd, hd, ft)

	// The key of this map is the name that will be exposed into the template, the second argument is the function itself.
	// We can not directly call ParseFiles, because the functions have to be added to the template before parsing
	// tmpl, err := template.ParseFiles(pd, hd, ft)

	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, teaPot); err != nil {
		panic(err)
	}
	fmt.Println("end")
}
