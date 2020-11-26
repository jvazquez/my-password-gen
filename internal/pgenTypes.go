package internal

import (
	"fmt"
	"github.com/sethvargo/go-diceware/diceware"
	"log"
	"math/rand"
	"strings"
	"time"
)

const DefaultWords int = 6

type PasswordGenerator struct {
	UseSymbols     bool
	Seed           int
	UseFixedSymbol bool
	Separator      string
}

func (p PasswordGenerator) Execute() string {
	var generatedPassword string

	list, err := diceware.Generate(p.Seed)

	if err != nil {
		log.Fatalf("Could not generate seed %s", err)
	}

	if p.UseSymbols == false {
		generatedPassword = p.generateWithSpaces(list)
	} else {
		generatedPassword = p.generateWithSymbols(list)
	}

	return generatedPassword
}

func (p PasswordGenerator) generateWithSpaces(words []string) string {
	return strings.Join(words, " ")
}

func (p PasswordGenerator) generateWithSymbols(words []string) string {
	var passwordBuffer strings.Builder
	for _, word := range words {
		fmt.Fprintf(&passwordBuffer, "%s%s", word, p.getSeparator())
	}

	return passwordBuffer.String()
}

func (p PasswordGenerator) getSeparator() string {
	if p.UseFixedSymbol {
		return p.Separator
	} else {
		defaultSymbols := getDefaultSymbols()
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(defaultSymbols))
		return defaultSymbols[index]
	}
}

func getDefaultSymbols() []string {
	var defaultSymbols []string
	for _, symbol := range strings.Split("! @ # $ % ^ & * ( ) _ - + =", " ") {
		defaultSymbols = append(defaultSymbols, symbol)
	}
	return defaultSymbols
}
