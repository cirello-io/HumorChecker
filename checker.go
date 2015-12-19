package HumorChecker // "cirello.io/HumorChecker"

import (
	"bufio"
	"regexp"
	"strings"
)

type sign int

const (
	positive sign = +1
	negative sign = -1
)

//Score is the result of sentiment calculation
type Score struct {
	// Score is the sum of the sentiment points of the analyzed text.
	// Negativity will render negative points only, and vice-versa.
	Score float64

	// Comparative establishes a ratio of sentiment per word
	Comparative float64

	// List of words for a given sentiment.
	Words []string
}

//FullScore is the difference between positive and negative sentiment
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

func calculateScore(phrase string, calcSign sign) Score {
	var hits float64
	var words []string
	var count int

	var lettersAndSpaceOnly = regexp.MustCompile(`[^a-zA-Z ]+`)
	scanner := bufio.NewScanner(strings.NewReader(strings.ToLower(lettersAndSpaceOnly.ReplaceAllString(phrase, " "))))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
		word := scanner.Text()
		if v, ok := afinn[word]; ok {
			if (calcSign == positive && v > 0) || (calcSign == negative && v < 0) {
				hits += v * float64(calcSign)
				words = append(words, word)
			}
		}
	}

	return Score{
		Score:       hits,
		Comparative: hits / float64(count),
		Words:       words,
	}
}

// Negativity calculates the negative sentiment of a sentence
func Negativity(phrase string) Score {
	return calculateScore(phrase, negative)
}

// Positivity calculates the positive sentiment of a sentence
func Positivity(phrase string) Score {
	return calculateScore(phrase, positive)
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
