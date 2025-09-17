package projetred

type Monster struct {
    Nom			string
	Niveau		int
	PV_max		int
	PV			int
	Attaque		int
}

func Monstenemy(nom string) Monster {
	return Monster{
	Nom: 		"Goblin", 
	PV_max:		100,
	PV:			100,
	Attaque: 	5,
	}	
}
