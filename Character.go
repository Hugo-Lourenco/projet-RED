package projetred

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PV_max     int
	PV         int
	Attaque    int
	Or         int
	Inventaire []string
	Skills     []Skill
	Max_Inv    int
}

var C1 Character

// simule le tour du joueur
func charTurn(player *Character, monster *Monster) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Menu :")
		fmt.Println("1) Attaquer")
		fmt.Println("2) Inventaire")
		fmt.Print("Choix ? ")
		choixRaw, _ := reader.ReadString('\n')
		choix := strings.TrimSpace(choixRaw)
		switch choix {
		case "1":
			// Attaque basique
			dmg := 5
			monster.PV -= dmg
			if monster.PV < 0 {
				monster.PV = 0
			}
			fmt.Printf("%s inflige %d dégâts à %s\n", player.Nom, dmg, monster.Nom)
			fmt.Printf("%s : PV %d / %d\n", monster.Nom, monster.PV, monster.PV_max)
			return
		case "2":
			fmt.Println("Inventaire :", player.Inventaire)
			// ici tu peux ajouter la logique pour utiliser un objet
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
