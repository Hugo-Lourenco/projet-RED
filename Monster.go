package projetred


var M Monster

type Monster struct {
	Nom     string
	Niveau  int
	PV_max  int
	PV      int
	Attaque int
}

//  monstre damageable
func (m *Monster) Applydamage(d int) {
	m.PV -= d
	if m.PV < 0 {
		m.PV = 0
	}
}

func (m *Monster) GetPV() int {
	return m.PV
}

func (m *Monster) GetPVMax() int {
	return m.PV_max
}

func (m *Monster) GetName() string {
	return m.Nom
}

func (m *Monster) IsDead() bool {
	return m.PV <= 0
}

// fonction goblin d'entraînement
func Monstenemy(nom string) Monster {
	m := Monster{
		Nom:     "Goblin",
		PV_max:  100,
		PV:      100,
		Attaque: 5,
	}
	return m
}
