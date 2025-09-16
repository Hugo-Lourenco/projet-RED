package projetred

import "fmt"

func InitCharacter() Character {
	var nom string
	fmt.Print("Entrez votre nom : ")
	fmt.Scanln(&nom)

	var choix string
	fmt.Print("Choisissez une classe (chevalier/magicien) : ")
	fmt.Scanln(&choix)

	switch choix {
	case "magicien", "Mage":
		return NewMagicienCharacter(nom)
	default:
		return NewChevalierCharacter(nom)
	}
}