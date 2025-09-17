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
	Max_Inv		int
}
var C1 Character