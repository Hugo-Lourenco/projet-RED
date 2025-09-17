package projetred

import "fmt"

func CanAddItem(inventaire []string) bool {
	return len(inventaire) < 10
}

func AddItem(c1 *Character, item string) {
	if CanAddItem(c1.Inventaire) {
		c1.Inventaire = append(c1.Inventaire, item)
		fmt.Printf("Ajouté : %s (%d/10)\n", item, len(c1.Inventaire))
	} else {
		fmt.Println("Ton inventaire est plein !")
	}
}
