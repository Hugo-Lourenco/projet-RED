package projetred

import "fmt"

func DisplayInfo() {
	fmt.Println("================= AVENTURE =================")
    fmt.Printf("Joueur: %s  | Classe: %s  | Niveau: %d\n", c1.Nom, c1.Classe, c1.Niveau)
    fmt.Printf("PV: %d/%d  | Attaque: %d  | Or: %d\n", c1.PV, c1.PV_max, c1.Attaque, c1.Or)
    fmt.Println("============================================")
}
