package mergeoverlap

import "testing"

//import "fmt"

func Test_MergeOverlap_1(t *testing.T) {
	shit := [][2]int{{1, 10}, {3, 5}, {4, 5}, {14, 20}}
	damn := MergeOverlap(shit)
	for i, j := range damn {
		switch {
		case i == 0 && j[0] == 1 && j[1] == 10:
		case i == 1 && j[0] == 14 && j[1] == 20:
		default:
			t.Error("Fuck")
		}
	}
}
func Test_MergeOverlap_2(t *testing.T) {
	shit := [][2]int{{0, 10}, {3, 5}, {4, 15}, {14, 50}, {51, 77}}
	damn := MergeOverlap(shit)
	for i, j := range damn {
		switch {
		case i == 0 && j[0] == 0 && j[1] == 50:
		case i == 1 && j[0] == 51 && j[1] == 77:
		default:
			t.Logf("$#v\n", damn)
			t.Error("Fuck")
		}
	}
}
