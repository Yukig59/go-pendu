package hangman

// TODO : NE PAS INTENTER CE FICHIER !!!
import "fmt"

// Fonction qui affiche la bannière d'accueil
// http://patorjk.com/software/taag/#p=display&f=Big&t=PENDU
func DrawWelcome(){
	fmt.Println(`          _____                   _____                   _____                   _____                   _____          
         /\    \                 /\    \                 /\    \                 /\    \                 /\    \         
        /::\    \               /::\    \               /::\____\               /::\    \               /::\____\        
       /::::\    \             /::::\    \             /::::|   |              /::::\    \             /:::/    /        
      /::::::\    \           /::::::\    \           /:::::|   |             /::::::\    \           /:::/    /         
     /:::/\:::\    \         /:::/\:::\    \         /::::::|   |            /:::/\:::\    \         /:::/    /          
    /:::/__\:::\    \       /:::/__\:::\    \       /:::/|::|   |           /:::/  \:::\    \       /:::/    /           
   /::::\   \:::\    \     /::::\   \:::\    \     /:::/ |::|   |          /:::/    \:::\    \     /:::/    /            
  /::::::\   \:::\    \   /::::::\   \:::\    \   /:::/  |::|   | _____   /:::/    / \:::\    \   /:::/    /      _____  
 /:::/\:::\   \:::\____\ /:::/\:::\   \:::\    \ /:::/   |::|   |/\    \ /:::/    /   \:::\ ___\ /:::/____/      /\    \ 
/:::/  \:::\   \:::|    /:::/__\:::\   \:::\____/:: /    |::|   /::\____/:::/____/     \:::|    |:::|    /      /::\____\
\::/    \:::\  /:::|____\:::\   \:::\   \::/    \::/    /|::|  /:::/    \:::\    \     /:::|____|:::|____\     /:::/    /
 \/_____/\:::\/:::/    / \:::\   \:::\   \/____/ \/____/ |::| /:::/    / \:::\    \   /:::/    / \:::\    \   /:::/    / 
          \::::::/    /   \:::\   \:::\    \             |::|/:::/    /   \:::\    \ /:::/    /   \:::\    \ /:::/    /  
           \::::/    /     \:::\   \:::\____\            |::::::/    /     \:::\    /:::/    /     \:::\    /:::/    /   
            \::/____/       \:::\   \::/    /            |:::::/    /       \:::\  /:::/    /       \:::\__/:::/    /    
             ~~              \:::\   \/____/             |::::/    /         \:::\/:::/    /         \::::::::/    /     
                              \:::\    \                 /:::/    /           \::::::/    /           \::::::/    /      
                               \:::\____\               /:::/    /             \::::/    /             \::::/    /       
                                \::/    /               \::/    /               \::/____/               \::/____/        
                                 \/____/                 \/____/                 ~~                      ~~
`)

}


// Fonction qui dessine l’échaffaud qui attend le nombe d'essai restant
func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |
   |    o
   |   /|\
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|

__/\\\\\\\\\\\\\____/\\\\\\\\\\\\\\\____/\\\\\\\\\______/\\\\\\\\\\\\_____/\\\________/\\\_______________/\\\____        
 _\/\\\/////////\\\_\/\\\///////////___/\\\///////\\\___\/\\\////////\\\__\/\\\_______\/\\\_____________/\\\\\\\__       
  _\/\\\_______\/\\\_\/\\\_____________\/\\\_____\/\\\___\/\\\______\//\\\_\/\\\_______\/\\\____________/\\\\\\\\\_      
   _\/\\\\\\\\\\\\\/__\/\\\\\\\\\\\_____\/\\\\\\\\\\\/____\/\\\_______\/\\\_\/\\\_______\/\\\___________\//\\\\\\\__     
    _\/\\\/////////____\/\\\///////______\/\\\//////\\\____\/\\\_______\/\\\_\/\\\_______\/\\\____________\//\\\\\___    
     _\/\\\_____________\/\\\_____________\/\\\____\//\\\___\/\\\_______\/\\\_\/\\\_______\/\\\_____________\//\\\____   
      _\/\\\_____________\/\\\_____________\/\\\_____\//\\\__\/\\\_______/\\\__\//\\\______/\\\_______________\///_____  
       _\/\\\_____________\/\\\\\\\\\\\\\\\_\/\\\______\//\\\_\/\\\\\\\\\\\\/____\///\\\\\\\\\/_________________/\\\____ 
        _\///______________\///////////////__\///________\///__\////////////________\/////////__________________\///_____

        `
	case 1:
		draw = `
    ____
   |    |
   |    o
   |   /|\
   |    |
   |
  _|_
 |   |______
 |          |
 |__________|
        `
	case 2:
		draw = `
    ____
   |    |
   |    o
   |    |
   |    |
   |
  _|_
 |   |______
 |          |
 |__________|
        `
	case 3:
		draw = `
    ____
   |    |
   |    o
   |
   |
   |
  _|_
 |   |______
 |          |
 |__________|
        `
	case 4:
		draw = `
    ____
   |    |
   |
   |
   |
   |
  _|_
 |   |______
 |          |
 |__________|
        `
	case 5:
		draw = `
    ____
   |
   |
   |
   |
   |
  _|_
 |   |______
 |          |
 |__________|
        `
	case 6:
		draw = `

   |
   |
   |
   |
   |
  _|_
 |   |______
 |          |
 |__________|
        `
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
        `
	case 8:
		draw = `

        `
	}
	fmt.Println(draw)
}

//Fonction utilitaire qui se charge de mettre un espace entre les lettres d'un SLICE
func drawLatters(l []string) {
	for _,c := range l {
		fmt.Printf("%v ",c)
	}
	fmt.Println()
}

// Fonction qui affiche les lettres + état du jeu
func drawState(g *Game, guess string) {
	fmt.Print("Lettres trouvées: ")
	drawLatters(g.FoundLetters)
	fmt.Print("Lettres utilisées: ")
	drawLatters(g.UsedLetters)
	//Affiche l'état de la lettre testée
	switch g.State {
	case "goodGuess":
		fmt.Println("Bonne pioche!")
		g.Points += 100
	case "alreadyGuessed":
		fmt.Printf("Lettre '%s' déjà utilisée\n", guess)
	case "badGuess":
		fmt.Printf("Mauvaise pioche, la lettre '%s' n'est pas dans le mot\n", guess)
		fmt.Printf("Il vous reste %v vies\n", g.TurnsLeft)
		g.Points -= 50
	case "lost":
		fmt.Print("Vous avez perdu :( ! Le mot était :")
		drawLatters(g.Letters)
		fmt.Printf("\nScore : %v points (100p par bonne réponse, -50p par mauvaise réponse)\n", g.Points)

	case "won":
		fmt.Print("Vous avez gagné :) ! Le mot était :")
		drawLatters(g.Letters)
		fmt.Printf("Il vous restait %v vies !\n", g.TurnsLeft)
		fmt.Printf("Score : %v points (100p par bonne réponse, -50p par mauvaise réponse)\n", g.Points)
	}
}

// Fonction qui va dessiner l’échafaud en fonction de l'état de la partie
// Fonction principale qui appelle plein de petite fonction
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}
