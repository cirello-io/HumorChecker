package HumorChecker

import "testing"

var negAndPosTests = []struct {
	Function string
	Type     string
	In       string
	Expected float64
}{
	{"Negativity", "Score", "bastard", 5.0},
	{"Negativity", "Score", "scumbag", 4.0},
	{"Negativity", "Score", "evil", 3.0},
	{"Negativity", "Score", "ache", 2.0},
	{"Negativity", "Score", "anti", 1.0},

	{"Negativity", "Comparative", "Hey scumbag", 2.0},
	{"Negativity", "Comparative", "I'll be here till 5", 0},

	{"Negativity", "WordCount", "This is two anti evil words", 2.0},

	{"Positivity", "Score", "superb", 5.0},
	{"Positivity", "Score", "amazing", 4.0},
	{"Positivity", "Score", "admire", 3.0},
	{"Positivity", "Score", "amaze", 2.0},
	{"Positivity", "Score", "cool", 1.0},

	{"Positivity", "Comparative", "Hey amazing", 2.0},
	{"Positivity", "Comparative", "I'll be here till 5", 0.0},

	{"Positivity", "WordCount", "This is two amazing cool words", 2.0},
}

func TestNegAndPos(t *testing.T) {
	var f func(string) Score
	for _, test := range negAndPosTests {
		if test.Function == "Negativity" {
			f = Negativity
		} else if test.Function == "Positivity" {
			f = Positivity
		} else {
			t.Fatal("Impossible humor type check")
		}

		result := f(test.In)
		if test.Type == "Score" {
			if result.Score != test.Expected {
				t.Errorf("got wrong score for %s in %s. got: %v. expected: %v", test.Function, test.In, result.Score, test.Expected)
			}
		} else if test.Type == "Comparative" {
			if result.Comparative != test.Expected {
				t.Errorf("got wrong comparative for %s in %s. got: %v. expected: %v", test.Function, test.In, result.Comparative, test.Expected)
			}
		} else if test.Type == "WordCount" {
			if len(result.Words) != int(test.Expected) {
				t.Errorf("got wrong comparative for %s in %s. got: %v. expected: %v", test.Function, test.In, len(result.Words), int(test.Expected))
			}
		} else {
			t.Fatal("Impossible test")
		}
	}
}

func TestAnalyze(t *testing.T) {
	// if v := Analyze("Hey Amazing Scumbag").Score; v != 0 {
	// 	t.Errorf("error analyzing score of sentence. got: %v", v)
	// }
	// if v := Analyze("Cool beans").Score; v != 1 {
	// 	t.Errorf("error analyzing score of sentence: should be positive for only positives. got: %v", v)
	// }
	// if v := Analyze("Hey scumbag").Score; v != -4 {
	// 	t.Errorf("error analyzing score of sentence: should be negative for only negatives. got: %v", v)
	// }
	if v := Analyze("Fearless!").Score; v != 2 {
		t.Errorf("error analyzing score of sentence: punctuation must be ignored (positive). got: %v", v)
	}
	if v := Analyze("Crash!").Score; v != -2 {
		t.Errorf("error analyzing score of sentence: punctuation must be ignored (negative). got: %v", v)
	}
	if v := Analyze("#fearless").Score; v != 2 {
		t.Errorf("error analyzing score of sentence: hashtags must be ignored (positive). got: %v", v)
	}
	if v := Analyze("#crash").Score; v != -2 {
		t.Errorf("error analyzing score of sentence: hashtags must be ignored (negative). got: %v", v)
	}
	// if v := Analyze("An amazing anti").Comparative; v != 1 {
	// 	t.Errorf("error analyzing comparative score of sentence. got: %v", v)
	// }
}
