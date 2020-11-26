package internal

import (
	"strings"
	"testing"
)

func TestGenerateWithSpaces(t *testing.T) {
	var fakeWords []string
	fakeWords = append(fakeWords, "horse")
	fakeWords = append(fakeWords, "duck")
	fakeWords = append(fakeWords, "cat")
	fakeWords = append(fakeWords, "car")

	generator := PasswordGenerator{}
	result := generator.generateWithSpaces(fakeWords)

	if strings.Contains(result, " ") == false {
		t.Errorf("generator with spaces does not has spaces.Got: %s", result)
	}

	if len(result) == 0 {
		t.Errorf("generator did not provide words.Got %d", len(result))
	}
}

func TestGenerateWithSymbols(t *testing.T) {
	var fakeWords []string
	fakeWords = append(fakeWords, "horse")
	fakeWords = append(fakeWords, "duck")
	fakeWords = append(fakeWords, "cat")
	fakeWords = append(fakeWords, "car")

	generator := PasswordGenerator{}
	result := generator.generateWithSymbols(fakeWords)

	if strings.Contains(result, " ") == true {
		t.Error("generator with spaces does not has spaces")
	}

	if len(result) == 0 {
		t.Error("generator did not provide words")
	}
}

func TestGetSeparatorWithFixedOption(t *testing.T) {
	fixedSymbol := "#"
	generator := PasswordGenerator{UseFixedSymbol: true, Separator: fixedSymbol}
	result := generator.getSeparator()

	if result != "#" {
		t.Errorf("Using fixed symbol, should return %s. Got %s", fixedSymbol, result)
	}
}

func TestGetSeparatorWithNoFixedOption(t *testing.T) {
	generator := PasswordGenerator{}
	var results []string
	results = append(results, generator.getSeparator())
	results = append(results, generator.getSeparator())

	if results[0] == results[1] {
		t.Error("Should not repeat symbols")
	}
}
