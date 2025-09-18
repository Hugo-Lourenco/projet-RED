package main

import (
	"bufio"
	"fmt"
	"os"
	"projetred"
)

func clearScreen() {
	for i := 0; i < 50; i++ {
		fmt.Println()
	}
}

func main() {
	projetred.C1 = projetred.InitCharacter()

	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		projetred.DisplayInfo()

		projetred.Menu()

		fmt.Print("Appuyez sur Entrée pour continuer...")
		_, _ = reader.ReadString('\n')
	}
}