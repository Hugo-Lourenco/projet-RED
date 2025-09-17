package projetred

func NewChevalierCharacter(nom string) Character {
	return Character{
		Nom:        nom,
		Classe:     "Chevalier",
		Niveau:     1,
		PV_max:     100,
		PV:         50,
		Attaque:    6,
		Or :		100,
		Inventaire: []string{},
	}
}

func NewMagicienCharacter(nom string) Character {
	return Character{
		Nom:        nom,
		Classe:     "Magicien",
		Niveau:     1,
		PV_max:     90,
		PV:         40,
		Attaque:    8,
		Inventaire: []string{},
	}
}
