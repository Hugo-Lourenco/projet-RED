package projetred

func takePot(C1 Character) {
	for i := 0; i < len(C1.Inventaire); i++ {
		if C1.Inventaire[i] == "Potion de soin" {
			C1.PV += 30
			C1.Inventaire = append(C1.Inventaire[:i], C1.Inventaire[i+1:]...)
			if C1.PV > C1.PV_max {
				C1.PV = C1.PV_max
			}
			break 
		}
	}
}
