package projetred

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func formatName(s string) string {
    s = strings.TrimSpace(s)
    runes := []rune(s)
    debutmot := true

    for i, r := range runes {
        if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
            if debutmot {
                if r >= 'a' && r <= 'z' {
                    runes[i] = r - ('a' - 'A')
                }
            } else {
                if r >= 'A' && r <= 'Z' {
                    runes[i] = r + ('a' - 'A')
                }
            }
            debutmot = false
        } else {
            debutmot = true
        }
    }
    return string(runes)
}

// <-- signature modifiée pour retourner un Character
func InitCharacter() Character {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Entrez votre nom : ")
    nomRaw, _ := reader.ReadString('\n')
    nom := formatName(nomRaw)

    for {
        fmt.Print("Choisissez une classe (chevalier/magicien) : ")
        choixRaw, _ := reader.ReadString('\n')
        choix := strings.ToLower(strings.TrimSpace(choixRaw))

        switch choix {
        case "magicien", "mage":
            return NewMagicienCharacter(nom)
        case "chevalier":
            return NewChevalierCharacter(nom)
        default:
            fmt.Println("Classe inexistante — choisissez parmi : chevalier, magicien")
        }
    }
}