package projetred

func takePot(C1 *Character) {
	// Vérifie que le joueur possède un potion
	for i, item := range C1.Inventaire {
		if item == "Potion de soin" {
			// Ajoute les PV
			C1.PV += 30
			if C1.PV > C1.PV_max {
				C1.PV = C1.PV_max
			}
			// Supprime la potion de soin de l'inventaire
			C1.Inventaire = append(C1.Inventaire[:i], C1.Inventaire[i+1:]...)
			break
		}
	}
}
