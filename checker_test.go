package HumorChecker

import "testing"

var negAndPosTests = []struct {
	Function            string
	Type                string
	In                  string
	Expected            int
	ExpectedComparative float64
}{
	{"Negativity", "Score", "bastard", 5.0, 0},
	{"Negativity", "Score", "scumbag", 4.0, 0},
	{"Negativity", "Score", "evil", 3.0, 0},
	{"Negativity", "Score", "ache", 2.0, 0},
	{"Negativity", "Score", "anti", 1.0, 0},
	{"Negativity", "WordCount", "This is two anti evil words", 2.0, 0},

	{"Positivity", "Score", "superb", 5.0, 0},
	{"Positivity", "Score", "amazing", 4.0, 0},
	{"Positivity", "Score", "admire", 3.0, 0},
	{"Positivity", "Score", "amaze", 2.0, 0},
	{"Positivity", "Score", "cool", 1.0, 0},
	{"Positivity", "WordCount", "This is two amazing cool words", 2.0, 0},

	{"Negativity", "Comparative", "Hey scumbag", 0, 2.0},
	{"Negativity", "Comparative", "I'll be here till 5", 0, 0},
	{"Positivity", "Comparative", "Hey amazing", 0, 2.0},
	{"Positivity", "Comparative", "I'll be here till 5", 0, 0.0},
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
			if result.Comparative != test.ExpectedComparative {
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

func BenchmarkAnalyze(b *testing.B) {
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent ut ligula semper, posuere nisl nec, sagittis tortor. Quisque tincidunt ante a auctor facilisis. Aenean dui lectus, tincidunt semper nisl non, convallis aliquet tortor. Vivamus hendrerit vehicula tempus. Suspendisse tellus elit, venenatis mollis dapibus nec, faucibus ut est. Proin a dolor et dolor gravida viverra in sed est. Pellentesque vulputate turpis vitae velit finibus mattis.

Integer ultrices ac nisi non sodales. Sed velit felis, efficitur nec dapibus non, malesuada eu sem. Curabitur laoreet libero lacinia varius tristique. Ut dictum quis ex sit amet pulvinar. Sed tincidunt, ante non efficitur ornare, erat ligula porta risus, sit amet porttitor ligula odio vulputate ante. Phasellus porttitor faucibus sagittis. Ut malesuada consectetur venenatis. Pellentesque eu magna at dolor pretium maximus. Pellentesque porta velit odio, sit amet efficitur tellus commodo sed. Nam vitae pulvinar metus. Integer vitae risus ac est mattis rutrum vitae a lacus.

In suscipit, risus at malesuada fringilla, dolor lorem interdum orci, at ullamcorper leo nunc sed purus. Vivamus molestie ultrices velit, et auctor dolor accumsan eu. Phasellus tincidunt tempus cursus. Fusce sit amet mattis justo, tincidunt congue neque. Aenean rutrum dolor vel tristique lacinia. Proin aliquet libero a enim iaculis maximus. Nulla est nisi, tincidunt ut feugiat ac, dictum sed tellus. Vestibulum elementum erat nec sem hendrerit sagittis. Etiam non sapien et libero fringilla elementum vel vitae augue. Aenean volutpat fermentum lorem. Nunc imperdiet ex in erat egestas, accumsan cursus diam iaculis. Donec non vehicula nisl, quis fermentum dolor. Pellentesque ultricies ex sit amet volutpat maximus. Vestibulum nec ligula sit amet sem semper tincidunt ut at magna. Integer nec pretium magna.

Phasellus pulvinar orci et dictum fermentum. Ut vestibulum, turpis eget blandit molestie, magna mauris consequat nulla, vitae varius nisi turpis vel nisi. Integer ac eros enim. Integer eu pharetra arcu. Quisque id luctus dui. Curabitur fringilla, nulla quis hendrerit facilisis, justo dui rutrum libero, at auctor nisi nisl viverra sem. Praesent ornare efficitur tincidunt. Vestibulum venenatis bibendum tortor, vel laoreet lacus sagittis eu. Maecenas accumsan, tortor eu euismod dignissim, dui nisl mollis nisl, ac varius nibh lorem sed velit. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec sit amet congue massa. Donec eget massa sed enim imperdiet efficitur. Ut tempus porta nulla, ac sollicitudin dui egestas vitae. Nullam euismod erat aliquet mattis molestie. Phasellus placerat, turpis euismod iaculis iaculis, massa quam euismod quam, eu aliquet ligula lectus quis justo. Suspendisse placerat ligula in egestas ullamcorper.

Fusce at eleifend tellus, ac luctus purus. Praesent eu eros sit amet odio bibendum scelerisque. Vestibulum in dui nec ex vulputate semper a vitae lorem. Donec rutrum ex sit amet mauris iaculis consectetur. Fusce ornare luctus augue ut laoreet. Proin dui diam, laoreet eget laoreet et, ultrices eget leo. Pellentesque ac mauris in elit placerat tincidunt id vitae urna. Donec id mauris in velit volutpat convallis. Mauris pulvinar orci sapien, sed pellentesque arcu elementum ac. Nulla facilisi. Cras aliquet mauris et sapien aliquam convallis. Curabitur eu tristique diam.`
	for n := 0; n < b.N; n++ {
		Analyze(text)
	}
}
