package design_a_leaderboard

import "testing"

func TestConstructor(t *testing.T) {
	leaderboard := Constructor()
	leaderboard.AddScore(1,73)   // leaderboard = [[1,73]];
	leaderboard.AddScore(2,56)   // leaderboard = [[1,73],[2,56]];
	leaderboard.AddScore(3,39)   // leaderboard = [[1,73],[2,56],[3,39]];
	leaderboard.AddScore(4,51)   // leaderboard = [[1,73],[2,56],[3,39],[4,51]];
	leaderboard.AddScore(5,4)    // leaderboard = [[1,73],[2,56],[3,39],[4,51],[5,4]];
	k := 1
	want := 73
	got := leaderboard.Top(k)           // returns 73;
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	leaderboard.Reset(1)         // leaderboard = [[2,56],[3,39],[4,51],[5,4]];
	leaderboard.Reset(2)         // leaderboard = [[3,39],[4,51],[5,4]];
	leaderboard.AddScore(2,51)   // leaderboard = [[2,51],[3,39],[4,51],[5,4]];
	k = 3
	want = 141
	got = leaderboard.Top(3)           // returns 141 = 51 + 51 + 39;
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
