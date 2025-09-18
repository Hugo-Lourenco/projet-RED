package projetred

import "fmt"

func accessInventory(c1 Character) {
	// Permet d'afficher l'inventaire 
	fmt.Println("════════════════════════ Inventaire ════════════════════════")
	if len(c1.Inventaire) == 0 {
        fmt.Println("                      	  [ Vide ]                 	     ")
        fmt.Println("════════════════════════════════════════════════════════════")
        return
	}
	for i := 0; i < len(c1.Inventaire); i++ {
        item := c1.Inventaire[i]
        fmt.Printf("%d) %s\n", i+1, item)
	}
	fmt.Printf("════════════════════ %d / %d emplacements ═══════════════════\n", len(c1.Inventaire), c1.Max_Inv)
}
