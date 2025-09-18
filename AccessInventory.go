package projetred

import "fmt"

func accessInventory(c1 Character) {
	fmt.Println("════════════════════════ Inventaire ════════════════════════")
	if len(c1.Inventaire) == 0 {
        fmt.Println("                      	  [ Vide ]                 	     ")
        fmt.Println("════════════════════════════════════════════════════════════")
        return
	}
	for i, item := range c1.Inventaire {
        fmt.Printf(" %2d. %-45s \n", i+1, item)
	}
	fmt.Printf("════════════════════ %d / %d emplacements ═══════════════════\n", len(c1.Inventaire), c1.Max_Inv)
}
