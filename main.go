package main

import "fmt"

var alphabet [26]rune

func main() {
	for i := 0; i < 26; i++ {
		alphabet[i] = rune('a' + i)
	}
	// encrypt
	encrypt("ashir", 1)
	// decrypt

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
