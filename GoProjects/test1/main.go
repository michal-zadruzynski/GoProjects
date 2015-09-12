package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var persons = map[string]string{
	"870406": "Michał",
	"891030": "Paulina",
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getPersons(persons), r.URL.Path[1:])
}

func main() {
	fmt.Println("GoLang absolute begin")
	rand.Seed(time.Now().Unix())
	fmt.Printf("My favourite number is not %d\n", rand.Intn(10))

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func getPersons(pers map[string]string) string {
	var names string
	for _, v := range pers {
		names += v + "\n"
	}
	return names
}
