package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 50)

//Fonction qui charge le fichier dans un slice (tous les mots)
func Load(filename string) error {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

//Fonction qui prend un mot aléatoirement depuis le slice
func PickWord() string {
	//Il n'y a pas de véritable aléatoire en informatique
	// le rand.Seed attend une racine, si on met toujours la meme racine on aura la meme séquence
	// Idéal en multijoueur afin que chaque joueur joue sur le meme mot par exemple
	rand.Seed(time.Now().Unix()) // On prend le nombre de millisecond depuis le 01/01/1970 pour que ça soit toujours différent
	i := rand.Intn(len(words))
	return words[i]
}
