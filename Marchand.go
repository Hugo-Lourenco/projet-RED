package projetred

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Marchand() {
	fmt.Println("Bienvenue chez le marchand!")
	if c1.Niveau == 1 {
		fmt.Println("Vous êtes nouveau ici ?")
		fmt.Println("Je vais vous faire un cadeau de bienvenue.")
		c1.Inventaire = append(c1.Inventaire, "Potion de soin")
		fmt.Println("Vous avez reçu une Potion de soin!")
	} else {
		for {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Que voulez-vous acheter ?")
			fmt.Print("1) Potion de soin à 20 pieces d'or ")
			fmt.Println("2) Épée en bois à 40 pieces d'or")
			fmt.Println("3) Livre de magie à 50 pieces d'or")
			fmt.Println("4) Épée en fer à 120 pieces d'or")

			choixRaw, _ := reader.ReadString('\n')
			choix := strings.ToLower(strings.TrimSpace(choixRaw))

			switch choix {
			case "1", "potion de soin":
				if c1.Or >= 20 {
					c1.Or -= 20
					c1.Inventaire = append(c1.Inventaire, "Potion de soin")
				}
			case "2", "Épée en bois":
				if c1.Or >= 40 {
					c1.Or -= 40
					c1.Inventaire = append(c1.Inventaire, "Épée en bois")
				}
			case "3", "Livre de magie":
				if c1.Or >= 50 {
					c1.Or -= 50
					c1.Inventaire = append(c1.Inventaire, "Livre de magie")
				}
			case "4", "Épée en fer":
				if c1.Or >= 120 {
					c1.Or -= 120
					c1.Inventaire = append(c1.Inventaire, "Épée en fer")
				}
			default:
				fmt.Println("Je ne possaide pas cette objet, que voulez vous ?")
				continue
			}
		}
	}
}
