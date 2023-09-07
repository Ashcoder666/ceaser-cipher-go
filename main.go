package main

import "fmt"

var alphabet [26]rune

func main() {
	for i := 0; i < 26; i++ {
		alphabet[i] = rune('a' + i)
	}
	// encrypt
	encrypt("ashizk", 20)
	// decrypt
	decrypt("umbcte", 20)

}

// encrypt
func encrypt(str string, shiftt int) {

	for j := 0; j < len(str); j++ {
		for index, letter := range alphabet {
			if letter == rune(str[j]) {
				if shitedChar := rune(str[j]) + rune(shiftt); shitedChar > 'z' {
					differenceFromStart := shitedChar - 123
					fmt.Printf("%c", alphabet[differenceFromStart])
				} else {
					fmt.Printf("%c", alphabet[index+shiftt])
				}

			}

		}

	}
	fmt.Println("")
}

func decrypt(str string, shiftt int) {
	for j := 0; j < len(str); j++ {
		for index, letter := range alphabet {
			if letter == rune(str[j]) {
				if unshiftedChar := rune(str[j]) - rune(shiftt); unshiftedChar < 97 {
					difference := 97 - unshiftedChar
					fmt.Printf("%c", alphabet[26-difference])

				} else {
					fmt.Printf("%c", alphabet[index-shiftt])
				}

			}

		}

	}
	fmt.Println("")
}
