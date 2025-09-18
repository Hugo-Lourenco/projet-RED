package projetred

import "fmt"

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

// pattern d'attaque du goblin tous les 3 tours il tape x2
func goblinPattern(goblin *Monster, target *Character, turn int) {
	var dmg int
	if turn%3 == 0 {
		dmg = goblin.Attaque * 2
	} else {
		dmg = goblin.Attaque
	}
	target.PV -= dmg
	if target.PV < 0 {
		target.PV = 0
	}
	fmt.Printf("%s inflige à %s %d dégâts\n", goblin.Nom, target.Nom, dmg)
	fmt.Printf("PV de %s : %d / %d\n", target.Nom, target.PV, target.PV_max)
}

// combat tour par tour joueur vs monstre
func trainingFight(player *Character, monster *Monster) {
	turn := 1
	for C1.PV > 0 && M.PV > 0 {
		fmt.Printf("\n===== Tour %d =====\n", turn)
		charTurn(&C1, &M)
		if M.PV <= 0 {
			fmt.Printf("%s est vaincu !\n", M.Nom)
			break
		}
		goblinPattern(monster, player, turn)
		if player.PV <= 0 {
			fmt.Printf("%s est vaincu !\n", C1.Nom)
			break
		}
		turn++
	}
	fmt.Println("Fin du combat d'entraînement.")
}
