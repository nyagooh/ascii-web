package asciiart

import (
	"fmt"
	"strings"
)

func DisplayAsciiArt(characterMap map[rune][]string, input string) {
	// Check if the character map is empty
	if len(characterMap) == 0 {
		fmt.Println("Error: Character map is empty. Please provide a valid character map.")
		return
	}
	inputSlice := strings.Split(input, "\\n")
	count := 0
	for _, value := range inputSlice {
			for i := 0; i < 8; i++ {
				for _, char := range value {
					line, ok := characterMap[char]
					if ok {
						//store in string
						fmt.Printf("%s ", line[i])
					} else {
						fmt.Println("Characters not found")
						return
					}
				}
				fmt.Println()
			}
		}
	}

