package projetred

import "fmt"

func IsDead(C1 *Character) bool {
	if C1.PV <= 0 {
		fmt.Println("Vous êtes mort !")
		C1.PV = C1.PV_max / 2
		fmt.Printf("Vous êtes ressuscité avec 50%% des PV : %d/%d PV\n", C1.PV, C1.PV_max)
		return true
	}
	return false
}
