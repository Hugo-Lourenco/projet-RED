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

func InitCharacter() {
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
            c1 = NewMagicienCharacter(nom)
            DisplayInfo(c1)
        case "chevalier":
            c1 = NewChevalierCharacter(nom)
            DisplayInfo(c1)
        default:
            fmt.Println("Classe inexistante — choisissez parmi : chevalier, magicien")
        }
    }
}