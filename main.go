package main

import (
	"fmt"
	"os"
	"yuki.go/pendu/dictionary"
	"yuki.go/pendu/hangman"
)

func main() {
	hangman.Clear()
	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Ne peut pas charger le dicitonnaire : %v\n", err)
		os.Exit(1)
	}
	g := hangman.New(8, dictionary.PickWord())
	hangman.DrawWelcome()
	guess := "a" // Lettre proposée
	for {
		hangman.Draw(g, guess) //On dessine le jeu
		switch g.State {
		case "won", "lost":
			os.Exit(0) // 0 = tout s'est bien passé pour l'OS on quitte la boucle
		}
		//On continue de jouer tant qu'on a ni gagné ni perdu
		l, err := hangman.ReadGuess() //l = la lettre qui est saisie
		if err != nil {
			fmt.Printf("Ne peux pas lire votre saisie: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
