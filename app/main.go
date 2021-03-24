package main

import (
	"fmt"

	"github.com/jaco60/saisies"
)

func main() {
	phrase := saisies.Input("Entrez une phrase : ")
	fmt.Println(phrase)
}
