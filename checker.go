package HumorChecker // import "cirello.io/HumorChecker"

import (
	"bufio"
	"regexp"
	"strings"
)

// analysis is the complete sentiment calculation
type analysis struct {
	// positivityScore is the sum of the positive sentiment points of the
	// analyzed text.
	positivityScore float64

	// negativityScore is the sum of the negativity sentiment points of the
	// analyzed text.
	negativityScore float64

	// positivityComparative establishes a ratio of sentiment per positive
	// word
	positivityComparative float64

	// negativityComparative establishes a ratio of sentiment per negative
	// word
	negativityComparative float64

	// positiveWords is the list of positive words for a given sentiment.
	positiveWords []string

	// negativeWords is the list of negative words for a given sentiment.
	negativeWords []string
}

// Score is the result of sentiment calculation
type Score struct {
	// Score is the sum of the sentiment points of the analyzed text.
	// Negativity will render negative points only, and vice-versa.
	Score float64

	// Comparative establishes a ratio of sentiment per word
	Comparative float64

	// Words is the list of words for a given sentiment.
	Words []string
}

// FullScore is the difference between positive and negative sentiment
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

func calculateScore(phrase string) analysis {
	var phits, nhits float64
	var pwords, nwords []string
	var count int

	scanner := bufio.NewScanner(strings.NewReader(strings.ToLower(lettersAndSpaceOnly.ReplaceAllString(phrase, " "))))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
		word := scanner.Text()
		if v, ok := afinn[word]; ok {
			if v > 0 {
				phits += v
				pwords = append(pwords, word)
			} else if v < 0 {
				nhits -= v
				nwords = append(nwords, word)
			}
		}
	}

	return analysis{
		positivityScore:       phits,
		positivityComparative: phits / float64(count),
		positiveWords:         pwords,
		negativityScore:       nhits,
		negativityComparative: nhits / float64(count),
		negativeWords:         nwords,
	}
}

func renderNegativeScore(a analysis) Score {
	return Score{
		Score:       a.negativityScore,
		Comparative: a.negativityComparative,
		Words:       a.negativeWords,
	}
}

// Negativity calculates the negative sentiment of a sentence
func Negativity(phrase string) Score {
	return renderNegativeScore(calculateScore(phrase))
}

func renderPositiveScore(a analysis) Score {
	return Score{
		Score:       a.positivityScore,
		Comparative: a.positivityComparative,
		Words:       a.positiveWords,
	}
}

// Positivity calculates the positive sentiment of a sentence
func Positivity(phrase string) Score {
	return renderPositiveScore(calculateScore(phrase))
}

// Analyze calculates overall sentiment
func Analyze(phrase string) FullScore {
	analysis := calculateScore(phrase)

	return FullScore{
		Score:       analysis.positivityScore - analysis.negativityScore,
		Comparative: analysis.positivityComparative - analysis.negativityComparative,
		Positive:    renderPositiveScore(analysis),
		Negative:    renderNegativeScore(analysis),
	}
}
