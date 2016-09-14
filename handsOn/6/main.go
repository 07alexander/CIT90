package main

import(
	"log"
	"os"
	"text/template"
)

type hotel struct{
	Name, Address, City, Zip, Region string
}
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
type hotels []hotel

func main() {

	Hotels:=hotels{
		hotel{
			Name: "inn",
			Address:"555 N street",
			City:"Fresno",
			Zip:"90000",
			Region:"Central",
		},
		hotel{
			Name: "asdasd",
			Address:"232 N street",
			City:"Clovis",
			Zip:"90000",
			Region:"Northern",
		},
	}

	err:=tpl.Execute(os.Stdout, Hotels)
	if err != nil{
		log.Fatalln(err)
	}

}
