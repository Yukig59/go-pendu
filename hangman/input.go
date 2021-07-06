package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {

	validate := false
	for !validate {
		fmt.Print("Quelle est votre lettre ? \n")
		guess, err = reader.ReadString('\n') //Lire jusqu'a "entrée"
		//Early returns
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess) // Vire les espaces devant et derriere

		if len(guess) != 1 {
			fmt.Printf("Saisie incorrecte lettres entrées : %v, longueur : %v caractères\nVotre lettre ne doit comporter qu'un seul caractère. Veuillez réessayer\n", guess, len(guess))
			continue // Redemarre la boucle
		}
		validate = true
	}
	Clear()
	return guess, nil

}
