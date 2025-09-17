package projetred

type Character struct {
	Nom			string
	Classe		string
	Niveau		int
	PV_max		int
	PV			int
	Attaque		int
	Or 			int 
	Inventaire	[]string
	Skills 		[]Skill
}
var c1 Character