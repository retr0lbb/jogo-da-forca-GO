package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	sorce := rand.NewSource(time.Now().UnixNano())
	r := rand.New(sorce)

	words := []string{"google", "semanadeti", "golang", "drive", "maps", "firebase", "grasshopper"}

	wordToGuess := words[r.Intn(len(words))]
	guessedWord := strings.Repeat("_", len(wordToGuess))
	maxErrors := 6
	errors := 0

	for errors < maxErrors {
		fmt.Println("JOGO DA FORCA SEMANA DE TI")
		drawHangman(errors)
		fmt.Println("A palavra é:", guessedWord)
		fmt.Print("Letra adivinhada: ")
		var letter string
		fmt.Scan(&letter)

		if wordToGuess == letter {
			fmt.Println("Parabéns, você acertou a palavra: ", wordToGuess)
			break
		}

		if strings.Contains(wordToGuess, letter) {
			// Atualize guessedWord com a letra corretamente adivinhada
			guessedWord = updateGuessedWord(wordToGuess, guessedWord, letter)
		} else {
			errors++
			drawHangman(errors)
		}
	}

	if errors >= maxErrors {
		fmt.Println("Você perdeu, a palavra certa era: ", wordToGuess)
	}
}

func updateGuessedWord(wordToGuess, guessedWord, letter string) string {
	newGuessedWord := ""
	for i := 0; i < len(wordToGuess); i++ {
		if wordToGuess[i] == letter[0] {
			newGuessedWord += string(letter[0])
		} else {
			newGuessedWord += string(guessedWord[i])
		}
	}
	return newGuessedWord
}

func drawHangman(errors int) {
	hangman := []string{
		"  ________     ",
		"  |       |    ",
		"  |            ",
		"  |            ",
		"  |            ",
		"  |            ",
		" _|_           ",
	}

	switch errors {
	case 0:
		break
	case 1:
		hangman[2] = "  |       O    "
	case 2:
		hangman[2] = "  |       O    "
		hangman[3] = "  |       |    "
	case 3:
		hangman[2] = "  |       O    "
		hangman[3] = "  |      /|    "
	case 4:
		hangman[2] = "  |       O    "
		hangman[3] = "  |      /|\\  "
	case 5:
		hangman[2] = "  |       O    "
		hangman[3] = "  |      /|\\  "
		hangman[4] = "  |      /     "
	case 6:
		hangman[2] = "  |       O    "
		hangman[3] = "  |      /|\\  "
		hangman[4] = "  |      / \\    "
	}

	for _, line := range hangman {
		fmt.Println(line)
	}
}
