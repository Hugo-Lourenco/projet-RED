package projetred

import "fmt"

func IsDead(c *Character) bool {
	if c.PV <= 0 {
		fmt.Println("Vous êtes mort !")
		if c.PV_max > 0 {
			c.PV = c.PV_max / 2
			if c.PV < 1 {
				c.PV = 1
			}
			fmt.Printf("Vous êtes ressuscité avec 50%% des PV : %d/%d PV\n", c.PV, c.PV_max)
		} else {
			fmt.Println("Impossible de ressusciter : PV_max = 0")
		}
		return true
	}
	return false
}
