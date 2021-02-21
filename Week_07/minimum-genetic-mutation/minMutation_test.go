package minimum_genetic_mutation

import "testing"

func TestMinMutation(t *testing.T) {
	start := "AACCGGTT"
	end := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	want := 1
	got := MinMutation(start, end, bank)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	start = "AACCGGTT"
	end = "AAACGGTA"
	bank = []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	want = 2
	got = MinMutation(start, end, bank)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	/*"AACCGGTT"
	"AAACGGTA"
	["AACCGATT","AACCGATA","AAACGATA","AAACGGTA"]*/
	start = "AACCGGTT"
	end = "AAACGGTA"
	bank = []string{"AACCGATT","AACCGATA","AAACGATA","AAACGGTA"}
	want = 4
	got = MinMutation(start, end, bank)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
