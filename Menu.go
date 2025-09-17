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
    fmt.Println(") Marchand")
    fmt.Print("Choix 1, 2 , 3 ou 4 ? ")
    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)

    for {
        switch choice {
        case "1":
            DisplayInfo()
        case "2":
            accessInventory(c1)
        case "3":
          os.Exit(0)
        case "4" : 
            Marchand()
        default:
            fmt.Println("Choix invalide.")
        continue 
        }
    }
}