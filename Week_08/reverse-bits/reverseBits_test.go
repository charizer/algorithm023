package reverse_bits

import "testing"

func TestReverseBits(t *testing.T) {
	num := 0b00000010100101000001111010011100
	want := 0b00111001011110000010100101000000
	got := ReverseBits(uint32(num))
	if got != uint32(want) {
		t.Errorf("want:%d got:%d", want, got)
	}
}
