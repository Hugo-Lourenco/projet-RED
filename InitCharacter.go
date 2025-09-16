package projetred

import "fmt"

func InitCharacter() Character{
	var nom string
	fmt.Print("Entrez votre nom :")
	fmt.Scanln(&nom)
	fmt.Printf("nom choisi : %s\n", nom)
	
	tmpl := NewChevalierTemplate()
	character := tmpl.CreateCharacter(nom)
	
	return character
}
