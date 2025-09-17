package projetred


func NewChevalierCharacter(nom string) Character {
	c1 := Character{
		Nom:        nom,
		Classe:     "Chevalier",
		Niveau:     1,
		PV_max:     100,
		PV:         50,
		Attaque:    6,
		Or:	100,
		Inventaire: []string{},
		Max_Inv: 	10,
	}
	c1.Skills = []Skill{HeavyAttack{Bonus: 4}}
	return c1
}

func NewMagicienCharacter(nom string) Character {
	c1 := Character{
		Nom:        nom,
		Classe:     "Magicien",
		Niveau:     1,
		PV_max:     90,
		PV:         40,
		Attaque:    8,
		Or: 	100,
		Inventaire: []string{},
		Max_Inv: 	10,
	}
	c1.Skills = []Skill{Fireball{Bonus: 6}}
	return c1
}
