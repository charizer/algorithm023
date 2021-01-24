package walking_robot_simulation

import "testing"

func TestRobotSim(t *testing.T) {
	commands := []int{4,-1,3}
	obstacles := [][]int{}
	want := 25
	got := RobotSim(commands, obstacles)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	commands = []int{4,-1,4,-2,4}
	obstacles = [][]int{{2,4}}
	want = 65
	got = RobotSim(commands, obstacles)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
