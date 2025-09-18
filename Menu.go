package projetred

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("1) Information perso")
		fmt.Println("2) Inventaire")
		fmt.Println("3) Marchand")
		fmt.Println("4) Quitter")
		fmt.Println("5) Retour")
		fmt.Println("6) Entrainement")
		fmt.Print("Choix 1, 2, 3, 4, 5 ou 6 ? ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			DisplayInfo()
		case "2":
			accessInventory(C1)
		case "3":
			Marchand()
		case "4":
			os.Exit(0)
		case "5":
			return
		case "6":
			goblin := Monstenemy("Goblin")
			Combat(&C1, &goblin)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
