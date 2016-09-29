package main

import(
	"html/template"
	"os"
)

var tpl *template.Template
func init(){
	tpl =  template.Must(template.ParseGlob("index/*"))
}

func main() {
	tpl.ExecuteTemplate(os.Stdout,"About.asp",nil)
	tpl.ExecuteTemplate(os.Stdout,"Contact.jsp",nil)
	tpl.ExecuteTemplate(os.Stdout,"index.php",nil)
}
