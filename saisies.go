package saisies

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

// Input affiche un message et renvoie la chaîne saisie par l'utilisateur
func Input(mess string) string {
	fmt.Print(mess)
	scanner.Scan()
	return scanner.Text()
}
