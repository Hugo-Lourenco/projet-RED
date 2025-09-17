package projetred

import "fmt"

func CanAddItem(inventaire []string) bool {
	return len(inventaire) < C1.Max_Inv
}

func AddItem(c1 *Character, item string) {
	if CanAddItem(c1.Inventaire) {
		c1.Inventaire = append(c1.Inventaire, item)
		fmt.Printf("Ajouté : %s (%d/%d)\n", item, len(c1.Inventaire), c1.Max_Inv)
	} else {
		fmt.Println("Ton inventaire est plein !")
	}
}
