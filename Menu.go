package projetred

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func Menu() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("1) Information perso")
    fmt.Println("2) Inventaire")
    fmt.Println("3) Quitter")
    fmt.Print("Choix 1, 2 ou 3 ? ")
    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)

    switch choice {
    case "1":
        DisplayInfo()
    case "2":
        accessInventory(c1)
    case "3":
        os.Exit(0)
    default:
        fmt.Println("Choix invalide.")
    }
}