package projetred

import "fmt"

func takePot(c1 Character) {
	for i := 0; i < len(c1.Inventaire); i++ {
		if c1.Inventaire[i] == "Potion de soin" {
			soin := 30
			c1.PV += soin
			c1.Inventaire = append(c1.Inventaire[:i], c1.Inventaire[i+1:]...)
			fmt.Printf("Vous utilisez %s et récupérez %d PV !\n", c1.Inventaire[i], soin)
			c1.Inventaire = append(c1.Inventaire[:i], c1.Inventaire[i+1:]...)
			if c1.PV > c1.PV_max {
				c1.PV = c1.PV_max
			}
			break 
		}
	}
}
