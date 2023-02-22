package examples

import (
	"log"
	"os"
	"path"
	"text/template"
)

func Parsestring() {
	td := Todo{
		"Test template",
		"Let's test a templdate to see the magic.",
	}

	t, err := template.New("todos").Parse("You have a task named \"{{ .Name }}\" with description: \"{{ .Description }}\"\n")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}

	td = Todo{
		"Go",
		"Contribete to any Go project",
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}

// set variable
const tmpl = `
{{$name := .Name}}
The name is {{.Name}}.
{{range .Emails}}
   Name is {{$name}} His email id is {{.}}
{{end}}
`

// use global context sign $
const tmpl2 = `
The name is {{.Name}}.
{{range .Emails}}
   Name is {{$.Name}} His email id is {{.}}
{{end}}
`

func Parsetmpl() {
	person := Person{
		Name:   "pzhang",
		Emails: []string{"zhangshengping2012@hotmail.com", "zhangshengping2012@gmail.com"},
	}

	t := template.New("Person template")
	// t, err := t.Parse(tmpl)
	t, err := t.Parse(tmpl2)
	if err != nil {
		log.Fatal("Parse: ", err)
	}

	err = t.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}

func Parsetmplfile() {
	lp := path.Join("templates", "layout.html")
	idx_lp := path.Join("templates", "index.html")

	templates, err := template.ParseFiles(lp, idx_lp)
	if err != nil {
		log.Fatal("Parse Files", err)
	}
	templates.ExecuteTemplate(os.Stdout, "layout", nil)
}
