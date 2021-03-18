package main

import(
	"net/http"
	"log"
	"WEB-INF/golang_todo/myapp"
	"github.com/urfave/negroni"
)

func main(){
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)
	log.Println("started app")
	err := http.ListenAndServe(":3000", n)
	if err != nil{
		panic(err)
	}
}
