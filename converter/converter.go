package converter

import (
	"strings"
)

func StringToMorse(text string, charToMorse map[rune]string) string {
	text = strings.ToLower(text)
	var morseString string
	for i, c := range text {
		morseString += charToMorse[c]
		if i != len(text)-1 {
			morseString += " "
		}
	}
	return morseString
}

func GetCharToMorse() map[rune]string {
	charToMorse := make(map[rune]string)
	charToMorse['a'] = ".-"
	charToMorse['b'] = "-..."
	charToMorse['c'] = "-.-."
	charToMorse['d'] = "-.."
	charToMorse['e'] = "."
	charToMorse['f'] = "..-."
	charToMorse['g'] = "--."
	charToMorse['h'] = "...."
	charToMorse['i'] = ".."
	charToMorse['j'] = ".---"
	charToMorse['k'] = "-.-"
	charToMorse['l'] = ".-.."
	charToMorse['m'] = "--"
	charToMorse['n'] = "-."
	charToMorse['o'] = "---"
	charToMorse['p'] = ".--."
	charToMorse['q'] = "--.-"
	charToMorse['r'] = ".-."
	charToMorse['s'] = "..."
	charToMorse['t'] = "-"
	charToMorse['u'] = "..-"
	charToMorse['v'] = "...-"
	charToMorse['w'] = ".--"
	charToMorse['x'] = "-..-"
	charToMorse['y'] = "-.--"
	charToMorse['z'] = "--.."
	charToMorse['0'] = "-----"
	charToMorse['1'] = ".----"
	charToMorse['2'] = "..---"
	charToMorse['3'] = "...--"
	charToMorse['4'] = "....-"
	charToMorse['5'] = "....."
	charToMorse['6'] = "-...."
	charToMorse['7'] = "--..."
	charToMorse['8'] = "---.."
	charToMorse['9'] = "----."
	charToMorse[' '] = "/"
	return charToMorse
}
