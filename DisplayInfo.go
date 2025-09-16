package projetred

import "fmt"

func DisplayInfo(c1 Character) {
	fmt.Println("Nom du personnage : ", c1.Nom)
	fmt.Println("Classe : ", c1.Classe)
	fmt.Println("Niveau du personnage : ", c1.Niveau)
	fmt.Println("Points de vie : ", c1.PV, "/", c1.PV_max)
	fmt.Println("Points de dégâts : ", c1.Attaque)
}
