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

type entree struct{
	Entree string
	Dish []dish
}
type menu []entree

type restaurant struct{
	Name string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init(){
	tpl=template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	R:= restaurants{
		restaurant{
			Name: "Diner",
			Menu: menu{
				entree{
					Entree:"Breakfast",
					Dish: []dish{
						dish{
							Name: "eggs",
							Description: "sunny-side",
							Price: 3.35,
						},
						dish{
							Name: "bacon",
							Description:"pork belly",
							Price: 2.51,
						},
						dish{
							Name: "coffee",
							Description: "black",
							Price: 1.32,
						},
					},
				},
				entree{
					Entree:"Lunch",
					Dish: []dish{
						dish{
							Name: "sandwich",
							Description: "bread and ham",
							Price: 5.11,
						},
						dish{
							Name: "burger",
							Description:"bread and ground beef",
							Price: 8.97,
						},
						dish{
							Name: "soda",
							Description: "carbonated water",
							Price: 2.72,
						},

					},
				},
				entree{
					Entree:"Dinner",
					Dish: []dish{
						dish{
							Name: "steak",
							Description: "seared beef",
							Price: 14.21,
						},
						dish{
							Name: "salad",
							Description:"green leaves doused in ranch",
							Price: 20.21,
						},
						dish{
							Name: "Whiskey",
							Description: "neat",
							Price: 8.22,
						},

					},
				},
			},
		},
		restaurant{
			Name: "Denny's",
			Menu: menu{
				entree{
					Entree:"Breakfast",
					Dish: []dish{
						dish{
							Name: "two eggs",
							Description: "sunny-side",
							Price: 3.14,
						},
						dish{
							Name: " extra bacon",
							Description:"pork belly",
							Price: 2.24,
						},
						dish{
							Name: "decafe coffee",
							Description: "black",
							Price: 1.04,
						},
					},
				},
				entree{
					Entree:"Lunch",
					Dish: []dish{
						dish{
							Name: " egg sandwich",
							Description: "bread and ham and egg",
							Price: 5.02,
						},
						dish{
							Name: "pancake burger",
							Description:"pancakes covering ground beef",
							Price: 8.910,
						},
						dish{
							Name: "soda",
							Description: "carbonated water",
							Price: 2.32,
						},

					},
				},
				entree{
					Entree:"Dinner",
					Dish: []dish{
						dish{
							Name: "t-bone steak",
							Description: "seared beef in the shape of a T",
							Price: 14.20,
						},
						dish{
							Name: "salad and nuts",
							Description:"green leaves doused in ranch with salted walnuts sprinkled over",
							Price: 20.21,
						},
						dish{
							Name: "ice cream shake",
							Description: "ice cream blended smoothie",
							Price: 8.01,
						},

					},
				},
			},
		},
	}
	err:= tpl.Execute(os.Stdout,R)
	if err !=nil{
		log.Fatalln(err)
	}
}

