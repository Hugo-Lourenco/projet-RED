package projetred

func takePot(c1 *Character) {
	for i := 0; i < len(c1.Inventaire); i++ {
		if c1.Inventaire[i] == "potion de soin" {
			c1.PV += 50
		}
		if c1.PV > c1.PV_max {
			c1.PV = c1.PV_max
		}
	}
}
