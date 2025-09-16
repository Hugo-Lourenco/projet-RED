package main

import (
    "fmt"
    "projetred"
)

func main() {
    _ = projetred.InitCharacter()
    projetred.DisplayInfo()
    fmt.Println("Jeu démarré.")
}