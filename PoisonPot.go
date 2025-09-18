package projetred

import (
	"fmt"
	"time"
)

func poisonPot(M *Monster, C1 Character) {
	for i := 0; i < 3; i++ {
		M.PV -= 10
		if M.PV < 0 {
			M.PV = 0
		}
		fmt.Printf("Effet du poison : %d/%d PV\n", M.PV, M.PV_max)
		time.Sleep(1 * time.Second)
	}
	for i := 0; i < len(C1.Inventaire); i++ {
		if C1.Inventaire[i] == "Potion de poison" {
			C1.Inventaire = append(C1.Inventaire[:i], C1.Inventaire[i+1:]...)
			break
		}
	}
}
