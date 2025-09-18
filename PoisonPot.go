package projetred

import (
	"fmt"
	"time"
)

func poisonPot(c *Character) {
	for i := 0; i < 3; i++ {
		c.PV -= 10
		if c.PV < 0 {
			c.PV = 0
		}
		fmt.Printf("Effet du poison : %d/%d PV\n", c.PV, c.PV_max)
		time.Sleep(1 * time.Second)
	}
	// casser la potion de poison après utilisation
	for i := 0; i < len(c.Inventaire); i++ {
		if c.Inventaire[i] == "Potion de poison" {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			break
		}
	}
}
