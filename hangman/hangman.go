package hangman

import (
	"os"
	"strings"
)

type Game struct {
	State        string   // Etat du jeu (perdu, gagné ...)
	Letters      []string //Le mot a deviner
	FoundLetters []string // L'affichage du mot avec les lettres a trouvées "B _ _ J _ _ R"
	UsedLetters  []string // Lettres jouées
	TurnsLeft    int      // Tours restants
	Points       int
	Indices      int
}

func New(turns int, word string) *Game {
	//Mot a deviner
	letters := strings.Split(strings.ToUpper(word), "")
	// Les lettres trouvées => il faut autant de "_" que de lettres dans le mot
	found := make([]string, len(letters)) //Slice à la main avec une capacité determinée
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	//On fabrique le struct du jeu et on le return
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
		Points:       0,
	}
	return g
}

// La lettre proposée est elle dans le mot à deviner ?
func LetterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

//Pour savoir si on a gagné on compare lettre à lettre les 2 tableaux (celui du mot à trouver et celui du mot trouvé)
func hasWon(letters []string, foundLetter []string) bool {
	for i := range letters {
		if letters[i] != foundLetter[i] {
			//Si il y'a la moindre différence alors on n'a pas gagné
			return false
		}
	}
	return true
}

// La lettre proposée, qui est bonne est affichée en clair à la place du "_"
func (g *Game) RevealLetter(guess string) {
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess //On affiche la lettre à la place du "_"
		}
	}
}

//Fonction qui test la lettre soumise par l'utilisateur et fait varier l'état du jeu
func (g *Game) MakeAGuess(guess string) {
	DrawWelcome()
	guess = strings.ToUpper(guess)
	l := 8
	if LetterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else {
		// On ajoute la lettre aux lettres utillisée
		g.UsedLetters = append(g.UsedLetters, guess)
		// Cas où la lettre est dans le mot à devinera
		if LetterInWord(guess, g.Letters) {
			g.State = "goodGuess"
			//On doit maintenant afficher en clair la lettre trouvée
			g.RevealLetter(guess)
			//Est ce que c'était la dernière lettre à trouver ? A ton gagné ?
			if hasWon(g.Letters, g.FoundLetters) {
				g.State = "won"
			}
		} else {
			g.State = "badGuess"
			g.TurnsLeft--
			drawTurns(l)
			l--
			if g.TurnsLeft <= 0 {
				g.State = "lost"
			}
		}
	}
}
