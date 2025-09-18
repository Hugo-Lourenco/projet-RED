package projetred

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// poison sur le monstre
func poisonPotMonster(M *Monster) {
    for i := 0; i < 3; i++ {
        M.PV -= 10
        if M.PV < 0 {
            M.PV = 0
        }
        fmt.Printf("Effet du poison sur le monstre : %d/%d PV\n", M.PV, M.PV_max)
    }
}

// Combat 
func Combat(C1 *Character, M *Monster) {
    reader := bufio.NewReader(os.Stdin)
    turn := 1
    skillAvailable := true

    for C1.PV > 0 && M.PV > 0 {
        fmt.Printf("\n===== Tour %d =====\n", turn)
        fmt.Printf("%s : %d/%d PV\n", C1.Nom, C1.PV, C1.PV_max)
        fmt.Printf("%s : %d/%d PV\n", C1.Nom, C1.PV, C1.PV_max)

        // au tour du joueur
        for {
            fmt.Println("Menu :")
            fmt.Println("1) Attaquer")
            fmt.Println("2) Inventaire")
            if skillAvailable && len(C1.Skills) > 0 {
                fmt.Println("3) Utiliser Skill")
            }
            fmt.Print("Choix ? ")
            choixRaw, _ := reader.ReadString('\n')
            choix := strings.TrimSpace(choixRaw)

            if choix == "1" {
                // attaque 
                dmg := C1.Attaque
                C1.PV -= dmg
                if C1.PV < 0 {
                    C1.PV = 0
                }
                fmt.Printf("%s inflige %d dégâts à %s\n", C1.Nom, dmg, M.Nom)
                break
            } else if choix == "2" {
                // inventaire
                if len(C1.Inventaire) == 0 {
                    fmt.Println("Inventaire vide.")
                    continue
                }
                fmt.Println("Inventaire :")
                for i, item := range C1.Inventaire {
                    fmt.Printf("%d) %s\n", i+1, item)
                }
                fmt.Print("Choisissez un objet à utiliser (numéro) ou 0 pour annuler : ")
                itemRaw, _ := reader.ReadString('\n')
                itemChoice := strings.TrimSpace(itemRaw)
                if itemChoice == "0" {
                    continue
                }
                var idx int
                fmt.Sscanf(itemChoice, "%d", &idx)
                idx--
                if idx >= 0 && idx < len(C1.Inventaire) {
                    item := C1.Inventaire[idx]
                    if item == "Potion de soin" {
                        takePot(*C1)
                    } else if item == "Potion de poison" {
                        poisonPotMonster(M)      
                        fmt.Println("Vous utilisez la potion sur le monstre")     
                        C1.Inventaire =  append(C1.Inventaire[:idx], C1.Inventaire[idx+1:]...)
                        break
                    } else {
                        fmt.Println("Objet non utilisable en combat.")
                        continue
                    }
                } else {
                    fmt.Println("Choix invalide.")
                    continue
                }
            } else if choix == "3" && skillAvailable && len(C1.Skills) > 0 {
                msg := C1.Skills[0].Use(C1, M)
                fmt.Println(msg)
                skillAvailable = false
                break
            } else {
                fmt.Println("Choix invalide.")
            }
        }

        if M.PV <= 0 {
            fmt.Printf("%s est vaincu !\n", M.Nom)
            break
        }

        // au tour du monstre
        var dmg int
        if turn%3 == 0 {
            dmg = M.Attaque * 2
        } else {
            dmg = M.Attaque
        }
        C1.PV -= dmg
        if C1.PV < 0 {
            C1.PV = 0
        }
        fmt.Printf("%s inflige à %s %d dégâts\n", M.Nom, C1.Nom, dmg)
        fmt.Printf("PV de %s : %d / %d\n", C1.Nom, C1.PV, C1.PV_max)

        turn++
        if turn%2 == 0 {
            skillAvailable = true
        }
    }

    fmt.Println("Fin du combat.")
    C1.PV = C1.PV_max / 2
    if C1.PV < 1 {
        C1.PV = 1
    }
    fmt.Printf("%s récupère 50%% de ses PV : %d/%d\n", C1.Nom, C1.PV, C1.PV_max)
}