package main

import(
	"log"
	"os"
	"text/template"
)

type dish struct{
	Name, Description, Hour string
	Price	float64
}

type dishes []dish

var tpl *template.Template

func init(){
	tpl=template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	menu := dishes{
		dish{
			Name:"eggs",
			Description:"sunny-side up eggs",
			Hour:"breakfast",
			Price:2.34,
		},
		dish{
			Name:"sandwich",
			Description:"meat and luttuce wrapped in bread",
			Hour:"lunch",
			Price:5.00,
		},
		dish{
			Name:"eggs",
			Description:"sunny-side up eggs",
			Hour:"breakfast",
			Price:2.34,
		},
	}
	err:= tpl.Execute(os.Stdout,menu)
	if err !=nil{
		log.Fatalln(err)
	}
}

