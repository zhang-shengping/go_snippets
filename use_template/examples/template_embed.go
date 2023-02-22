package examples

import (
	"embed"
	"fmt"
	"os"
	"path"
	"text/template"
	"time"
)

//https://medium.com/@leo_hetsch/using-gos-embed-package-to-build-a-small-webpage-6175953fccea

//go:embed templates/*
var templates embed.FS

func EmbedTmpl() {

	teaPot := Product{
		Name:            "Red Tea Pot",
		Description:     "250ml",
		Price:           19.99,
		ShippingDate:    time.Now(),
		Sale:            true,
		SaleImagePath:   []string{"test_image"},
		MyFunc:          Myfunc,
		ShippingOptions: []string{"Extra Priority", "Normal", "Low Priority"},
	}

	hd := path.Join("templates", "header.html")
	ft := path.Join("templates", "footer.html")
	pd := path.Join("templates", "product.html")

	tmpl, err := template.ParseFiles(pd, hd, ft)
	// tmpl, err := template.ParseFiles(hd, ft, pd)
	if err != nil {
		panic(err)
	}

	// var tpl bytes.Buffer
	// if err := tmpl.Execute(&tpl, teaPot); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(tpl.String())

	if err := tmpl.Execute(os.Stdout, teaPot); err != nil {
		panic(err)
	}

	// 在 {{ define "header" }} block 里面的，代码不能显示出来。
	// if err = tmpl.ExecuteTemplate(os.Stdout, "header.html", teaPot); err != nil {
	// 	panic(err)
	// }
	// tmpl.ExecuteTemplate(os.Stdout, "product.html", teaPot)
	fmt.Println("test")
}
