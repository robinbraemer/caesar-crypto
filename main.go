package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to Caesar Crypto by Robin Brämer!")
	fmt.Printf("Let's En- & Decrypt your text.\n\n")

	fmt.Print("Do you want to encrypt or decrypt? (e/d):")
	s, _ := readString()
	if s == "e" {
		fmt.Println()
		letsEncrypt()
	} else {
		letsDecrypt()
	}
}

func letsEncrypt() {
	fmt.Print("[ENCRYPT] Enter your key (int):")

	key, err := readKey()
	if err != nil {
		fmt.Println("[ENCRYPT] FAIL! Please insert an integer next time.")
		return
	}

	fmt.Print("[ENCRYPT] Enter your text (string):")

	txt, err := readString()
	if err != nil {
		fmt.Println("[ENCRYPT] FAIL! Please insert an string next time.")
		return
	}

	fmt.Println("[ENCRYPT] OK! Let's encrypt:", txt)
	en, err := encrypt(txt, key)
	if err != nil {
		fmt.Println("[ENCRYPT] FAIL!", err.Error())
		return
	}
	fmt.Println("[ENCRYPT] DONE! Your encryption is:", en)
}

func letsDecrypt() {
	fmt.Print("[DECRYPT] Enter your key (int):")

	key, err := readKey()
	if err != nil {
		fmt.Println("[DECRYPT] FAIL! Please insert an integer next time.")
		return
	}

	fmt.Print("[DECRYPT] Enter your text (string):")

	txt, err := readString()
	if err != nil {
		fmt.Println("[DECRYPT] FAIL! Please insert an string next time.")
		return
	}

	fmt.Println("[DECRYPT] OK! Let's decrypt:", txt)
	de, err := decrypt(txt, key)
	if err != nil {
		fmt.Println("[DECRYPT] FAIL!", err.Error())
		return
	}
	fmt.Println("[DECRYPT] DONE! Your decryption is:", de)
}

var alphabet = []string{
	"a", "b", "c",
	"d", "e", "f",
	"g", "h", "i",
	"j", "k", "l",
	"m", "n", "o",
	"p", "q", "r",
	"s", "t", "u",
	"v", "w", "x",
	"y", "z"}

func encrypt(txt string, key int) (string, error) {
	var buf bytes.Buffer

	for i := 0; i < len(txt); i++ {
		sl, err := shiftLetter(string(txt[i]), key, true)
		if err != nil {
			return "", err
		}

		buf.WriteString(sl)
	}

	return buf.String(), nil
}

func decrypt(txt string, key int) (string, error) {
	var buf bytes.Buffer

	for i := 0; i < len(txt); i++ {
		sl, err := shiftLetter(string(txt[i]), key, false)
		if err != nil {
			return "", err
		}

		buf.WriteString(sl)
	}

	return buf.String(), nil
}

func shiftLetter(letter string, key int, shiftForward bool) (string, error) {
	i, err := findLetterIndex(letter)
	if err != nil {
		return "", err
	}

	if shiftForward {
		if i+key <= len(alphabet) {
			return alphabet[i+key], nil
		}

		// TODO shifte öffter durch alphabet
		return alphabet[i+key-len(alphabet)], nil
	} else {
		if i-key >= 0 {
			return alphabet[i-key], nil
		}

		// TODO shifte öffter durch alphabet
		return alphabet[i-key+len(alphabet)], nil
	}
}

func findLetterIndex(letter string) (int, error) {
	le := strings.ToLower(letter)
	for i, c := range alphabet {
		if le == c {
			return i, nil
		}
	}
	return 0, errors.New("Letter not in alphabet.")
}

func readKey() (int, error) {
	var key int
	_, err := fmt.Scanf("%d", &key)
	return key, err
}

func readString() (string, error) {
	var input string
	fmt.Scanln(&input)
	return input, nil
}
