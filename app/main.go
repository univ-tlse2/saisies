package main

import (
	"fmt"

	. "github.com/univ-tlse2/saisies"
)

func main() {
	phrase := Input("Entrez une phrase : ")
	fmt.Println(phrase)
}
