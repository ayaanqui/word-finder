package main

import (
	"fmt"
	"os"
	"strings"
)

type WordAt struct {
	Value string
	Index int
}

func userData() (string, []WordAt, []string) {
	var user_word string
	fmt.Print("Enter word: ")
	fmt.Scanln(&user_word)

	// add all letter that are not "_" in the user_words_fixed list
	user_words_fixed := []WordAt{}
	for i, letter := range user_word {
		if letter == '_' {
			continue
		}
		user_words_fixed = append(user_words_fixed, WordAt{
			Value: string(letter),
			Index: i,
		})
	}

	var user_letters_raw string
	fmt.Print("Enter letters (Ex. a,l,b,t): ")
	fmt.Scanln(&user_letters_raw)

	user_letters := strings.Split(user_letters_raw, ",")
	available_letters := []string{}
	for _, letter := range user_letters {
		// TODO: this also removes extra repeating letters if any
		if strings.Contains(user_word, letter) {
			continue
		}
		available_letters = append(available_letters, letter)
	}
	return user_word, user_words_fixed, available_letters
}

func main() {
	data, file_err := os.ReadFile("./words.txt")
	if file_err != nil {
		panic(file_err)
	}
	file_words := strings.Split(string(data), "\n")
	if len(file_words) == 0 {
		panic("")
	}

	user_word, user_words_fixed, available_letters := userData()
	fmt.Println(user_word, user_words_fixed, available_letters)

	words_in_len := []string{}
	for _, file_word := range file_words {
		if len(file_word) != len(user_word) {
			continue
		}

		invalid_word := false
		for _, fixed_letter := range user_words_fixed {
			if string(file_word[fixed_letter.Index]) == fixed_letter.Value {
				continue
			}
			invalid_word = true
			break
		}
		if invalid_word {
			continue
		}
		words_in_len = append(words_in_len, file_word)
	}

	valid_words := []string{}
	for _, word := range words_in_len {
		invalid_word := false
		for _, valid_letter := range available_letters {
			if strings.Contains(word, valid_letter) {
				valid_words = append(valid_words, word)
				continue
			}
			invalid_word = true
			break
		}

		if invalid_word {
			continue
		}
	}

	for _, w := range valid_words {
		fmt.Println(w)
	}
}
