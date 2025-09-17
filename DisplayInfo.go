package projetred

import "fmt"

func DisplayInfo() {
	fmt.Println("================= AVENTURE =================")
    fmt.Printf("Joueur: %s  | Classe: %s  | Niveau: %d\n", C1.Nom, C1.Classe, C1.Niveau)
    fmt.Printf("PV: %d/%d  | Attaque: %d  | Or: %d\n", C1.PV, C1.PV_max, C1.Attaque, C1.Or)
    fmt.Println("============================================")
}
