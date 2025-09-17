package projetred

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Marchand() {
	fmt.Println("Bienvenue chez le marchand!")
	if C1.Niveau == 1 {
		fmt.Println("Vous êtes nouveau ici ?")
		fmt.Println("Je vais vous faire un cadeau de bienvenue.")
		C1.Inventaire = append(C1.Inventaire, "Potion de soin")
		fmt.Println("Vous avez reçu une Potion de soin!")
	}
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Que voulez-vous acheter ?")
		fmt.Println("1) Potion de soin à 20 pièces d'or ")
		fmt.Println("2) Épée en bois à 40 pièces d'or")
		fmt.Println("3) Livre de magie réservé aux magiciens à 50 pièces d'or")
		fmt.Println("4) Épée en fer à 120 pièces d'or")
		if C1.Max_Inv < 40 {
			fmt.Println("5) Sac à dos à 50")
		}
		fmt.Println("6) Retour")

		choixRaw, _ := reader.ReadString('\n')
		choix := strings.ToLower(strings.TrimSpace(choixRaw))

		switch choix {
		case "1", "potion de soin":
			if C1.Or >= 20 {
				C1.Or -= 20
				C1.Inventaire = append(C1.Inventaire, "Potion de soin")
				fmt.Println("Vous avez reçu une Potion de soin!")
			} else {
				fmt.Println("Vous n'avez pas assez de pièces d'or")
				continue
			}
		case "2", "Épée en bois":
			if C1.Or >= 40 {
				C1.Or -= 40
				C1.Inventaire = append(C1.Inventaire, "Épée en bois")
				fmt.Println("Vous avez reçu une épée en bois!")
			} else {
				fmt.Println("Vous n'avez pas assez de pièces d'or")
				continue
			}
		case "3", "Livre de magie":
			if C1.Or >= 50 && C1.Classe == "Magicien" {
				C1.Or -= 50
				C1.Inventaire = append(C1.Inventaire, "Livre de magie")
				fmt.Println("Vous avez reçu une livre de magie!")
			} else if C1.Or < 50 {
				fmt.Println("Vous n'avez pas assez de pièces d'or")
				continue
			} else {
				fmt.Println("Vous n'êtes pas magicien, vous ne pouvez acheter cette objet")
				continue
			}
		case "4", "Épée en fer":
			if C1.Or >= 120 {
				C1.Or -= 120
				C1.Inventaire = append(C1.Inventaire, "Épée en fer")
				fmt.Println("Vous avez reçu une épée en fer!")
			} else {
				fmt.Println("Vous n'avez pas assez de pièces d'or")
				continue
			}
		case "5", "Sac à dos":
		if C1.Or >= 50 {
			C1.Or -= 50
			C1.Max_Inv += 10
			fmt.Println("Vous avez reçu un sac à dos!")
		} else {
			fmt.Println("Vous n'avez pas assez de pièces d'or")
			continue
		}
		case "6", "retour":
			return
		default:
			fmt.Println("Je ne possède pas cette objet, que voulez vous ?")
			continue

		}
	}
}
