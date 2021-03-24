package saisies

import (
	"bufio"
	"fmt"
	"os"
)

// Input affiche un message et renvoie la chaîne saisie par l'utilisateur
func Input(mess string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(mess)
	scanner.Scan()
	return scanner.Text()
}
