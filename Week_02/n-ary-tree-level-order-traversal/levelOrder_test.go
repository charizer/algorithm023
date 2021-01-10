package n_ary_tree_level_order_traversal

import "testing"

func creatTree(arr []int) *Node {
	root := &Node{Val: 1}
	c1 := &Node{Val: 3}
	c2 := &Node{Val: 2}
	c3 := &Node{Val: 4}
	root.Children = []*Node{c1,c2,c3}
	cc1 := &Node{Val: 5}
	cc2 := &Node{Val: 6}
	c1.Children = []*Node{cc1,cc2}
	return root
}

func TestLevelOrder(t *testing.T) {
	root := creatTree([]int{1,3,2,4,5,6})
	want := [][]int{{1},{3,2,4},{5,6}}
	got := LevelOrder(root)
	for idx, g := range got {
		for i := range g {
			if got[idx][i] != want[idx][i] {
				t.Errorf("want:%d got:%d idx:%d, i:%d", want[idx], got[idx], idx, i)
			}
		}
	}
}

func TestLevelOrder2(t *testing.T) {
	root := creatTree([]int{1,3,2,4,5,6})
	want := [][]int{{1},{3,2,4},{5,6}}
	got := LevelOrder2(root)
	for idx, g := range got {
		for i := range g {
			if got[idx][i] != want[idx][i] {
				t.Errorf("want:%d got:%d idx:%d, i:%d", want[idx], got[idx], idx, i)
			}
		}
	}
}
