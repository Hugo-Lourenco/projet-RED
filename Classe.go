package projetred

type Chevalier struct {
	Nom         string
	BasePV      int
	PVParNiveau int
	AttaqueBase int
}

func NewChevalierTemplate() Chevalier {
	return Chevalier{
		Nom:         "Chevalier",
		BasePV:      40,
		PVParNiveau: 8,
		AttaqueBase: 6,
	}
}

func (tmpl Chevalier) CreateCharacter(nomJoueur string) Character {
	return Character{
		nom:        nomJoueur,
		classe:     tmpl.Nom,
		niveau:     1,
		PV_max:     tmpl.BasePV,
		PV:         tmpl.BasePV,
		attaque:    tmpl.AttaqueBase,
		inventaire: []string{},
	}
}


