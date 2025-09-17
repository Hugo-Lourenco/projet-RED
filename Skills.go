package projetred

import 	"fmt"


type Skill interface {
	Name() string
	Use(user *Character, target Damageable) string
}

type Damageable interface {
	Applydamage(d int)
	GetPV() int 
	GetPVMax() int 
	GetName() string
	IsDead() bool
}

type HeavyAttack struct {
	Bonus int 
}

func (h HeavyAttack) Name() string {return "Attaque lourde"}

func (h HeavyAttack) Use(user *Character, target Damageable) string {
	if user == nil || target == nil {
		return "Pas de cible"
	}
		dmg := user.Attaque + h.Bonus
		target.Applydamage(dmg)
		msg := fmt.Sprintf("%s utilise %s et inflige %d dégâts à %s", user.Nom, h.Name(),dmg, target.GetName())
	return msg
}

type Fireball struct {
	Bonus int 
}
func (f Fireball) Name() string {return "Boule de feu"}
func (f Fireball) Use(user *Character, target Damageable) string {
	if user == nil || target == nil {
		return "Pas de cible "
	}
	dmg := user.Attaque + f.Bonus 
	target.Applydamage(dmg)
	msg := fmt.Sprintf("%s lance %s et inflige %d dégâts à %s", user.Nom, f.Name(), dmg, target.GetName())
return msg
}

type Thunederbolt struct {
	Bonus int
}
func (t Thunederbolt) Name() string {return "éclair de génie"}
func (t Thunederbolt) Use(user *Character, target Damageable) string {
	if user == nil || target == nil {
		return "Pas de cible"
	}
dmg := user.Attaque + t.Bonus
target. Applydamage(dmg)
msg := fmt.Sprintf("%s lance %s et inflige %d dégâts à %s", user.Nom, t.Name(), dmg, target.GetName())
return msg
}
