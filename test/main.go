package main

import (
    "fmt"
    "projetred"
)

func main() {
    char := projetred.InitCharacter()
    projetred.DisplayInfo(char)
    fmt.Println("Jeu démarré.")
}