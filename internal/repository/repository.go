package repository

import (
	"fume/config"
	"fume/internal/storage"
	"fume/internal/utils"
	"log/slog"
	"strings"
)

type Repository struct {
	DataBase  *storage.Storage
	Config 	  *config.Config
	Logger 	  *slog.Logger
}

func NewRepository(config *config.Config, logger *slog.Logger, database *storage.Storage) (*Repository, error) {
	return &Repository{
		Config: config,
		Logger: logger,
		DataBase: database,
	}, nil
}

func (r *Repository) Run(s string) []string {
	arr := strings.Split(s, " ")

	alphabet := r.Config.Alphabet
	dictionary := r.Config.Dictionary

	output := make([]string, len(arr))

	for i, word := range arr {
		var answer []string
		if strings.Contains(word, "*") {
			BackTracking(word, 0, &alphabet, &dictionary, &answer)
			if len(answer) > 0 {
				output[i] = answer[0]
			} else {
				output[i] = "unrecognized"
			}
		} else {
			output[i] = word
		}
	}

	return output
}

func BackTracking(word string, index int, alphabet *[]string, dict *[]string, answer *[]string) {
	for index < len(word) && word[index] != '*' {
		index += 1
	}

	if index >= len(word) {
		if utils.Contains(*dict, word) {
			*answer = append(*answer, word)
		}
		return
	}

	for _, character := range *alphabet {
		newWord := word[:index] + character + word[index+1:]
		BackTracking(newWord, index+1, alphabet, dict, answer)
	}
}
