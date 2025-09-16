package projetred

func NewChevalierCharacter(nom string) Character {
    return Character{
        nom:        nom,
        classe:     "Chevalier",
        niveau:     1,
        PV_max:     100,
        PV:         50,
        attaque:    6,
        inventaire: []string{},
    }
}

func NewMagicienCharacter(nom string) Character {
    return Character{
        nom:        nom,
        classe:     "Magicien",
        niveau:     1,
        PV_max:     90,
        PV:         40,
        attaque:    8,
        inventaire: []string{},
    }
}
