package main

import (
	"fmt"

	asciiart "asciiart/functionFiles"
)

func main() {
	// handle paths
	//open server
	characterMap, err := asciiart.CreateMap(filePath)
	if err != nil {
		fmt.Println("Error reading map:", err)
		return
	}
	if len(characterMap) == 0 {
		fmt.Println("File is empty")
		return
	}
	asciiart.DisplayAsciiArt(characterMap, input)
}
