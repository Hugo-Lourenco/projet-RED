package main

import (
    "fmt"

    "projetred"
)

func main() {
    char := projetred.InitCharacter()
    // Attention : si les champs de Character sont non-exportés (minuscules),
    // main ne peut pas y accéder directement. Prévois des getters ou rends
    // les champs exportés si tu veux les afficher ici.
    _ = char
    fmt.Println("Jeu démarré.")
}