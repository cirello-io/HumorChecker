package HumorChecker // "cirello.io/HumorChecker"

import (
	"bufio"
	"regexp"
	"strings"
)

type Score struct {
	// Score is the sum of the sentiment points of the analyzed text.
	// Negativity will render negative points only, and vice-versa.
	Score float64

	// Comparative establishes a ratio of sentiment per word
	Comparative float64

	// List of words for a given sentiment.
	Words []string
}

type FullScore struct {
	// Score is the difference between positive and negative sentiment
	// scores.
	Score float64

	// Comparative is the difference between positive and negative sentiment
	// comparative scores.
	Comparative float64

	// Positive score object
	Positive Score

	// Negative score object
	Negative Score
}

var lettersAndSpaceOnly = regexp.MustCompile(`[^a-zA-Z ]+`)

// Negativity calculates the negative sentiment of a sentence
func Negativity(phrase string) Score {
	var hits float64
	var words []string

	addPush := func(t string, score float64) {
		hits -= score
		words = append(words, t)
	}

	scanner := bufio.NewScanner(strings.NewReader(strings.ToLower(lettersAndSpaceOnly.ReplaceAllString(phrase, " "))))
	scanner.Split(bufio.ScanWords)

	var count float64
	for scanner.Scan() {
		count++
		word := scanner.Text()
		if v, ok := afinn[word]; ok && v < 0 {
			addPush(word, v)
		}
	}

	return Score{
		Score:       hits,
		Comparative: hits / count,
		Words:       words,
	}
}

// Positiviy calculates the positive sentiment of a sentence
func Positivity(phrase string) Score {
	var hits float64
	var words []string

	addPush := func(t string, score float64) {
		hits += score
		words = append(words, t)
	}

	scanner := bufio.NewScanner(strings.NewReader(strings.ToLower(lettersAndSpaceOnly.ReplaceAllString(phrase, " "))))
	scanner.Split(bufio.ScanWords)

	var count float64
	for scanner.Scan() {
		count++
		word := scanner.Text()
		if v, ok := afinn[word]; ok && v > 0 {
			addPush(word, v)
		}
	}

	return Score{
		Score:       hits,
		Comparative: hits / count,
		Words:       words,
	}
}

// Analyze calculates overall sentiment
func Analyze(phrase string) FullScore {
	pos := Positivity(phrase)
	neg := Negativity(phrase)

	return FullScore{
		Score:       pos.Score - neg.Score,
		Comparative: pos.Comparative - neg.Comparative,
		Positive:    pos,
		Negative:    neg,
	}
}
