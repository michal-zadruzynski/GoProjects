package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("GoLang absolute begin")
	rand.Seed(time.Now().Unix())
	fmt.Printf("My favourite number is not %d", rand.Intn(10))
}
