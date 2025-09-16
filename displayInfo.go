package projetred

import "fmt"

func displayInfo(c1 Character) {
	fmt.Println(c1.nom)
	fmt.Println(c1.classe)
	fmt.Println(c1.niveau)
	fmt.Println(c1.PV_max)
	fmt.Println(c1.PV)
	fmt.Println(c1.attaque)
}