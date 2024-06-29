package main

import (
	"bytes"
	"fmt"
	"morse/converter"
	"morse/sound"
)

func main() {
	runSound()
	fmt.Println("leave")
}

func runSound() {
	charToMorse := converter.GetCharToMorse()
	textString := "SOS"
	morseString := converter.StringToMorse(textString, charToMorse)

	FREQUENCY := 2000

	fmt.Println(morseString)
	for _, c := range morseString {
		var buffer []byte
		switch c {
		case '-':
			fmt.Print("-")
			buffer = sound.SquareBuffer(FREQUENCY, 1)
		case '.':
			fmt.Print(".")
			buffer = sound.SquareBuffer(FREQUENCY, 3)
		case ' ':
			fmt.Print(" ")
			buffer = sound.SilentBuffer(2)
		case '/':
			fmt.Print("/")
			buffer = sound.SilentBuffer(6)
		default:
			fmt.Printf("Incorrect char passed %c\n", c)
			return
		}

		gapBuffer := sound.SilentBuffer(1)
		buffer = append(buffer, gapBuffer...)
		audioReader := bytes.NewReader(buffer)
		sound.PlaySound(audioReader)
	}
	fmt.Print("\n")
}
