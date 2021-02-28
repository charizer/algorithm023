package number_of_1_bits

import "testing"

func TestHammingWeight(t *testing.T) {
	num := 0b00000000000000000000000000001011
	want := 3
	got := HammingWeight(uint32(num))
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
