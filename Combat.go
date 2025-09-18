package projetred

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// poison sur le monstre
func poisonPotMonster(m *Monster) {
	for i := 0; i < 3; i++ {
		m.PV -= 10
		if m.PV < 0 {
			m.PV = 0
		}
		fmt.Printf("Effet du poison sur le monstre : %d/%d PV\n", m.PV, m.PV_max)
	}
}

// Combat
func Combat(player *Character, monster *Monster) {
	reader := bufio.NewReader(os.Stdin)
	turn := 1
	skillAvailable := true
	for player.PV > 0 && monster.PV > 0 {
		fmt.Printf("\n===== Tour %d =====\n", turn)
		fmt.Printf("%s : %d/%d PV\n", player.Nom, player.PV, player.PV_max)
		fmt.Printf("%s : %d/%d PV\n", monster.Nom, monster.PV, monster.PV_max)

		// au tour du joueur
		for {
			fmt.Println("Menu :")
			fmt.Println("1) Attaquer")
			fmt.Println("2) Inventaire")
			if skillAvailable && len(player.Skills) > 0 {
				fmt.Println("3) Utiliser Skill")
			}
			fmt.Print("Choix ? ")
			choixRaw, _ := reader.ReadString('\n')
			choix := strings.TrimSpace(choixRaw)

			if choix == "1" {
				// attaque
				dmg := player.Attaque
				monster.PV -= dmg
				if monster.PV < 0 {
					monster.PV = 0
				}
				fmt.Printf("%s inflige %d dégâts à %s\n", player.Nom, dmg, monster.Nom)
				break
			} else if choix == "2" {
				// inventaire
				if len(player.Inventaire) == 0 {
					fmt.Println("Inventaire vide.")
					continue
				}
				fmt.Println("Inventaire :")
				for i, item := range player.Inventaire {
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
				if idx >= 0 && idx < len(player.Inventaire) {
					item := player.Inventaire[idx]
					if item == "Potion de soin" {
						takePot(&C1)
						fmt.Printf("Vous utilisez %s et récupérez %d PV !\n", item, 30)
						break
					} else if item == "Potion de poison" {
						poisonPotMonster(monster)
						fmt.Println("Vous utilisez la potion de poison sur le monstre !")
						player.Inventaire = append(player.Inventaire[:idx], player.Inventaire[idx+1:]...)
						break
					} else {
						fmt.Println("Objet non utilisable en combat.")
						continue
					}
				} else {
					fmt.Println("Choix invalide.")
					continue
				}
			} else if choix == "3" && skillAvailable && len(player.Skills) > 0 {
				msg := player.Skills[0].Use(player, monster)
				fmt.Println(msg)
				skillAvailable = false
				break
			} else {
				fmt.Println("Choix invalide.")
			}
		}

		if monster.PV <= 0 {
			fmt.Printf("%s est vaincu !\n", monster.Nom)
			break
		}

		// au tour du monstre
		var dmg int
		if turn%3 == 0 {
			dmg = monster.Attaque * 2
		} else {
			dmg = monster.Attaque
		}
		player.PV -= dmg
		if player.PV < 0 {
			player.PV = 0
		}
		fmt.Printf("%s inflige à %s %d dégâts\n", monster.Nom, player.Nom, dmg)
		fmt.Printf("PV de %s : %d / %d\n", player.Nom, player.PV, player.PV_max)

		turn++
		if turn%2 == 0 {
			skillAvailable = true
		}
	}

	fmt.Println("Fin du combat.")
	// réssussiter avec 50% de la vie
	player.PV = player.PV_max / 2
	if player.PV < 1 {
		player.PV = 1
	}
	fmt.Printf("%s récupère 50%% de ses PV : %d/%d\n", player.Nom, player.PV, player.PV_max)
	// retour au menu
}
